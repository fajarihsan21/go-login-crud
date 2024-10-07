package middleware

import (
	"go-login-crud/internal/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")
		if headerToken == "" {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		token := strings.ReplaceAll(headerToken, "Bearer ", "")
		if token == "" {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		checkToken, err := util.CheckToken(token)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		ctx.Set("userid", *checkToken)
		ctx.Next()
	}
}