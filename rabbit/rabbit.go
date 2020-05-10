package rabbit

import (
	"context"
	"errors"
	"time"

	"github.com/streadway/amqp"
	"github.com/zhs/loggr"
)

const (
	timeout = 600 * time.Second
)

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	ctx     context.Context
	url     string
}

type Publisher func(body []byte) error

func NewConnection(ctx context.Context, url string) (*Connection, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn := &Connection{
		ctx: ctx,
		url: url,
	}
	err := conn.connect()
	if err != nil {
		return nil, errors.New("can't connect to Rabbit after timeout")
	}
	go conn.reconnect()
	return conn, nil
}

func (c *Connection) connect() error {
	log := loggr.WithContext(c.ctx)
	err := Retry(c.ctx, 2*time.Second, func() error {
		conn, err := amqp.DialConfig(c.url, amqp.Config{
			Dial: amqp.DefaultDial(time.Second * 120),
		})
		if err != nil {
			log.Warn("Trying to connect rabbit...")
			return err
		}

		ch, err := conn.Channel()
		if err != nil {
			return err
		}

		c.conn = conn
		c.channel = ch
		log.Info("Connected to rabbit")
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) reconnect() {
	errs := c.conn.NotifyClose(make(chan *amqp.Error))
	for {
		select {
		case <-errs:
			err := c.connect()
			if err != nil {
				return
			}
		}
	}
}

func (c *Connection) Consumer(exchange string, key string) (<-chan amqp.Delivery, error) {
	q, err := c.channel.QueueDeclare(
		"",
		true,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err := c.channel.QueueBind(
		q.Name,
		key,
		exchange,
		false,
		nil); err != nil {
		return nil, err
	}

	msgs, err := c.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (c Connection) NewPublisher(exchange string, key string) (Publisher, error) {
	return func(body []byte) error {
		err := c.channel.Publish(
			exchange,
			key,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			return err
		}
		return nil
	}, nil
}

func Retry(ctx context.Context, interval time.Duration, f func() error) error {
	tick := time.NewTicker(interval)
	defer tick.Stop()

	var err error
	for {
		if err = f(); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return err
		case <-tick.C:
		}
	}
}
