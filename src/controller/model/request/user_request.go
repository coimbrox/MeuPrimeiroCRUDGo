package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20,containsany=!@#$%^&*()"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Name     string `json:"name" binding:"required,min=3,max=80"`
	Age      int    `json:"age" binding:"required,min=1,max=140"`
}