package repository

import (
	"github.com/rokoga/filas-backend/domain"
)

// StoreRepository - Repositório para persistencia de um estabelecimento
type StoreRepository interface {
	Create(store *domain.Store) (*domain.Store, error)
	Get(id string) (*domain.Store, error)
	GetStore(URLname string) (*domain.Store, error)
	AddConsumer(id string, consumer *domain.Consumer) error
	RemoveConsumer(id string, phone string) error
	GetConsumer(id string, phone string) (*domain.Consumer, error)
	GetAllConsumers(id string) ([]*domain.Consumer, error)
}

// app.filas/outback/token?=24238971alkajrealm
