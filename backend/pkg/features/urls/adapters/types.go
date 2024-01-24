package adapters

type OriginalUrl struct {
	Ou *string
}

type AuthUserInfo struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
