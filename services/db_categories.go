package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"clinicacl/ccl-website-api/models"
	"encoding/json"
	"net/http"
)

var categories_database = "ccl"
var categories_collection = "ccl.photos.categories"

func PostCategory(w http.ResponseWriter, r *http.Request) {

	var item models.Category

	itemmodel := mymodel{I: item, DatabaseName: categories_database, Collection: categories_collection}
	res, err := itemmodel.postInterface(w, r)
	if err != nil {
		Log.Error(err)
		return
	}

	Log.Info("POST", "PATIENT", "ID:", res.InsertedID.(primitive.ObjectID).Hex(), "Title:", "", "SUCCESS:", IsObjectIDValid(res.InsertedID.(primitive.ObjectID)))

}

func GetCategory(w http.ResponseWriter, r *http.Request) {

	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)

	var filter interface{}

	filter = bson.M{"_id": item.ID}

	itemmodel := mymodel{I: &item, F: filter, DatabaseName: categories_database, Collection: categories_collection}

	itemmodel.getInterface(w, r)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	var item models.Category
	//_ = json.NewDecoder(r.Body).Decode(&item)

	var filter interface{}
	filter = bson.M{}

	itemmodel := mymodel{I: &item, F: filter, DatabaseName: categories_database, Collection: categories_collection}

	itemmodel.getInterfaces(w, r)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	keys, _ := r.URL.Query()["byID"]
	key := keys[0]

	var update models.Category
	_ = json.NewDecoder(r.Body).Decode(&update)

	var filter interface{}
	if key != "" {
		idPrimitive, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			Log.Error(err)
			http.Error(w, err.Error(), 500)
			return
		}
		filter = bson.M{"_id": idPrimitive}
	} else {
		filter = bson.M{}

	}

	itemmodel := mymodel{I: update, F: filter, DatabaseName: categories_database, Collection: categories_collection}

	res, err := itemmodel.putInterface(w, r)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	Log.Info("UPDATE", "CATEGORY", "ID:", update.ID, "SUCCESS:", res.ModifiedCount)
	//Log.Info("UPDATE", "INVOICE", "ID:", items[0].ID, "SUCCESS:", IsObjectIDValid(res.UpsertedID.(primitive.ObjectID)))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	categorykeys, _ := r.URL.Query()["byID"]
	categorykey := categorykeys[0]

	var filter interface{}
	if categorykey != "" {
		idPrimitive, err := primitive.ObjectIDFromHex(categorykey)
		if err != nil {
			Log.Error(err)
			http.Error(w, err.Error(), 500)
			return
		}
		filter = bson.M{"_id": idPrimitive}
	} else {
		filter = bson.M{}

	}

	itemmodel := mymodel{F: filter, DatabaseName: categories_database, Collection: categories_collection}
	res, err := itemmodel.deleteInterface(w, r)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	Log.Info("DELETE", "ID:", categorykey, "SUCCESS:", res.DeletedCount)
}
