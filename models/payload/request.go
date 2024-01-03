package payload

type LoginAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type CreateMenuRequest struct {
	Nama   string `json:"nama" form:"nama" validate:"required"`
	Nasi   string `json:"nasi" form:"nasi"`
	Harga  int    `json:"harga" form:"harga" validate:"required"`
	Jumlah int    `json:"jumlah" form:"jumlah"`
	Cabe   int    `json:"cabe" form:"cabe"`
}
type UpdateMenuRequest struct {
	Nama   string `json:"nama" form:"nama" validate:"required"`
	Nasi   string `json:"nasi" form:"nasi"`
	Harga  int    `json:"harga" form:"harga" validate:"required"`
	Jumlah int    `json:"jumlah" form:"jumlah"`
	Cabe   int    `json:"cabe" form:"cabe"`
}

type UploadImageCloudinaryBase64 struct {
	Image string `json:"image"`
}
