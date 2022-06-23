package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// New create a new mongodb instance with the given options.
func NewMongodbPool(opts *MongodbConf) (*mongo.Client, error) {
	var dns string
	if len(opts.Username) > 0 {
		dns = fmt.Sprintf(`mongodb://%s:%s@%s/?authSource=%s`,
			opts.Username,
			opts.Password,
			opts.Host,
			opts.Database,
		)
	} else {
		dns = fmt.Sprintf(`mongodb://%s/?authSource=%s`,
			opts.Host,
			opts.Database,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//连接池最大连接数默认最大100
	options.Client().SetMaxPoolSize(opts.MaxPoolSize)

	//mongo最大同时连接数默认为2
	options.Client().SetMaxConnecting(opts.MaxOpenConnections)

	//mongo最大空余连接存在时长默认不限

	options.Client().SetMaxConnIdleTime(opts.MaxConnectionIdleTime * time.Millisecond)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dns))
	if err != nil {
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}

func InitMongodbPool(mongodbConf *MongodbConf) error {
	//普通的db方式
	var err error
	if len(mongodbConf.Host) == 0 {
		fmt.Printf("[INFO] %s%s\n", time.Now().Format("2006-01-02 15:04:05"), " empty mongo config.")
		return errors.New("empty mongo config")
	}
	MongodbClient, err = NewMongodbPool(mongodbConf)
	if err != nil {
		return err
	}

	Mongodb = MongodbClient.Database(mongodbConf.Database)
	if Mongodb == nil {
		return errors.New("get mongodb pool error")
	}

	return nil
}

func GetMongodb(name string) *mongo.Database {

	return MongodbClient.Database(name)
}
