package payload

type LoginAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type CreateMenuRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Price      int    `json:"price" form:"price" validate:"required"`
	ImageUrl   string `json:"image_url" form:"image_url"`
	CategoryID int    `json:"category_id" form:"category_id" validate:"required"`
}
type UpdateMenuRequest struct {
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

type UploadImageCloudinaryBase64 struct {
	Image string `json:"image"`
}

type AddMenuToCartRequest struct {
	CustomerId int    `json:"customer_id" form:"customer_id"`
	MenuID     int    `json:"menu_id" form:"menu_id" validate:"required"`
	Quantity   int    `json:"quantity" form:"quantity" validate:"required"`
	TakeAway   bool   `json:"take_away" form:"take_away"`
	Note       string `json:"note" form:"note"`
}

type UpdateCartItemRequest struct {
	CustomerId int    `json:"customer_id" form:"customer_id"`
	CartItemID int    `json:"cart_item_id" form:"cart_item_id" validate:"required"`
	Quantity   int    `json:"quantity" form:"quantity" validate:"required"`
	TakeAway   bool   `json:"take_away" form:"take_away"`
	Note       string `json:"note" form:"note"`
}

type DeleteCartItemRequest struct {
	CustomerId int `json:"customer_id" form:"customer_id"`
	CartItemID int `json:"cart_item_id" form:"cart_item_id" validate:"required"`
}

type CreateCustomerRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"max=11,min=10,number"`
}
