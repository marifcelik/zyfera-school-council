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
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	StdNumber string `json:"stdNumber"`
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
