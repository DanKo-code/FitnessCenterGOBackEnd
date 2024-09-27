package http

import (
	"FitnessCenter_GoBackEnd/auth"
	"FitnessCenter_GoBackEnd/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

type Handler struct {
	useCase  auth.UseCase
	validate *validator.Validate
}

type signInput struct {
	FirstName string `json:"first_name" validate:"validFirstName"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password"`
}

func NewHandler(useCase auth.UseCase) *Handler {
	validate := validator.New()
	err := validate.RegisterValidation("validFirstName", validators.ValidateUsreFisrtName)
	if err != nil {
		return nil
	}

	return &Handler{
		useCase:  useCase,
		validate: validate,
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		log.Println(err)
		return
	}

	err := h.validate.Struct(inp)
	if err != nil {
		log.Println(err)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), inp.FirstName, inp.LastName, inp.Email, inp.Password); err != nil {
		log.Println(err)
	}
}
