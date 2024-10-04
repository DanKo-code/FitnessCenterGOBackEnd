package dtos

type SignUpDTO struct {
	FirstName   string `json:"first_name" validate:"validFirstName"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" binding:"email"`
	Password    string `json:"password"`
	FingerPrint string `json:"fingerPrint"`
}

type SignInDTO struct {
	Email       string `json:"email" binding:"email"`
	Password    string `json:"password"`
	FingerPrint string `json:"fingerPrint"`
}
