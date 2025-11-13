package repository

import (
	"fmt"
)

// DBType is an enum for the database types
type DBType string

const (
	Postgres DBType = "postgres"
	MongoDB  DBType = "mongodb"
)

// Repositories holds all the entity-specific repositories
type Repositories struct {
	UserRepo     UserRepository
  ProductRepo  ProductRepository
  CategoryRepo CategoryRepository
}

func NewRepositories(dbType DBType, connString string) (*Repositories, error) {
	switch dbType {
	case Postgres:
		// Initialize ALL PostgreSQL concrete implementations
		fmt.Printf("Initializing PostgreSQL Repositories...\n")
		return &Repositories{
			UserRepo:     NewPostgresUserRepository(connString),
      ProductRepo:  NewPostgresProductRepository(connString),
      CategoryRepo: NewPostgresCategoryRepository(connString),
		}, nil
	case MongoDB:
		// Initialize ALL MongoDB concrete implementations
    /*
		fmt.Printf("Initializing MongoDB Repositories...\n")
		return &Repositories{
			UserRepo:     NewMongoDBUserRepository(connString),
		}, nil
    */
		return &Repositories{
			UserRepo:     nil,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
