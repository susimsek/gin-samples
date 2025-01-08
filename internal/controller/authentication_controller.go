package controller

import (
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	"gin-samples/internal/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController interface {
	Login(c *gin.Context)
}

type authenticationControllerImpl struct {
	authService service.AuthenticationService
	validator   *validator.Validate
	trans       ut.Translator
}

// NewAuthenticationController creates a new instance of AuthenticationController
func NewAuthenticationController(authService service.AuthenticationService, validator *validator.Validate, trans ut.Translator) AuthenticationController {
	return &authenticationControllerImpl{
		authService: authService,
		validator:   validator,
		trans:       trans,
	}
}

// Login godoc
// @Summary Authenticate user and generate token
// @Description Validates user credentials and returns a JWT token
// @Tags authentication
// @Accept json
// @Produce json
// @Param input body dto.LoginInput true "Login Input"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} dto.ProblemDetail
// @Failure 401 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/auth/login [post]
func (a *authenticationControllerImpl) Login(c *gin.Context) {
	var input dto.LoginInput

	// Parse and validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(&customError.MessageNotReadableError{
			Detail: err.Error(),
		})
		return
	}

	if err := a.validator.Struct(input); err != nil {
		_ = c.Error(err)
		return
	}

	// Authenticate the user
	tokenResponse, err := a.authService.Authenticate(input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	// Return the token response
	c.JSON(http.StatusOK, tokenResponse)
}
