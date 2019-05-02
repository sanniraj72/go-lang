package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tenant-management/model"
)

type Property struct {
	OwnerEmail    string 	`json: "ownerEmail"`
	PropertyName  string 	`json: "propertyName"`
	AvailableFlat int 		`json: "availableFlat"`
	OccupiedFlat  int 		`json: "occupiedFlat"`
}

var ownerList = make(map[string]model.Owner)
var propertyList = make(map[string]Property)
//var tenantList = make(map[string]Tenant)

func main() {

	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:8080")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/add/property", addProperty)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var owner model.Owner
		errM := json.Unmarshal(data, &owner)

		// save data
		ownerList[owner.OwnerEmail] = owner

		if errM != nil {
			http.Error(w, errM.Error(), 500)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "Welcome %s!", owner.OwnerName)
	}

}

func addProperty(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var property Property
		errM := json.Unmarshal(data, &property)

		if _, ok := ownerList[property.OwnerEmail]; !ok {
			_, _ = fmt.Fprintf(w, "%s not registered as an owner", property.OwnerEmail)
			return
		}

		if _, ok := propertyList[property.PropertyName]; ok {
			_, _ = fmt.Fprintf(w, "Property \"%s\" already exist for %s", property.PropertyName, ownerList[property.OwnerEmail].OwnerName)
			return
		}

		// save data
		propertyList[property.PropertyName] = property

		if errM != nil {
			http.Error(w, errM.Error(), 500)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "Your property added successfully")
	}
}
