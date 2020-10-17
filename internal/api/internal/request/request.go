package request

type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Create struct {
	Title string `json:"Title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
