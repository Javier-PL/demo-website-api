package routers

import (
	"clinicacl/ccl-website-api/models"
	"clinicacl/ccl-website-api/services"

	"github.com/gorilla/mux"
)

var routesDBphotos = []models.Route{
	{Path: "/photo/c", Function: services.PostPhoto, Method: "POST", Mw: "auth"},
	{Path: "/photo/u", Function: services.UpdatePhoto, Method: "POST", Mw: "auth"},
	{Path: "/photo/g", Function: services.GetPhoto, Method: "POST", Mw: "auth"},
	{Path: "/photos/g", Function: services.GetPhotos, Method: "GET", Mw: "auth"},
	{Path: "/photo/d", Function: services.DeletePhoto, Method: "DELETE", Mw: "auth"},
}

func SetRoutesDBphotos(router *mux.Router) *mux.Router {
	for _, r := range routesDBphotos {

		if r.Mw == "" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
		} else if r.Mw == "auth" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
			//router.HandleFunc(r.Path, auth.RequireTokenAuthentication(r.Function)).Methods(r.Method)
		}

	}

	return router
}

func GetRoutesDBphotos() []models.Route {
	return routesDBphotos
}
