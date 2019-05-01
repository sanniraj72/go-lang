package main

import (
	"fmt"
	"log"
	"net/http"
)

type Owner struct {
	ownerName  string
	ownerEmail string
}

//type Property struct {
//	ownerEmail    string
//	propertyName  string
//	availableFlat int
//	occupiedFlat  int
//}
//
//type Tenant struct {
//	propertyName string
//	tenantName   string
//	tenantPh     string
//	tenantEmail  string
//}

var ownerList = make(map[string]Owner)
//var propertyList = make(map[string]Property)
//var tenantList = make(map[string]Tenant)

func main() {

	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:8080")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/signup", signup)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var ownerName string
		var ownerEmail string

		fmt.Println("Enter owner name:")
		_, _ = fmt.Scanf("%s", ownerName)

		fmt.Println("Enter owner email:")
		_, _ = fmt.Scanf("%s", ownerEmail)

		owner := Owner{
			ownerName:  ownerName,
			ownerEmail: ownerEmail,
		}

		ownerList[ownerEmail] = owner
		fmt.Printf("Welcome %v", ownerName)
	}
}
