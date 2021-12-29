package domain

import (
	"reflect"
)

type Product struct {
	ID    interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title string      `json:"title" bson:"title" validate:"required"`
}

// ProductRepository ...
type ProductRepository interface {
	CreateMany(databaseName, collectionName string, books []Product) (interface{}, error)
	Read(databaseName, collectionName string, filter interface{}, limit int64, dataModel reflect.Type) (interface{}, error)
	Update(databaseName, collectionName string, filter, update interface{}) (interface{}, error)
	Delete(databaseName, collectionName string, filter interface{}) (interface{}, error)
}

// ProductUsecase ..
type ProductUsecase interface {
	InsertBooks(books *[]Product) (interface{}, error)
	ListBooks(limit int64, dataModel reflect.Type) (interface{}, error)
	UpdateBook(update Product) (interface{}, error)
	DeleteBook(bookID string) (interface{}, error)
}
