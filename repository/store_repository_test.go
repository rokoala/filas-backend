package repository

import (
	"errors"
	"testing"

	"github.com/rokoga/filas-backend/domain"
	"github.com/rokoga/filas-backend/infra"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	dbClient, dbCollection, err := infra.GetConnection("../config/tests/.env")
	if err != nil {
		panic(err)
	}
	defer infra.CloseConnection(dbClient)

	store := NewStoreRepository(dbCollection)

	newStore := domain.Store{
		Name:    "Test Store",
		URLName: "test",
		Queue:   nil,
	}
	store.Create(&newStore)

	tests := []struct {
		urlName    string
		name       string
		resultURL  string
		resultName string
		err        error
	}{
		{urlName: "outback", name: "Outback", resultURL: "outback", resultName: "Outback", err: nil},
		{urlName: "jeronimo", name: "Jeronimo", resultURL: "jeronimo", resultName: "Jeronimo", err: nil},
		{urlName: "", name: "", resultURL: "", resultName: "", err: errors.New(ErrorNotFoundStore)},
	}

	for _, test := range tests {
		newStore := domain.Store{
			Name:    test.name,
			URLName: test.urlName,
			Queue:   nil,
		}
		result, err := store.Create(&newStore)
		if err == nil {
			assert.NotNil(t, result)
			assert.Equal(t, test.resultURL, result.URLName)
			assert.Equal(t, test.resultName, result.Name)
			assert.NotEmpty(t, result.ID)
		} else {
			assert.Equal(t, test.err, err)
			assert.Nil(t, store)
		}
	}

}
