package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rokoga/filas-backend/service"
	"github.com/rokoga/filas-backend/vo"
)

func Run(done chan string) {
	const PORT = ":8080"
	router := gin.Default()

	svc := service.NewStoreServiceImpl()

	router.PUT("/create", func(c *gin.Context) {
		createRequest := vo.CreateRequest{}
		c.BindJSON(&createRequest)

		store, err := svc.Create(createRequest.URLName, createRequest.Name)
		if err != nil {
			c.Error(err)
		}

		c.JSON(200, store)
	})

	router.PUT("/addconsumer", func(c *gin.Context) {
		addConsumerRequest := vo.AddConsumerRequest{}
		c.BindJSON(&addConsumerRequest)

		accessURL, err := svc.AddConsumer(addConsumerRequest.StoreID, addConsumerRequest.Name, addConsumerRequest.Phone)
		if err != nil {
			c.Error(err)
		}

		c.JSON(200, accessURL)
	})

	fmt.Printf("Server is listening at %s", PORT)
	router.Run(PORT)

	done <- "Server shutdown"
}
