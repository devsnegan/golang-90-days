package main

import "fmt"

type UserReader interface {
	GetUser(id int)
}

type UserWriter interface {
	CreateUser(name string)
}

type UserRepository interface {
	UserReader
	UserWriter
}

type PostgresRepository struct{}

func (p PostgresRepository) GetUser(id int) {
	fmt.Println("Getting user:", id)
}

func (p PostgresRepository) CreateUser(name string) {
	fmt.Println("Creating user:", name)
}

func main() {
	var repo UserRepository = PostgresRepository{}

	repo.GetUser(101)
	repo.CreateUser("Snegan")
}