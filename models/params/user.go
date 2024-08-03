package params

type CreateUserParams struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
