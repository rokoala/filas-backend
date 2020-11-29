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

	router.GET("/getstore/:name", func(c *gin.Context) {
		name := c.Param("name")

		domainStore, err := svc.GetStore(name)
		if err != nil {
			c.Error(err)
		}

		c.JSON(200, domainStore)
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

	router.DELETE("/removeconsumer/:storeid/:number", func(c *gin.Context) {
		storeid := c.Param("storeid")
		number := c.Param("number")

		err := svc.RemoveConsumer(storeid, number)
		if err != nil {
			c.Error(err)
		}

		c.JSON(200, nil)
	})

	router.GET("/getconsumer/:storeid/:number", func(c *gin.Context) {
		storeid := c.Param("storeid")
		number := c.Param("number")

		consumer, err := svc.GetConsumer(storeid, number)
		if err != nil {
			c.Error(err)
		}

		c.JSON(200, consumer)
	})

	router.GET("/getallconsumers/:storeid", func(c *gin.Context) {
		storeid := c.Param("storeid")

		allConsumers, err := svc.GetAllConsumers(storeid)
		if err != nil {
			c.Error(err)
		}

		c.JSON(200, allConsumers)
	})

	fmt.Printf("Server is listening at %s", PORT)
	router.Run(PORT)

	done <- "Server shutdown"
}
