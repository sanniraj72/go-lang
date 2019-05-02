package model


type Property struct {
	OwnerEmail    string 	`json: "ownerEmail"`
	PropertyName  string 	`json: "propertyName"`
	AvailableFlat int 		`json: "availableFlat"`
	OccupiedFlat  int 		`json: "occupiedFlat"`
}