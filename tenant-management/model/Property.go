package model


type Property struct {
	ownerEmail    string 	`json: "ownerEmail"`
	propertyName  string 	`json: "propertyName"`
	availableFlat int 		`json: "availableFlat"`
	occupiedFlat  int 		`json: "occupiedFlat"`
}