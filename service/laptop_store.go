package service


import (
	"sync"
	"errors"
	"fmt"

	"example.com/laptop_store/proto"
	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface{
	// saves a laptop to the store
	Save(*pb.Laptop) error
	// finds a laptop by ID
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct{
	mutex sync.RWMutex
	data map[string]*pb.Laptop
}

func NewInMemoryLaptopStore () *InMemoryLaptopStore{
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}


func (store *InMemoryLaptopStore)Save (laptop *pb.Laptop) error{
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil{
		return ErrAlreadyExists
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil{
		return fmt.Errorf("Cannot Copy laptop data: %w", err)
	}
	store.data[other.Id] = other
	return nil
}


func (store *InMemoryLaptopStore)Find(id string)(*pb.Laptop, error){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	laptop := store.data[id]
	if laptop == nil{
		return nil, fmt.Errorf("There is no Laptop with %s", id)
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil{
		return nil, fmt.Errorf("Cannot Copy laptop data: %w", err)
	}
	return other, nil
}