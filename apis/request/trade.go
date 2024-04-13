package request

type ReqBoothUpdateDescription struct {
	AccessToken string `json:"access_token"`
	Description string `json:"description"`
}
type ReqOrderAdd struct {
	ProductCategory2 string  `json:"product_category_2,omitempty"`
	AccessToken      string  `json:"access_token,omitempty"`
	ProductId        string  `json:"product_id,omitempty"`
	TotalAmount      float64 `json:"total_amount,omitempty"`
}
type ReqBoothUpdateBrand struct {
	AccessToken string `json:"access_token"`
	Brand       string `json:"brand"`
}
