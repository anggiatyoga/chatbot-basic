package config

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDatabase(conf MongoDB) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()
	credential := options.Credential{
		AuthMechanism: conf.AuthMechanism,
		Username:      conf.Username,
		Password:      conf.Password,
		AuthSource:    conf.AuthSource,
	}

	// Server Production
	option := options.Client().
		ApplyURI(conf.UriServer).
		SetAuth(credential)

	// Local Dev
	//option := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(option)
	if err != nil {
		log.Error().Str("mongo", "init").Msg(fmt.Sprintf("mongo error %s", err))
		panic(fmt.Sprintf("mongo error %s", err))
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Error().Str("mongo", "connect").Msg(fmt.Sprintf("mongo error %s", err))
		panic(fmt.Sprintf("mongo error %s", err))
	}

	db := client.Database(conf.CredentialName)
	return db
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
