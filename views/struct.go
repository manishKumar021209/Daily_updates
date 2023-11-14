package views

type Students struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	RollNumber int    `json:"roll_ number"`
	ParentID   string `json:"student_id"`
}

type Parents struct {
	ParentID        string `json:"student_id"`
	ParentFirstName string `json:"parent_first_name"`
	ParentLastName  string `json:"parent_last_name"`
}

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
