package request

type ReqAddProduct struct {
	Category1   string   `json:"category_1" binding:"required"`
	Category2   string   `json:"category_2" binding:"required"`
	AccessToken string   `json:"access_token" binding:"required"`
	AvatarUrl   string   `json:"avatar_url" binding:"required"`
	PicturesUrl []string `json:"pictures_url" binding:"required"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       float64  `json:"price" binding:"required"`
	Sale        float64  `json:"sale" binding:"required"`
}
type ReqUpdateSaleProduct struct {
	Category2 string  `json:"category_2,omitempty"`
	Id        string  `json:"id,omitempty"`
	Sale      float64 `json:"sale,omitempty"`
}
type ReqAddRatingUser struct {
	Category2 string `json:"category_2,omitempty"`
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FeedBack  string `json:"feed_back,omitempty"`
	Score     int    `json:"score,omitempty"`
}
type ReqAddProductRanking struct {
	Category2 string `json:"category_2,omitempty"`
	Id        string `json:"id,omitempty"`
}
