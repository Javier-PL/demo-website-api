package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"clinicacl/ccl-website-api/models"
	"encoding/json"
	"net/http"
)

var photos_database = "ccl"
var photos_collection = "ccl.photos"

func PostPhoto(w http.ResponseWriter, r *http.Request) {

	var item models.Photo

	itemmodel := mymodel{I: item, DatabaseName: photos_database, Collection: photos_collection}
	res, err := itemmodel.postInterface(w, r)
	if err != nil {
		Log.Error(err)
		return
	}

	Log.Info("POST", "PATIENT", "ID:", res.InsertedID.(primitive.ObjectID).Hex(), "Title:", "", "SUCCESS:", IsObjectIDValid(res.InsertedID.(primitive.ObjectID)))

}

func GetPhoto(w http.ResponseWriter, r *http.Request) {

	var item models.Photo
	_ = json.NewDecoder(r.Body).Decode(&item)

	var filter interface{}

	filter = bson.M{"_id": item.ID}

	itemmodel := mymodel{I: &item, F: filter, DatabaseName: photos_database, Collection: photos_collection}

	itemmodel.getInterface(w, r)
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {

	categorykeys, _ := r.URL.Query()["bycategoryID"]
	categorykey := categorykeys[0]

	var item models.Photo
	//_ = json.NewDecoder(r.Body).Decode(&item)

	Log.Info(categorykey)
	var filter interface{}
	if categorykey != "" {
		filter = bson.M{"categoryID": categorykey}
	} else {
		filter = bson.M{}

	}

	itemmodel := mymodel{I: &item, F: filter, DatabaseName: photos_database, Collection: photos_collection}

	itemmodel.getInterfaces(w, r)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {

	var items []models.Photo //awaits 2 items, the item to update, and the update
	_ = json.NewDecoder(r.Body).Decode(&items)

	if len(items) < 2 {
		Log.Error(items)
		http.Error(w, "Expecting 2 structures, the original and the update", 500)
		return
	}
	filter := bson.M{"_id": items[0].ID}

	itemmodel := mymodel{I: items[1], F: filter, DatabaseName: photos_database, Collection: photos_collection}

	res, err := itemmodel.putInterface(w, r)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	Log.Info("UPDATE", "PHOTO", "ID:", items[0].ID, "SUCCESS:", res.ModifiedCount)
	//Log.Info("UPDATE", "INVOICE", "ID:", items[0].ID, "SUCCESS:", IsObjectIDValid(res.UpsertedID.(primitive.ObjectID)))
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {

	keys, _ := r.URL.Query()["byID"]
	key := keys[0]

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

	itemmodel := mymodel{F: filter, DatabaseName: photos_database, Collection: photos_collection}
	res, err := itemmodel.deleteInterface(w, r)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	Log.Info("DELETE", "ID:", key, "SUCCESS:", res.DeletedCount)
}
