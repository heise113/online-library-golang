package online_lib_api

type Profile struct {
	Name string `json:"name" db:"name"`
	Username string `json:"username"`
	Mybooks []Book `json:"my_books"`
}