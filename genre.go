package online_lib_api

type Genre struct {
	Id int `json:"id" db:"id"`
	Genre string `json:"genre" db:"genre"`
}