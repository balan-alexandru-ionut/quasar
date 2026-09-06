package user

type CreateUserRequest struct {
	FirstName string `json:"name" validate:"required, min=1,max=128"`
	LastName  string `json:"last_name" validate:"required, min=1,max=128"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=128,password_strength"`
}
