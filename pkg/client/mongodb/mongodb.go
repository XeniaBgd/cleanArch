package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authSource string) (*mongo.Database, error) {
	var mongoDBURL string
	var anon bool
	if username == "" || password == "" {
		anon = true
		mongoDBURL = fmt.Sprintf("mongodb://#{host}:#{port}")
	} else {
		mongoDBURL = fmt.Sprintf("mongodb:/#{username}:#{password}@#{host}:#{port}")
	}

	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if !anon {
		clientOptions.SetAuth(options.Credential{
			AuthSource:  authSource,
			Username:    username,
			Password:    password,
			PasswordSet: true,
		})
	}

	client, err := mongo.Connect(reqCtx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	return client.Database(database), nil
}
