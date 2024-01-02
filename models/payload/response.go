package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
