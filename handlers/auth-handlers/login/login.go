package loginHandler

import (
	loginAuth "GOFILEGO/controllers/auth-controllers/login"
	"GOFILEGO/models"
	"GOFILEGO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service   loginAuth.Service
	validator *validator.Validate
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input loginAuth.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(ctx, "Invalid input", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, validationError := range validationErrors {
			switch validationError.Tag() {
			case "required":
				errorMessages[validationError.Field()] = validationError.Field() + " is required"
			case "email":
				errorMessages[validationError.Field()] = "Invalid email format"
			default:
				errorMessages[validationError.Field()] = validationError.Error()
			}
		}
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errorMessages)
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)
	h.handleLoginResponse(ctx, resultLogin, errLogin)
}

func (h *handler) handleLoginResponse(ctx *gin.Context, resultLogin *models.UserEntity, errLogin int) {
	switch errLogin {
	case http.StatusNotFound:
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
	case http.StatusUnauthorized:
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
	case http.StatusAccepted:
		h.handleSuccessfulLogin(ctx, resultLogin)
	default:
		utils.APIResponse(ctx, "Unknown error occurred", http.StatusInternalServerError, http.MethodPost, nil)
	}
}

func (h *handler) handleSuccessfulLogin(ctx *gin.Context, resultLogin *models.UserEntity) {
	accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
	accessToken, errToken := utils.Sign(accessTokenData, utils.GodotEnv("JWT_SECRET"), 24*60)

	if errToken != nil {
		logrus.Error(errToken.Error())
		utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
}
