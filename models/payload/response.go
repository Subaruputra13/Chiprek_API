package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type CreateCustomerResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
}

type DashboardAdminResponse struct {
	TotalMenus  uint `json:"total_menus"`
	TotalOrders uint `json:"total_orders"`
	TotalUsers  uint `json:"total_users"`
	TotalIncome uint `json:"total_income"`
}

// type MenuResponse struct {
// 	ID     int    `json:"id"`
// 	Nama   string `json:"nama"`
// 	Nasi   string `json:"nasi"`
// 	Harga  int    `json:"harga"`
// 	Jumlah int    `json:"jumlah"`
// 	Cabe   int    `json:"cabe"`
// }
