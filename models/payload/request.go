package payload

type LoginAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type CreateMenuRequest struct {
	Nama       string `json:"nama" form:"nama" validate:"required"`
	Harga      int    `json:"harga" form:"harga" validate:"required"`
	ImageUrl   string `json:"image_url" form:"image_url"`
	CategoryID int    `json:"category_id" form:"category_id" validate:"required"`
}
type UpdateMenuRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Harga    int    `json:"harga" form:"harga"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

type UploadImageCloudinaryBase64 struct {
	Image string `json:"image"`
}
