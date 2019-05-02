package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tenant-management/model"
)

var ownerList = make(map[string]model.Owner)
var propertyList = make(map[string]model.Property)
var tenantList = make(map[string]model.Tenant)

func main() {

	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:8080")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/add/property", addProperty)
	http.HandleFunc("/add/tenant", addTenant)
	http.HandleFunc("/get/property", getProperty)
	http.HandleFunc("/get/tenant", getTenant)
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
		var property model.Property
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

func addTenant(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var tenant model.Tenant
		errM := json.Unmarshal(data, &tenant)

		if _, ok := propertyList[tenant.PropertyName]; !ok {
			_, _ = fmt.Fprintf(w, "Property \"%s\" does not exist", tenant.PropertyName)
			return
		}

		if _, ok := tenantList[tenant.TenantEmail]; ok {
			_, _ = fmt.Fprintf(w, "Tenant \"%s\" already exist", tenant.TenantName)
			return
		}

		// save data
		tenantList[tenant.TenantEmail] = tenant

		if errM != nil {
			http.Error(w, errM.Error(), 500)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "Tenant added successfully to property %s", tenant.PropertyName)

	}
}

func getProperty(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "GET" {

		emails, ok := r.URL.Query()["ownerEmail"]

		if !ok || len(emails) < 1 {
			_, _ = fmt.Fprintf(w, "Please provide correct email id")
			return
		}

		l := 0
		for _, property := range propertyList{
			if property.OwnerEmail == emails[0] {
				l++
			}
		}

		var pList = make([]model.Property, l)
		var i = 0
		for _, property := range propertyList{
			if property.OwnerEmail == emails[0] {
				pList[i] = property
				i++
			}
		}

		if len(pList) == 0 {
			_, _ = fmt.Fprintf(w, "You don't have any property")
			return
		}

		b, err := json.Marshal(pList)

		if err != nil {
			_, _ = fmt.Fprintf(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(b)
	}
}

func getTenant(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "GET" {

		properties, ok := r.URL.Query()["propertyName"]

		if !ok || len(properties) < 1 {
			_, _ = fmt.Fprintf(w, "Please provide correct property name")
			return
		}

		l := 0
		for _, tenant := range tenantList{
			if tenant.PropertyName == properties[0] {
				l++
			}
		}

		var tList = make([]model.Tenant, l)
		var i = 0
		for _, tenant := range tenantList{
			if tenant.PropertyName == properties[0] {
				tList[i] = tenant
				i++
			}
		}

		if len(tList) == 0 {
			_, _ = fmt.Fprintf(w, "You don't have any tenant in %s", properties[0])
			return
		}

		b, err := json.Marshal(tList)

		if err != nil {
			_, _ = fmt.Fprintf(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(b)
	}
}