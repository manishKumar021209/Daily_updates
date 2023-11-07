package views

type Students struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	RollNumber int    `json:"roll_ number"`
}

type Parents struct {
	StudentID       string `json:"student_id"`
	ParentFirstName string `json:"parent_first_name"`
	ParentLastName  string `json:"parent_last_name"`
}
