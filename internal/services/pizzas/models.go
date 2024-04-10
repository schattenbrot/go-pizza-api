package pizzas

// Pizza represents a pizza object.
// swagger:model
type Pizza struct {
	// ID represents the pizza ID as ObjectID.
	// Example: ESESRIRIESAEREEEEEiseseseaiN
	ID string `json:"id"`

	// Name represents the pizza name.
	// Example: Pizza Salami
	Name string `json:"name"`

	// Description represents a short description like an ingredients-list.
	// Example: Salami, Mushrooms, Cheddar
	Description string `json:"description"`

	// ImageURL represents the pizza's image url.
	// Example: https://files.schattenbrot.com/api/v1/files/file-0000000000.jpg
	ImageURL string `json:"imageURL"`

	// Price represents the pizza's price in euro.
	// Example: 5.99
	Price float32 `json:"price"`
}
