package repository

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/rokoga/filas-backend/domain"
)

// MockStore implements
type MockStore struct {
	aStore []*domain.Store
}

// StoreMockRepositoryImpl implements
type StoreMockRepositoryImpl struct {
	mockStore MockStore
}

// NewStoreMockRepository implements
func NewStoreMockRepository() StoreRepository {
	return &StoreMockRepositoryImpl{
		mockStore: MockStore{
			aStore: nil,
		},
	}
}

// Create implements
func (repo *StoreMockRepositoryImpl) Create(store *domain.Store) (*domain.Store, error) {

	if store.ID == "" {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		store.ID = strconv.Itoa(r1.Int())
	}

	repo.mockStore.aStore = append(repo.mockStore.aStore, store)

	return store, nil
}

// RemoveStore implements
func (repo *StoreMockRepositoryImpl) RemoveStore(id string) error {

	for i, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			copy(repo.mockStore.aStore[i:], repo.mockStore.aStore[i+1:])
			repo.mockStore.aStore[len(repo.mockStore.aStore)-1] = nil
			repo.mockStore.aStore = repo.mockStore.aStore[:len(repo.mockStore.aStore)-1]

			return nil
		}
	}

	return errors.New(ErrorNotFoundStore)
}

// GetStoreByID implements
func (repo *StoreMockRepositoryImpl) GetStoreByID(id string) (*domain.Store, error) {
	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			return elem, nil
		}
	}
	return nil, errors.New(ErrorNotFoundStore)
}

// GetStore implements
func (repo *StoreMockRepositoryImpl) GetStore(name string) (*domain.Store, error) {
	for _, elem := range repo.mockStore.aStore {
		if elem.Name == name {
			return elem, nil
		}
	}

	return nil, errors.New(ErrorNotFoundStore)
}

// AddConsumer implements
func (repo *StoreMockRepositoryImpl) AddConsumer(id string, consumer *domain.Consumer) error {

	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			elem.Queue = append(elem.Queue, consumer)

			return nil
		}
	}

	return errors.New(ErrorNotFoundStore)
}

// RemoveConsumer implements
func (repo *StoreMockRepositoryImpl) RemoveConsumer(id string, phone string) error {

	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			for i, cons := range elem.Queue {
				if cons.Number == phone {
					copy(elem.Queue[i:], elem.Queue[i+1:])
					elem.Queue[len(elem.Queue)-1] = nil
					elem.Queue = elem.Queue[:len(elem.Queue)-1]

					return nil
				}
			}
		}
	}

	return errors.New(ErrorNotFoundConsumer)
}

// GetConsumer implements
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

	return nil, errors.New(ErrorNotFoundConsumer)
}

// GetAllConsumers implements
func (repo *StoreMockRepositoryImpl) GetAllConsumers(id string) ([]*domain.Consumer, error) {

	for _, elem := range repo.mockStore.aStore {
		if elem.ID == id {
			return elem.Queue, nil
		}
	}

	return nil, errors.New(ErrorNotFoundStore)
}
