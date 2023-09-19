package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}
type PostgresNumberStore struct{
	// some db connection
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 3, 9,7,99}, nil
}

func (p PostgresNumberStore) Put(number int) error {
	fmt.Println("Store the number into the mongoDB storage")
	return nil
}

type MongoDBNumberStore struct {
	//some values
}



func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 3, 9}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("Store the number into the mongoDB storage")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	err:=apiServer.numberStore.Put(12)
	if err != nil{
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(numbers)

}
