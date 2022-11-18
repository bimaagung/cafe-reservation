package web

type UserReq struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
}

type UserRes struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}