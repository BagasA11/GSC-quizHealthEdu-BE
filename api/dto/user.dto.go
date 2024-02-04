package dto

type AdminLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type AdminCreate struct {
	Email      string `json:"email" binding:"required,min=5"`
	Password   string `json:"password" binding:"required,min=8"`
	C_password string `json:"c-password" binding:"required,min=8"`
}

type UserCreate struct {
	Username   string `json:"username" binding:"required,min=8"`
	Email      string `json:"email" binding:"required,min=5"`
	Password   string `json:"password" binding:"required,min=8"`
	C_password string `json:"c-password" binding:"required,min=8"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdatePassword struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	Password    string `json:"password" binding:"required,min=8"`
}

type UpdateUsername struct {
	Username string `json:"username" binding:"required,min=8"`
}

type DeleteUser struct {
	Password string `json:"password" binding:"required,min=8"`
}
