package service

import (
	"fmt"

	"sync"
)

type UserStore interface{
	Save(u *User) error
	Find(username string)(*User, error)
}

type InMemoryUserStore struct{
	mutex sync.RWMutex 
	data map[string]*User
}

func NewInMemoryUserStore ()*InMemoryUserStore{
	return &InMemoryUserStore{
		data:make(map[string]*User),
	}
}

func (store *InMemoryUserStore)Save(u *User)error{
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[u.Username] != nil{
		return ErrAlreadyExists
	}

	store.data[u.Username] = u.Clone()
	return nil

}

func (store *InMemoryUserStore)Find(username string)(*User, error){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	user := store.data[username]
	if user == nil{
		return nil, fmt.Errorf("There is no User by this username")
	}
	return user.Clone(), nil
}
