package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type scheme int

const (
	//HTTP usng the HTTP protocol
	HTTP scheme = iota
	//HTTPS using the HTTPS protocol
	HTTPS
)

var schemeStr = [...]string{
	"HTTPS",
	"HTTP",
}

//Auth a struct representing the data needed for authentication
type Auth struct {
	username string
	password string
	client   http.Client
	url      string
}

//NewAuth creates a new Auth struct
func NewAuth(
	username string,
	password string,
	client http.Client,
	url string,
) Auth {
	a := Auth{
		username,
		password,
		client,
		url,
	}
	return a
}

//Authenticate an authentication method in the works
func (a Auth) Authenticate(jsonByteData []byte) *http.Response {
	reqBody := bytes.NewBuffer(jsonByteData)
	req, err := http.NewRequest("POST", a.url, reqBody)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp
}

//Parse Parses the token data and refreshes access token before expirary
func Parse(res http.Response) {
	//data = {k.replace("-", "_"): v for k, v in resp.json().items()} - what is this doing and how do i replicate it here ??

	// IDK what to do with this
	//	self._token_data = TokenData(**data)
	//	self._client.headers[
	//		"Authorization"
	//	] = f"{self._token_data.token_type} {self._token_data.access_token}"
	//	self._timer = Timer(self._token_data.expires_in - 10, self._refresh)
	//  self._timer = ensure_future(self._sleep_then_refresh())

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))
	var parsed []string
	if err := json.Unmarshal(body, &parsed); err != nil {
		log.Fatalf("JSON unmarshal: %s", err)
	}

}
