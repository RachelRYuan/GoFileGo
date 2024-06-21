package middlewares

import (
	"net/http"

	"GOFILEGO/utils"

	"github.com/gin-gonic/gin"
)

type UnauthorizedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

// Auth middleware to check for authorization header and validate JWT token
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errorResponse UnauthorizedError
		errorResponse.Method = ctx.Request.Method

		authorizationHeader := ctx.GetHeader("Authorization")
		if authorizationHeader == "" {
			errorResponse.Status = "Forbidden"
			errorResponse.Code = http.StatusForbidden
			errorResponse.Message = "Authorization is required for this endpoint"
			ctx.JSON(http.StatusForbidden, errorResponse)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		token, err := utils.VerifyTokenHeader(ctx, "JWT_SECRET")
		if err != nil {
			errorResponse.Status = "Unauthorized"
			errorResponse.Code = http.StatusUnauthorized
			errorResponse.Message = "Invalid Access token"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set user claims in the context
		ctx.Set("user", token.Claims)

		// Proceed to the next handler
		ctx.Next()
	}
}
