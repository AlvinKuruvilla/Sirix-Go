package internal

//Resource Struct for manipulating a resource
type Resource struct {
	//todo: include DBType in here somehow...
	dbName       string
	resourceName string
	auth         Auth
	client       HTTPClient
}
