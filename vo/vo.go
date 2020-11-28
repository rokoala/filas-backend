package vo

type CreateRequest struct {
	Name    string `json:name`
	URLName string `json:urlName`
}

type AddConsumerRequest struct {
	StoreID string `json:storeId`
	Name    string `json:name`
	Phone   string `json:phone`
}
