package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabaseProvider interface {
	Collection(name string) CollectionProvider
	Client() ClientProvider
}

type CollectionProvider interface {
	FindOne(context.Context, interface{}) SingleResultProvider
	InsertOne(context.Context, interface{}) (interface{}, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
}

type SingleResultProvider interface {
	Decode(v interface{}) error
}

type ClientProvider interface {
	Database(string) MongoDatabaseProvider
	Connect() error
	StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

func NewClientProvider(url string) (ClientProvider, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(url))

	return &mongoClient{cl: c}, err

}

func NewMongoDatabase(dbName string, client ClientProvider) MongoDatabaseProvider {
	return client.Database(dbName)
}

func (mc *mongoClient) Database(dbName string) MongoDatabaseProvider {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect() error {
	return mc.cl.Connect(nil)
}

func (md *mongoDatabase) Collection(colName string) CollectionProvider {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() ClientProvider {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResultProvider {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
