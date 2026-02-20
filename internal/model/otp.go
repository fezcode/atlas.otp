package model

type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Issuer   string `json:"issuer"`
	Secret   string `json:"secret"`
	IsLocked bool   `json:"is_locked"`
}
