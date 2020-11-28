package repository

import "github.com/rokoga/filas-backend/domain"

type StoreRepositoryImpl struct {
	// Conn mongo
}

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (repo *StoreRepositoryImpl) Create(store *domain.Store) (*domain.Store, error) {
	return nil, nil
}

func (repo *StoreRepositoryImpl) Get(id string) (*domain.Store, error) {
	return nil, nil
}

func (repo *StoreRepositoryImpl) GetStore(URLname string) (*domain.Store, error) {
	return nil, nil
}

func (repo *StoreRepositoryImpl) AddConsumer(id string, consumer *domain.Consumer) error {
	return nil
}

func (repo *StoreRepositoryImpl) RemoveConsumer(id string, phone string) error {
	return nil
}

func (repo *StoreRepositoryImpl) GetConsumer(id string, phone string) (*domain.Consumer, error) {
	return nil, nil
}

func (repo *StoreRepositoryImpl) GetAllConsumers(id string) ([]*domain.Consumer, error) {
	return nil, nil
}
