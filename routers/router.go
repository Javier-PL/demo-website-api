package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetRoutesDBphotos(router)
	router = SetRoutesDBcategories(router)

	return router
}
