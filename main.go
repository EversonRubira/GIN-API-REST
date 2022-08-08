package main

import (
	routes "github.com/EversonRubira/api-go-gin/Routes"
	database "github.com/EversonRubira/api-go-gin/database"
)

func main() {

	database.DB.connectDatabase()
	routes.HandleRequest()

}
