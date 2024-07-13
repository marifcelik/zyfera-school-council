package dto

type GradeRequest struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type GradeResponse struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type Student struct {
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	StdNumber string `json:"stdNumber,omitempty"`
}

type StudentUpdate struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}

type CreateRequest struct {
	Student
	Grades []GradeRequest `json:"grades"`
}

type CreateResponse struct {
	Student
	ID     uint            `json:"id"`
	Grades []GradeResponse `json:"grades"`
}

type UpdateRequest struct {
	StudentUpdate
	Grades []GradeRequest `json:"grades"`
}

type UpdateResponse struct {
	Student
	ID        uint            `json:"id"`
	Grades    []GradeResponse `json:"grades"`
	UpdatedAt string          `json:"updated_at"`
}
