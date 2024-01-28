package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GiovaniGitHub/cep-weather/configs"
	_ "github.com/GiovaniGitHub/cep-weather/docs"
	"github.com/GiovaniGitHub/cep-weather/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Desafio 2.0 - cep-weather
// @version         1.0
// @description     Fullcycle Pós Go Expert Go Expert

// @termsOfService  http://swagger.io/terms/

// @contact.name   Giovani Angelo
// @contact.email  giovani.angelo@gmail.com

// @host      localhost:8080
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

	if configs.ENVIRONMENT == "dev" {
		apiURL := configs.URL_BASE + ":" + configs.WebServerPort + "/cep"
		log.Printf("API está disponível em: %s", apiURL)

		log.Printf("API Swagger está disponível em: %s", configs.URL_BASE+":"+configs.WebServerPort+"/docs/index.html")

		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(configs.URL_BASE+":"+configs.WebServerPort+"/docs/doc.json")))
		http.ListenAndServe(fmt.Sprintf(":%s", configs.WebServerPort), r)
	}
	if configs.ENVIRONMENT == "prod" {
		apiURL := configs.URL_BASE + "/cep"
		log.Printf("API está disponível em: %s", apiURL)
	}

}
