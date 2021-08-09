package defs

//reqeusts
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Data model
type User struct {
	Id        int
	LoginName string
	Pwd       string
}
