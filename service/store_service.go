package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rokoga/filas-backend/domain"
	"github.com/rokoga/filas-backend/repository"
)

const (
	// ErrorArgumentNotValidGetStore for invalid argument
	ErrorArgumentNotValidGetStore = "Os parametros para pesquisa do estabelecimento devem ser preenchidos"
	// ErrorArgumentNotValidAddStore for invalid argument
	ErrorArgumentNotValidAddStore = "Os parametros para inserção do estabelecimento devem ser preenchidos"
	// ErrorArgumentNotValidRemoveStore for invalid argument
	ErrorArgumentNotValidRemoveStore = "Os parametros para remoção do estabelecimento devem ser preenchidos"
	// ErrorArgumentNotValidAddConsumer for invalid argument
	ErrorArgumentNotValidAddConsumer = "Os parametros para inserção de consumidor devem ser preenchidos"
	// ErrorArgumentNotValidRemoveConsumer for invalid argument
	ErrorArgumentNotValidRemoveConsumer = "Os parametros para remoção de consumidor devem ser preenchidos"
	// ErrorArgumentNotValidGetConsumer for invalid argument
	ErrorArgumentNotValidGetConsumer = "Os parametros para pesquisa de consumidor devem ser preenchidos"
	// ErrorStoreExists for already created store
	ErrorStoreExists = "Estabelecimento com nome já cadastrado"
)

// StoreServiceImpl implements
type StoreServiceImpl struct {
	storeRepository repository.StoreRepository
}

// NewStoreServiceImpl implements
func NewStoreServiceImpl(db *mongo.Collection) StoreService {
	return &StoreServiceImpl{
		storeRepository: repository.NewStoreRepository(db),
	}
}

// Create implements
func (svc *StoreServiceImpl) Create(URLname, name string) (*domain.Store, error) {

	if URLname == "" || name == "" {
		return nil, errors.New(ErrorArgumentNotValidAddStore)
	}

	lstore, err := svc.storeRepository.GetStore(name)
	if err != nil {
		if err.Error() != repository.ErrorNotFoundStore {
			return nil, err
		}
	}

	if lstore != nil {
		return nil, errors.New(ErrorStoreExists)
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
func (svc *StoreServiceImpl) RemoveStore(id string) error {

	if id == "" {
		return errors.New(ErrorArgumentNotValidRemoveStore)
	}

	err := svc.storeRepository.RemoveStore(id)
	if err != nil {
		return err
	}

	return nil
}

// GetStore implements
func (svc *StoreServiceImpl) GetStore(name string) (*domain.Store, error) {

	if name == "" {
		return nil, errors.New(ErrorArgumentNotValidGetStore)
	}

	store, err := svc.storeRepository.GetStore(name)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// GetStoreByID implements
func (svc *StoreServiceImpl) GetStoreByID(id string) (*domain.Store, error) {

	if id == "" {
		return nil, errors.New(ErrorArgumentNotValidGetStore)
	}

	store, err := svc.storeRepository.GetStoreByID(id)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// AddConsumer implements
func (svc *StoreServiceImpl) AddConsumer(id, name, phone string) (string, error) {

	if id == "" || name == "" || phone == "" {
		return "", errors.New(ErrorArgumentNotValidAddConsumer)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	consumer := domain.Consumer{
		Name:      name,
		Phone:     phone,
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
func (svc *StoreServiceImpl) RemoveConsumer(id, phone string) error {

	if id == "" || phone == "" {
		return errors.New(ErrorArgumentNotValidRemoveConsumer)
	}

	if err := svc.storeRepository.RemoveConsumer(id, phone); err != nil {
		return err
	}

	return nil
}

// GetConsumer implements
func (svc *StoreServiceImpl) GetConsumer(id, phone string) (*domain.Consumer, error) {

	if id == "" || phone == "" {
		return nil, errors.New(ErrorArgumentNotValidGetConsumer)
	}

	consumer, err := svc.storeRepository.GetConsumer(id, phone)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

// GetAllConsumers implements
func (svc *StoreServiceImpl) GetAllConsumers(id string) ([]*domain.Consumer, error) {

	if id == "" {
		return nil, errors.New(ErrorArgumentNotValidGetConsumer)
	}

	consumers, err := svc.storeRepository.GetAllConsumers(id)
	if err != nil {
		return nil, err
	}

	return consumers, nil
}
