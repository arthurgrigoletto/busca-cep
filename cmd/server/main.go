package main

import (
	"net/http"

	_ "github.com/arthurgrigoletto/busca-cep/docs"
	"github.com/arthurgrigoletto/busca-cep/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Busca CEP
// @version         1.0
// @description     API para buscar cep em diversos providers e retornar o mais rápido
// @termsOfService  http://swagger.io/terms/

// @contact.name   Arthur Grigoletto
// @contact.url    http://www.swagger.io/support
// @contact.email  arthur.grigoletto.lima@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3333
// @BasePath  /
func main() {
	cepHandler := handlers.NewCepHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/search/{cep}", cepHandler.GetCepInfo)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3333/docs/doc.json")))

	http.ListenAndServe(":3333", r)
}
