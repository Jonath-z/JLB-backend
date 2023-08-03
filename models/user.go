package models

type User struct {
	UserId           string `json:"userId,omitempty"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	ProfileUrl       string `json:"profileUrl,omitempty"`
	ProfileThumbnail string `json:"profileThumbnail,omitempty"`
	IsVerified       bool   `json:"isVerified"`
	Password         string `json:"password"`
}

type Product struct {
	ProductId string  `json:"productId,omitempty"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Image     string  `json:"image"`
}

type LoginData struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
