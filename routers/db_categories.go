package routers

import (
	"clinicacl/ccl-website-api/models"
	"clinicacl/ccl-website-api/services"

	"github.com/gorilla/mux"
)

var routesDBcategories = []models.Route{
	{Path: "/category/c", Function: services.PostCategory, Method: "POST", Mw: "auth"},
	{Path: "/category/u", Function: services.UpdateCategory, Method: "PUT", Mw: "auth"},
	{Path: "/category/g", Function: services.GetCategory, Method: "POST", Mw: "auth"},
	{Path: "/categories/g", Function: services.GetCategories, Method: "GET", Mw: "auth"},
	{Path: "/category/d", Function: services.DeleteCategory, Method: "DELETE", Mw: "auth"},
}

func SetRoutesDBcategories(router *mux.Router) *mux.Router {
	for _, r := range routesDBcategories {

		if r.Mw == "" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
		} else if r.Mw == "auth" {
			router.HandleFunc(r.Path, r.Function).Methods(r.Method)
			//router.HandleFunc(r.Path, auth.RequireTokenAuthentication(r.Function)).Methods(r.Method)
		}

	}

	return router
}

func GetRoutesDBcategories() []models.Route {
	return routesDBcategories
}
