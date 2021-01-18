package service

import "github.com/rokoga/filas-backend/domain"

// StoreService - Provides a Store services layer
type StoreService interface {
	Create(URLname, name string) (*domain.Store, error)
	RemoveStore(id string) error
	GetStore(name string) (*domain.Store, error)
	GetStoreByID(id string) (*domain.Store, error)
	AddConsumer(id, name, phone string) (string, error)
	RemoveConsumer(id string, phone string) error
	GetConsumer(id string, phone string) (*domain.Consumer, error)
	GetAllConsumers(id string) ([]*domain.Consumer, error)
}
