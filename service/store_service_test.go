package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	svc := NewStoreServiceImpl()

	tests := []struct {
		urlName    string
		name       string
		resultURL  string
		resultName string
		err        error
	}{
		{urlName: "outback", name: "Outback", resultURL: "outback", resultName: "Outback", err: nil},
		{urlName: "jeronimo", name: "Jeronimo", resultURL: "jeronimo", resultName: "Jeronimo", err: nil},
		{urlName: "", name: "", resultURL: "", resultName: "", err: errors.New(ERROR_ARGUMENT_NOT_VALID_ADD_STORE)},
	}

	for _, test := range tests {
		store, err := svc.Create(test.urlName, test.name)
		if err == nil {
			assert.NotNil(t, store)
			assert.Equal(t, test.resultURL, store.URLName)
			assert.Equal(t, test.resultName, store.Name)
			assert.NotEmpty(t, store.ID)
		} else {
			assert.IsType(t, test.err, err)
			assert.Nil(t, store)
		}
	}

	// fmt.Printf("Store: %v", store)
}

func TestRemoveStore(t *testing.T) {

	svc := NewStoreServiceImpl()

	URLname := "outback"
	name := "Outback"

	store, err := svc.Create(URLname, name)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	tests := []struct {
		id  string
		err error
	}{
		{id: store.ID, err: nil},
		{id: store.ID, err: errors.New("ERROR_REMOVE_CONSUMER")},
		{id: "fakeID", err: errors.New("ERROR_REMOVE_CONSUMER")},
		{id: "", err: errors.New(ERROR_ARGUMENT_NOT_VALID_REMOVE_STORE)},
	}

	for _, test := range tests {
		err := svc.RemoveStore(test.id)
		assert.IsType(t, test.err, err)
	}

}

func TestGetStore(t *testing.T) {

	svc := NewStoreServiceImpl()

	URLname := "outback"
	name := "Outback"

	store, err := svc.Create(URLname, name)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	tests := []struct {
		name       string
		resultURL  string
		resultName string
		err        error
	}{
		{name: "Outback", resultURL: "outback", resultName: "Outback", err: nil},
		{name: "Jeronimo", resultURL: "", resultName: "", err: errors.New("Não foi encontrado o estabelecimento")},
		{name: "", resultURL: "", resultName: "", err: errors.New("Parametro de url não deve ser vazio")},
	}

	for _, test := range tests {
		store, err := svc.GetStore(test.name)
		if err == nil {
			assert.NotNil(t, store)
			assert.Equal(t, test.resultURL, store.URLName)
			assert.Equal(t, test.resultName, store.Name)
			assert.NotEmpty(t, store.ID)
		} else {
			assert.IsType(t, test.err, err)
			assert.Nil(t, store)
		}
	}

}

func TestAddConsumer(t *testing.T) {

	svc := NewStoreServiceImpl()

	URLname := "outback"
	name := "Outback"

	store, err := svc.Create(URLname, name)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	tests := []struct {
		id    string
		name  string
		phone string
		err   error
	}{
		{id: store.ID, name: "Fulano", phone: "011998989898", err: nil},
		{id: store.ID, name: "Ciclano", phone: "011922222222", err: nil},
		{id: store.ID, name: "", phone: "", err: errors.New(ERROR_ARGUMENT_NOT_VALID_ADD)},
		{id: "", name: "Fulaninho", phone: "011888888888", err: errors.New(ERROR_ARGUMENT_NOT_VALID_ADD)},
		{id: "FakeID", name: "Fulaninho", phone: "011888888888", err: errors.New("Não foi encontrado o estabelecimento")},
	}

	for _, test := range tests {
		accessURL, err := svc.AddConsumer(test.id, test.name, test.phone)
		if err == nil {
			assert.NotNil(t, accessURL)
		} else {
			assert.IsType(t, test.err, err)
			assert.Empty(t, accessURL)
		}
	}

}

func TestRemoveConsumer(t *testing.T) {

	svc := NewStoreServiceImpl()

	URLname := "outback"
	name := "Outback"

	store, err := svc.Create(URLname, name)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	consumerName := "Fulano"
	consumerPhone := "011998989898"
	consumerFakePhone := "011988888888"

	accessConsumerURL, err2 := svc.AddConsumer(store.ID, consumerName, consumerPhone)

	assert.Nil(t, err2)
	assert.NotNil(t, accessConsumerURL)

	err3 := svc.RemoveConsumer(store.ID, consumerPhone)

	assert.Nil(t, err3)

	err4 := svc.RemoveConsumer(store.ID, consumerFakePhone)

	assert.NotNil(t, err4)
}

func TestGetConsumer(t *testing.T) {

	svc := NewStoreServiceImpl()

	URLname := "outback"
	name := "Outback"

	store, err := svc.Create(URLname, name)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	consumerName := "Fulano"
	consumerPhone := "011998989898"
	consumerFakePhone := "011988888888"

	accessConsumerURL, err2 := svc.AddConsumer(store.ID, consumerName, consumerPhone)

	assert.Nil(t, err2)
	assert.NotNil(t, accessConsumerURL)

	consumer, err3 := svc.GetConsumer(store.ID, consumerPhone)

	assert.Nil(t, err3)
	assert.NotNil(t, consumer)

	assert.NotEmpty(t, consumer.Name)
	assert.NotEmpty(t, consumer.Number)
	assert.NotEmpty(t, consumer.Accesskey)

	consumer, err4 := svc.GetConsumer(store.ID, consumerFakePhone)

	assert.NotNil(t, err4)
	assert.Nil(t, consumer)

}

func TestGetAllConsumers(t *testing.T) {

	svc := NewStoreServiceImpl()

	URLname := "outback"
	name := "Outback"

	store, err := svc.Create(URLname, name)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	c1Name := "Fulano Um"
	c1Phone := "011998989899"
	c2Name := "Fulano Dois"
	c2Phone := "011998989855"
	c3Name := "Fulano Tres"
	c3Phone := "011998989877"

	accessConsumerURL, err := svc.AddConsumer(store.ID, c1Name, c1Phone)

	assert.Nil(t, err)
	assert.NotNil(t, accessConsumerURL)

	accessConsumerURL1, err := svc.AddConsumer(store.ID, c2Name, c2Phone)

	assert.Nil(t, err)
	assert.NotNil(t, accessConsumerURL1)

	accessConsumerURL2, err := svc.AddConsumer(store.ID, c3Name, c3Phone)

	assert.Nil(t, err)
	assert.NotNil(t, accessConsumerURL2)

	consumers, err := svc.GetAllConsumers(store.ID)

	assert.Nil(t, err)
	assert.NotNil(t, consumers)

	// assert.NotEmpty(t, consumer.Name)
	// assert.NotEmpty(t, consumer.Number)
	// assert.NotEmpty(t, consumer.Accesskey)

	consumer, err4 := svc.GetAllConsumers("")

	assert.NotNil(t, err4)
	assert.Nil(t, consumer)

}
