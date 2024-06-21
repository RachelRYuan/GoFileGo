package loginHandler

import (
	loginAuth "GOFILEGO/controllers/auth-controllers/login"
	"GOFILEGO/controllers/auth-controllers/register"
	"GOFILEGO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service loginAuth.Service
}

// NewHandlerLogin initializes a new login handler with the given service.
func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

// LoginHandler handles user login requests.
func (h *handler) LoginHandler(ctx *gin.Context) {
	var input loginAuth.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(ctx, "Invalid JSON format", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Email",
				Message: "Email is required in the body",
			},
			{
				Tag:     "email",
				Field:   "Email",
				Message: "Email format is not valid",
			},
			{
				Tag:     "required",
				Field:   "Password",
				Message: "Password is required in the body",
			},
		},
	}
	errResponse, errCount := utils.GoValidator(&input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {
	case http.StatusNotFound:
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case http.StatusUnauthorized:
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return

	case http.StatusAccepted:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := utils.Sign(accessTokenData, "JWT_SECRET", 24*60*60) // 24 hours in seconds

		if errToken != nil {
			logrus.Error(errToken.Error())
			utils.APIResponse(ctx, "Failed to generate access token", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		var data register.RegisterResponse
		utils.ObjectToJson(resultLogin, &data)
		data.Token = accessToken

		utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, data)
		return

	default:
		utils.APIResponse(ctx, "Unknown error occurred", http.StatusInternalServerError, http.MethodPost, nil)
	}
}
