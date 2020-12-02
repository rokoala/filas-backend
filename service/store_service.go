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

const (
	ERROR_ARGUMENT_NOT_VALID_GET_STORE       = "Os parametros para pesquisa do estabelecimento devem ser preenchidos"
	ERROR_ARGUMENT_NOT_VALID_ADD_STORE       = "Os parametros para inserção do estabelecimento devem ser preenchidos"
	ERROR_ARGUMENT_NOT_VALID_REMOVE_STORE    = "Os parametros para remoção do estabelecimento devem ser preenchidos"
	ERROR_ARGUMENT_NOT_VALID_ADD_CONSUMER    = "Os parametros para inserção de consumidor devem ser preenchidos"
	ERROR_ARGUMENT_NOT_VALID_REMOVE_CONSUMER = "Os parametros para remoção de consumidor devem ser preenchidos"
	ERROR_ARGUMENT_NOT_VALID_GET_CONSUMER    = "Os parametros para pesquisa de consumidor devem ser preenchidos"
	ERROR_STORE_EXISTS                       = "Estabelecimento com nome já cadastrado"
)

type StoreServiceImpl struct {
	storeRepository repository.StoreRepository
}

func NewStoreServiceImpl() StoreService {
	return &StoreServiceImpl{
		storeRepository: repository.NewStoreMockRepository(),
	}
}

func (svc *StoreServiceImpl) Create(URLname, name string) (*domain.Store, error) {

	if URLname == "" || name == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_ADD_STORE)
	}

	lstore, err := svc.storeRepository.GetStore(name)
	if err != nil {
		if err.Error() != repository.ERROR_NOT_FOUND {
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

func (svc *StoreServiceImpl) RemoveStore(id string) error {

	if id == "" {
		return errors.New(ERROR_ARGUMENT_NOT_VALID_REMOVE_STORE)
	}

	err := svc.storeRepository.RemoveStore(id)
	if err != nil {
		return err
	}

	return nil
}

func (svc *StoreServiceImpl) GetStore(name string) (*domain.Store, error) {

	if name == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_STORE)
	}

	store, err := svc.storeRepository.GetStore(name)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (svc *StoreServiceImpl) AddConsumer(id, name, number string) (string, error) {

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

	store, err := svc.storeRepository.Get(id)
	if err != nil {
		return "", err
	}

	accessConsumerURL := fmt.Sprintf("%s/%s", store.URLName, consumer.Accesskey)

	return accessConsumerURL, nil
}

func (svc *StoreServiceImpl) RemoveConsumer(id, phone string) error {

	if id == "" || phone == "" {
		return errors.New(ERROR_ARGUMENT_NOT_VALID_REMOVE_CONSUMER)
	}

	if err := svc.storeRepository.RemoveConsumer(id, phone); err != nil {
		return err
	}

	return nil
}

func (svc *StoreServiceImpl) GetConsumer(id, phone string) (*domain.Consumer, error) {

	if id == "" || phone == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_CONSUMER)
	}

	consumer, err := svc.storeRepository.GetConsumer(id, phone)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (svc *StoreServiceImpl) GetAllConsumers(id string) ([]*domain.Consumer, error) {

	if id == "" {
		return nil, errors.New(ERROR_ARGUMENT_NOT_VALID_GET_CONSUMER)
	}

	consumers, err := svc.storeRepository.GetAllConsumers(id)
	if err != nil {
		return nil, err
	}

	return consumers, nil
}
