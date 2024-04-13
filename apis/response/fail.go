package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type FailResponse struct {
	C *gin.Context
}

func (r *FailResponse) Login(mess string) {
	r.C.Redirect(302, os.Getenv("CLIENT_FALLBACK_URL")+"?error="+mess+"&code=100")
}
func (r *FailResponse) CannotBindBodyReq(mess string) {
	r.C.JSON(http.StatusBadRequest, gin.H{
		"code": 101,
		"mess": mess,
	})
}
func (r *FailResponse) CannotValidateBodyReq(mess string) {
	r.C.JSON(http.StatusBadRequest, gin.H{
		"code": 102,
		"mess": mess,
	})
}
func (r *FailResponse) ErrorPersist(mess string) {
	r.C.JSON(http.StatusInternalServerError, gin.H{
		"code": 103,
		"mess": mess,
	})
}
func (r *FailResponse) FieldGivenEmpty(name ...string) {
	r.C.JSON(http.StatusBadRequest, gin.H{
		"code": 104,
		"mess": strings.Join(name, " "),
	})
}
func (r *FailResponse) TokenIsExpire() {
	r.C.JSON(http.StatusBadRequest, gin.H{
		"code": 105,
	})
}
func (r *FailResponse) CannotParseToken(mess string) {
	r.C.JSON(http.StatusBadRequest, gin.H{
		"code": 106,
		"mess": mess,
	})
}
func (r *FailResponse) InvalidToken() {
	r.C.JSON(http.StatusBadRequest, gin.H{
		"code": 107,
		"mess": "invalid token",
	})
}
