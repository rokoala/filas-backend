package vo

// CreateRequest struct
type CreateRequest struct {
	Name string `json:"name"`
}

// AddConsumerRequest struct
type AddConsumerRequest struct {
	StoreID string `json:"storeId"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
}
