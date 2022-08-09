package database

import (
	"fmt"
	"log"

	"github.com/EversonRubira/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	stringDeConexao := "host=150.136.55.154 user=root password=root dbname=dev_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	fmt.Println("Conexao efetuada com sucesso.")
	DB.AutoMigrate(&models.Student{})
}
