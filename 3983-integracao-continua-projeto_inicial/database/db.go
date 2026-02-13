package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// função auxiliar para pegar variável de ambiente com valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ConectaComBancoDeDados() {
	var err error

	// pega variáveis com fallback
	endereco := getEnv("DB_HOST", "localhost")
	usuario := getEnv("DB_USER", "root")
	senha := getEnv("DB_PASSWORD", "root")
	nomeBanco := getEnv("DB_NAME", "root")
	portaBanco := getEnv("DB_PORT", "5432")

	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		endereco,
		usuario,
		senha,
		nomeBanco,
		portaBanco,
	)

	fmt.Println("Conectando ao banco em:", endereco)

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	// AutoMigrate
	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Fatal("Erro ao fazer migration: ", err)
	}

	fmt.Println("Banco conectado com sucesso 🚀")
}
