package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GiovaniGitHub/cpf-weather/configs"
	_ "github.com/GiovaniGitHub/cpf-weather/docs"
	"github.com/GiovaniGitHub/cpf-weather/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Desafio 2.0 - cpf-weather
// @version         1.0
// @description     Fullcycle Pós Go Expert Go Expert

// @termsOfService  http://swagger.io/terms/

// @contact.name   Giovani Angelo
// @contact.email  giovani.angelo@gmail.com

// @host      localhost:8000
// @BasePath  /
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", handlers.GetTemperature)
	})

	// Inicia o servidor
	apiURL := "http://localhost:" + configs.WebServerPort + "/cep"
	log.Printf("API está disponível em: %s", apiURL)

	// Inicia o Swagger
	log.Printf("API Swagger está disponível em: %s", "http://localhost:"+configs.WebServerPort+"/docs/index.html")

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:"+configs.WebServerPort+"/docs/doc.json")))
	http.ListenAndServe(fmt.Sprintf(":%s", configs.WebServerPort), r)
}
