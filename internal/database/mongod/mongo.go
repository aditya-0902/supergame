package mongod

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Option func(*MongoRepository)

type MongoRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, opts ...Option) *MongoRepository {
	repo := &MongoRepository{
		client: client,
	}

	for _, opt := range opts {
		opt(repo)
	}

	return repo
}

func WithCollection(dbName string, collectionName string) Option {
	return func(m *MongoRepository) {
		m.collection = m.client.Database(dbName).Collection(collectionName)
	}
}

func WithDatabase(dbName string) Option {
	return func(m *MongoRepository) {
		m.db = m.client.Database(dbName)
	}
}

func (m *MongoRepository) SetCollection(dbName, collectionName string) {
	m.collection = m.client.Database(dbName).Collection(collectionName)
}
