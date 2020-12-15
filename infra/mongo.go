package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/rokoga/filas-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetConnection implements MongoDB connection
func GetConnection(configFile string) (*mongo.Client, *mongo.Collection, error) {
	cfg, err := config.ReadConfig(configFile, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo de configuração: %v", err)
	}

	port := cfg.GetString("dbport")
	dbhost := cfg.GetString("dbhost")
	dbdriver := cfg.GetString("dbdriver")
	// dbuser := cfg.GetString("dbuser")
	// dbpass := cfg.GetString("dbpass")
	database := cfg.GetString("dbname")
	dbcollection := cfg.GetString("dbcollection")
	uri := fmt.Sprintf("%s://%s:%s", dbdriver, dbhost, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao criar cliente de conexão com o banco: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao criar conexão com o banco: %v", err)
	}

	collection := client.Database(database).Collection(dbcollection)

	return client, collection, nil
}

// CloseConnection implements database client disconnect
func CloseConnection(dbClient *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := dbClient.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("Erro ao fechar conexão com o banco: %v", err)
	}

	return nil
}
