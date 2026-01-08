package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

// Repository interface for data access
type UserRepository interface {
	Save(user User) error
	FindByID(id int) (*User, error)
	Delete(id int) error
}

// PostgreSQL implementation
type PostgresRepo struct {
	connectionString string
}

func (p PostgresRepo) Save(user User) error {
	fmt.Printf("Saving user to PostgreSQL: %+v\n", user)
	return nil
}

func (p PostgresRepo) FindByID(id int) (*User, error) {
	fmt.Printf("Finding user %d in PostgreSQL\n", id)
	return &User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

func (p PostgresRepo) Delete(id int) error {
	fmt.Printf("Deleting user %d from PostgreSQL\n", id)
	return nil
}

// MongoDB implementation
type MongoRepo struct {
	database string
}

func (m MongoRepo) Save(user User) error {
	fmt.Printf("Saving user to MongoDB: %+v\n", user)
	return nil
}

func (m MongoRepo) FindByID(id int) (*User, error) {
	fmt.Printf("Finding user %d in MongoDB\n", id)
	return &User{ID: id, Name: "Jane", Email: "jane@example.com"}, nil
}

func (m MongoRepo) Delete(id int) error {
	fmt.Printf("Deleting user %d from MongoDB\n", id)
	return nil
}

// Business logic - doesn't care about implementation
type UserService struct {
	repo UserRepository
}

func (s UserService) CreateUser(name, email string) error {
	user := User{Name: name, Email: email}
	return s.repo.Save(user)
}

func (s UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func main() {
	// Postgress service implementation
	postgresService := UserService{repo: PostgresRepo{}}
	postgresService.CreateUser("Alice", "alice@example.com")

	// MongoDB service implementation
	mongoService := UserService{repo: MongoRepo{}}
	mongoService.CreateUser("Bob", "bob@example.com")
	mongoService.DeleteUser(1)
}
