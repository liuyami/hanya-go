package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liuyami/hanya-go/bootstrap"
)

func main() {
	router := gin.New()

	bootstrap.SetRoute(router)

	err := router.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
