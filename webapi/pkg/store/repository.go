package store

import (
	"context"
	"database/sql"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"webapi/schema/v1"
)

type Repository[T v1.Entity] interface {
	Count() (int64, error)
	List() ([]*T, error)
	Get(id uint) (*T, error)
	Create(item *T) (*T, error)
	Update(item *T) error
	Delete(id uint) error
}

func CreateRepository[T v1.Entity](db *DbStorage) Repository[T] {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	switch typeName {
	case "User":
		repo := &OrmRepository[v1.User]{db: db}
		return any(repo).(Repository[T])
	case "Project":
		repo := &OrmRepository[v1.Project]{db: db}
		return any(repo).(Repository[T])
	case "Task":
		repo := &OrmRepository[v1.Task]{db: db}
		return any(repo).(Repository[T])
	default:
		panic("Unknown type")
	}
}

type OrmRepository[T v1.Entity] struct {
	db *DbStorage
}

var _ Repository[v1.Entity] = &OrmRepository[v1.Entity]{}

func (o *OrmRepository[T]) Count() (int64, error) {
	tableName := v1.GetTableName[T]()
	if o.db.DB != nil {
		var count int64
		result := o.db.DB.Table(tableName).Count(&count)
		if result.Error != nil {
			return 0, result.Error
		}
		return count, nil
	} else if o.db.MongoClient != nil {
		collection := o.db.MongoDb.Collection(tableName)
		count, err := collection.CountDocuments(context.Background(), bson.D{})
		if err != nil {
			return 0, err
		}
		return count, nil
	} else {
		//TODO implement me
		panic("implement me")
	}
}

func (o *OrmRepository[T]) List() ([]*T, error) {
	tableName := v1.GetTableName[T]()
	if o.db.DB != nil {
		rows, err := o.db.DB.Table(tableName).Rows()
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if err != nil {
			return nil, err
		}

		var items []*T
		for rows.Next() {
			item := new(T)
			err := o.db.DB.ScanRows(rows, item)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		return items, nil
	} else if o.db.MongoClient != nil {
		collection := o.db.MongoDb.Collection(tableName)
		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			return nil, err
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			_ = cursor.Close(ctx)
		}(cursor, context.Background())

		var items []*T
		if err = cursor.All(context.Background(), &items); err != nil {
			return nil, err
		}
		return items, nil
	} else {
		//TODO implement me
		panic("implement me")
	}
}

func (o *OrmRepository[T]) Get(id uint) (*T, error) {
	tableName := v1.GetTableName[T]()
	if o.db.DB != nil {
		item := new(T)
		result := o.db.DB.Table(tableName).First(item, id)
		if result.Error != nil {
			return nil, result.Error
		}
		return item, nil
	} else if o.db.MongoClient != nil {
		var item T
		collection := o.db.MongoDb.Collection(tableName)
		err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&item)
		if err != nil {
			return nil, err
		}
		return &item, nil
	} else {
		//TODO implement me
		panic("implement me")
	}
}

func (o *OrmRepository[T]) Create(item *T) (*T, error) {
	tableName := v1.GetTableName[T]()
	if o.db.DB != nil {
		result := o.db.DB.Create(item)
		if result.Error != nil {
			return nil, result.Error
		}
		return item, nil
	} else if o.db.MongoClient != nil {
		collection := o.db.MongoDb.Collection(tableName)
		_, err := collection.InsertOne(context.Background(), item)
		if err != nil {
			return nil, err
		}
		return item, nil
	} else {
		//TODO implement me
		panic("implement me")
	}
}

func (o *OrmRepository[T]) Update(item *T) error {
	tableName := v1.GetTableName[T]()
	if o.db.DB != nil {
		result := o.db.DB.Save(item)
		if result.Error != nil {
			return result.Error
		}
		return nil
	} else if o.db.MongoClient != nil {
		collection := o.db.MongoDb.Collection(tableName)
		_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": any(item).(v1.Entity).GetID()}, item)
		return err
	} else {
		//TODO implement me
		panic("implement me")
	}
}

func (o *OrmRepository[T]) Delete(id uint) error {
	tableName := v1.GetTableName[T]()
	if o.db.DB != nil {
		result := o.db.DB.Table(tableName).Delete(nil, id)
		if result.Error != nil {
			return result.Error
		}
		return nil
	} else if o.db.MongoClient != nil {
		collection := o.db.MongoDb.Collection(tableName)
		_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
		return err
	} else {
		//TODO implement me
		panic("implement me")
	}
}
