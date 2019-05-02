package model


type Tenant struct {
	PropertyName string 	`json: "propertyName"`
	TenantName   string		`json: "tenantName"`
	TenantPh     string		`json: "tenantPh"`
	TenantEmail  string		`json: "tenantEmail"`
}
