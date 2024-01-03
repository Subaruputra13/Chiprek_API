package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// type MenuResponse struct {
// 	ID     int    `json:"id"`
// 	Nama   string `json:"nama"`
// 	Nasi   string `json:"nasi"`
// 	Harga  int    `json:"harga"`
// 	Jumlah int    `json:"jumlah"`
// 	Cabe   int    `json:"cabe"`
// }
