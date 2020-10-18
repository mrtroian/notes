package request

type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewRegister() *Register {
	return new(Register)
}

// Stub for objpool
func (*Register) Reset() {}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLogin() *Login {
	return new(Login)
}

func (*Login) Reset() {}

type CreateNote struct {
	Title string `json:"Title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

func NewCreateNote() *CreateNote {
	return new(CreateNote)
}

func (*CreateNote) Reset() {}
