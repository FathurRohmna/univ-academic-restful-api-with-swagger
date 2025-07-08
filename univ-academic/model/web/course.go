package web

type CourseDetails struct {
	CourseID      string `json:"course_id"`
	CourseName    string `json:"course_name"`
	CourseDesc    string `json:"course_description"`
	Credits       int    `json:"credits"`
	Department    string `json:"department"`
	ProfessorName string `json:"professor_name"`
}
