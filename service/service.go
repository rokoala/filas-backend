package service

import "github.com/rokoga/filas-backend/domain"

type StoreService interface {
	Create(URLname, name string) (*domain.Store, error)
	RemoveStore(id string) error
	GetStore(name string) (*domain.Store, error)
	AddConsumer(id, name, number string) (string, error)
	RemoveConsumer(id string, phone string) error
	GetConsumer(id string, phone string) (*domain.Consumer, error)
	GetAllConsumers(id string) ([]*domain.Consumer, error)
}
