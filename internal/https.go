package internal

//HTTPClient a synchronous or async http client wrapper
type HTTPClient struct {
	a Auth
}

func new(auth Auth) *HTTPClient {
	return &HTTPClient{
		auth,
	}
}
func (h HTTPClient) Authenticate() {
	//h.a.Authenticate()
}
