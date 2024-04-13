package response

type RespBooth struct {
	Email         string
	Brand         string
	IsAuthentic   bool
	Description   string
	MainCategory  string
	ProductId     []map[string]string
	ProductSoldId []map[string]string
}
