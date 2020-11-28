package repository

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/rokoga/filas-backend/domain"
)

const (
	ERROR_NOT_FOUND          = "Não foi encontrado o estabelecimento"
	ERROR_REMOVE_CONSUMER    = "Não foi possível remover o consumidor"
	ERROR_NOT_FOUND_CONSUMER = "Não foi possível encontrar consumidor"
)

type MockStore struct {
	aStore []*domain.Store
}

type StoreMockRepositoryImpl struct {
	mockStore MockStore
}

func NewStoreMockRepository() StoreRepository {
	return &StoreMockRepositoryImpl{
		mockStore: MockStore{
			aStore: nil,
		},
	}
}

func (repo *StoreMockRepositoryImpl) Create(store *domain.Store) (*domain.Store, error) {

	if store.ID == "" {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		store.ID = strconv.Itoa(r1.Int())
	}

	repo.mockStore.aStore = append(repo.mockStore.aStore, store)

	return store, nil
}

func (repo *StoreMockRepositoryImpl) Get(id string) (*domain.Store, error) {
	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			return elem, nil
		}
	}
	return nil, errors.New(ERROR_NOT_FOUND)
}

func (repo *StoreMockRepositoryImpl) GetStore(URLname string) (*domain.Store, error) {
	for _, elem := range repo.mockStore.aStore {
		if elem.URLName == URLname {
			return elem, nil
		}
	}

	return nil, errors.New(ERROR_NOT_FOUND)
}

func (repo *StoreMockRepositoryImpl) AddConsumer(id string, consumer *domain.Consumer) error {

	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			elem.Queue = append(elem.Queue, consumer)
			return nil
		}
	}

	return errors.New(ERROR_NOT_FOUND)
}

func (repo *StoreMockRepositoryImpl) RemoveConsumer(id string, phone string) error {

	for i, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			copy(repo.mockStore.aStore[i:], repo.mockStore.aStore[i+1:])
			repo.mockStore.aStore[len(repo.mockStore.aStore)-1] = nil
			repo.mockStore.aStore = repo.mockStore.aStore[:len(repo.mockStore.aStore)-1]

			return nil
		}
	}

	return errors.New("ERROR_REMOVE_CONSUMER")
}

func (repo *StoreMockRepositoryImpl) GetConsumer(id string, phone string) (*domain.Consumer, error) {

	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			for _, consumer := range elem.Queue {
				if consumer.Number == phone {
					return consumer, nil
				}
			}
		}
	}

	return nil, errors.New(ERROR_NOT_FOUND_CONSUMER)
}

func (repo *StoreMockRepositoryImpl) GetAllConsumers(id string) ([]*domain.Consumer, error) {

	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			return elem.Queue, nil
		}
	}

	return nil, errors.New(ERROR_NOT_FOUND)
}
