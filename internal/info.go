package internal

//TokeData Represents a connection token
type TokeData struct {
	accessToken      string
	expireAt         uint
	expiresIn        uint
	notBeforePolicy  uint
	refreshExpiresIn uint
	refreshToken     string
	scope            string
	sessionState     string
	tokenType        string
}

//TokePostDaa a struct containing thw sata in a post token
type TokePostDaa struct {
	username  string
	password  string
	grantType string
}
