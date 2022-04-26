package router

import (
	"API/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar vai retornar um router com as totas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
