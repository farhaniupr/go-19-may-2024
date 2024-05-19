package model

type User struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama_lengkap"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
