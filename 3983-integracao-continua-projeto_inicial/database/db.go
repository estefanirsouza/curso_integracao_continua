package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConectaComBancoDeDados() {
	var err error
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")
	stringDeConexao := "host=" + endereco + " user=" + usuario + " password=" + senha + " dbname=" + nomeBanco + " port=" + portaBanco + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
    log.Fatal(err)
}
}
