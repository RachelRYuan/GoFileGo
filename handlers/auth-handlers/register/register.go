package registerHandler

import (
	registerAuth "GOFILEGO/controllers/auth-controllers/register"
	"GOFILEGO/models"
	"GOFILEGO/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service   registerAuth.Service
	validator *validator.Validate
}

func NewHandlerRegister(service registerAuth.Service) *handler {
	return &handler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input registerAuth.RegisterInput
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
			case "gte":
				errorMessages[validationError.Field()] = validationError.Field() + " must be at least " + validationError.Param() + " characters"
			case "lowercase":
				errorMessages[validationError.Field()] = validationError.Field() + " must be in lowercase"
			default:
				errorMessages[validationError.Field()] = validationError.Error()
			}
		}
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errorMessages)
		return
	}

	registerResult, errorCode := h.service.RegisterService(&input)
	h.handleRegisterResponse(ctx, registerResult, errorCode)
}

func (h *handler) handleRegisterResponse(ctx *gin.Context, registerResult *models.UserEntity, errorCode int) {
	switch errorCode {
	case http.StatusCreated:
		h.handleSuccessfulRegistration(ctx, registerResult)
	case http.StatusConflict:
		utils.APIResponse(ctx, "Email already taken", http.StatusConflict, http.MethodPost, nil)
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to create an account", http.StatusExpectationFailed, http.MethodPost, nil)
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, http.MethodPost, nil)
	}
}

func (h *handler) handleSuccessfulRegistration(ctx *gin.Context, registerResult *models.UserEntity) {
	accessTokenData := map[string]interface{}{"id": registerResult.ID, "email": registerResult.Email}
	token, errToken := utils.Sign(accessTokenData, utils.GodotEnv("JWT_SECRET"), 60)
	if errToken != nil {
		logrus.Error(errToken.Error())
		utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	var data registerAuth.RegisterResponse
	jsonData, _ := json.Marshal(registerResult)
	if err := json.Unmarshal(jsonData, &data); err != nil {
		logrus.Error("Failed to parse registration result: ", err)
		utils.APIResponse(ctx, "Failed to process registration result", http.StatusInternalServerError, http.MethodPost, nil)
		return
	}

	data.Token = token
	utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, data)
}
