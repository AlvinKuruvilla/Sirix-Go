package internal

//Resource Struct for manipulating a resource
type Resource struct {
	dbName       string
	resourceName string
	auth         Auth
	client       HTTPClient
	DBT          DBType
}
