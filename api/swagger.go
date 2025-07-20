package api

import (
	"net/http"

	_ "github.com/dae-vercel-function/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title        Drink and Eat vercel function API
// @version      1.0
// @description  Drink and Eat API that hosted on vercel as microservice.

// @contact.name   Liem Tran
// @contact.email  liemtran1414@gmail.com

// @host      drinkandeat.vercel.app
// @BasePath  /api
func SwaggerHandler(w http.ResponseWriter, r *http.Request){
	httpSwagger.Handler(httpSwagger.URL("./doc.json"))(w, r)
}  