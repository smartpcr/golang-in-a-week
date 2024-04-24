package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"webapi/pkg/config"
	"webapi/schema/v1"
)

type DbStorage struct {
	DB          *gorm.DB
	MongoClient *mongo.Client
	MongoDb     *mongo.Database
}

func NewDbStorage(cfg *config.DbConfig) (*DbStorage, error) {
	switch cfg.Type {
	case config.MySQL:
		dnsMySql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.Host, cfg.Port, cfg.DBName)
		mysqlDialect := mysql.Open(dnsMySql)
		db, err := gorm.Open(mysqlDialect, &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return &DbStorage{
			DB: db,
		}, nil
	case config.PgSQL:
		dnsPostgres := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.DBUser, cfg.DBPassword, cfg.DBName)
		pgDialect := postgres.Open(dnsPostgres)
		db, err := gorm.Open(pgDialect, &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return &DbStorage{
			DB: db,
		}, nil
	case config.Mongo:
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return nil, err
		}
		db := client.Database(cfg.DBName)
		return &DbStorage{
			MongoClient: client,
			MongoDb:     db,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
}

func (m *DbStorage) Init() error {
	if m.DB != nil {
		err := m.DB.AutoMigrate(&v1.User{}, &v1.Project{}, &v1.Task{})
		if err != nil {
			return err
		}
	}

	if m.MongoClient != nil {
		collOptions := options.CreateCollectionOptions{}
		*collOptions.Capped = true
		*collOptions.SizeInBytes = 1048576
		err := m.MongoDb.CreateCollection(context.TODO(), "Users", &collOptions)
		if err != nil {
			return err
		}
		err = m.MongoDb.CreateCollection(context.TODO(), "Projects", &collOptions)
		if err != nil {
			return err
		}
		err = m.MongoDb.CreateCollection(context.TODO(), "Tasks", &collOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *DbStorage) Close() {
	if m.MongoClient != nil {
		err := m.MongoClient.Disconnect(nil)
		if err != nil {
			log.Printf("Error closing the mongo client: %v", err)
		}
	}

	if m.DB != nil {
		db, err := m.DB.DB()
		if err != nil {
			log.Printf("Error reading underlying db: %v", err)
		} else if db != nil {
			err = db.Close()
			if err != nil {
				log.Printf("Error closing the database connection: %v", err)
			}
		}
	}
}
