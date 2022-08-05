package main

import (
	routes "github.com/EversonRubira/api-go-gin/Routes"
	"github.com/EversonRubira/api-go-gin/database"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
