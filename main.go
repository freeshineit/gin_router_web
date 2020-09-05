package main

import (
	"gin-router-web/router"
	"log"
)

func main() {
	r := router.SetupRouter()

	err := r.Run(":8089") // listen and serve on 0.0.0.0:8089

	if err == nil {
		log.Println("listen and serve on localhost:8089")
	}
}
