package model


type Tenant struct {
	propertyName string 	`json: "propertyName"`
	tenantName   string		`json: "tenantName"`
	tenantPh     string		`json: "tenantPh"`
	tenantEmail  string		`json: "tenantEmail"`
}
