package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/rokoga/filas-backend/domain"
	"github.com/rokoga/filas-backend/repository"
)

// StoreMockServiceImpl implements
type StoreMockServiceImpl struct {
	storeRepository repository.StoreRepository
}

// NewStoreMockServiceImpl implements
func NewStoreMockServiceImpl() StoreService {
	return &StoreMockServiceImpl{
		storeRepository: repository.NewStoreMockRepository(),
	}
}

// Create implements
func (svc *StoreMockServiceImpl) Create(URLname, name string) (*domain.Store, error) {

	if URLname == "" || name == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_ADD_STORE)
	}

	lstore, err := svc.storeRepository.GetStore(name)
	if err != nil {
		if err.Error() != repository.ErrorNotFoundStore {
			return nil, err
		}
	}

	if lstore != nil {
		return nil, errors.New(ERROR_STORE_EXISTS)
	}

	store := domain.Store{
		Name:    name,
		URLName: URLname,
	}

	newStore, err := svc.storeRepository.Create(&store)
	if err != nil {
		return nil, err
	}

	return newStore, nil
}

// RemoveStore implements
func (svc *StoreMockServiceImpl) RemoveStore(id string) error {

	if id == "" {
		return errors.New(ERROR_ARGUMENT_NOT_VALID_REMOVE_STORE)
	}

	err := svc.storeRepository.RemoveStore(id)
	if err != nil {
		return err
	}

	return nil
}

// GetStore implements
func (svc *StoreMockServiceImpl) GetStore(name string) (*domain.Store, error) {

	if name == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_STORE)
	}

	store, err := svc.storeRepository.GetStore(name)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// GetStoreByID implements
func (svc *StoreMockServiceImpl) GetStoreByID(id string) (*domain.Store, error) {

	if id == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_STORE)
	}

	store, err := svc.storeRepository.GetStoreByID(id)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// AddConsumer implements
func (svc *StoreMockServiceImpl) AddConsumer(id, name, number string) (string, error) {

	if id == "" || name == "" || number == "" {
		return "", errors.New(ERROR_ARGUMENT_NOT_VALID_ADD_CONSUMER)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	consumer := domain.Consumer{
		Name:      name,
		Number:    number,
		Accesskey: strconv.Itoa(r1.Int()),
	}

	if err := svc.storeRepository.AddConsumer(id, &consumer); err != nil {
		return "", err
	}

	store, err := svc.storeRepository.GetStoreByID(id)
	if err != nil {
		return "", err
	}

	accessConsumerURL := fmt.Sprintf("%s/%s", store.URLName, consumer.Accesskey)

	return accessConsumerURL, nil
}

// RemoveConsumer implements
func (svc *StoreMockServiceImpl) RemoveConsumer(id, phone string) error {

	if id == "" || phone == "" {
		return errors.New(ERROR_ARGUMENT_NOT_VALID_REMOVE_CONSUMER)
	}

	if err := svc.storeRepository.RemoveConsumer(id, phone); err != nil {
		return err
	}

	return nil
}

// GetConsumer implements
func (svc *StoreMockServiceImpl) GetConsumer(id, phone string) (*domain.Consumer, error) {

	if id == "" || phone == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_CONSUMER)
	}

	consumer, err := svc.storeRepository.GetConsumer(id, phone)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

// GetAllConsumers implements
func (svc *StoreMockServiceImpl) GetAllConsumers(id string) ([]*domain.Consumer, error) {

	if id == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_CONSUMER)
	}

	consumers, err := svc.storeRepository.GetAllConsumers(id)
	if err != nil {
		return nil, err
	}

	return consumers, nil
}
