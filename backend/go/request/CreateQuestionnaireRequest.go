package request

type CreateQuestionnaireRequest struct {
	CourseName   string   `json:"courseName"`
	StudentName  string   `json:"studentName"`
	Satisfaction int    `json:"satisfaction"`
	Comment      string `json:"comment"`
	Suggestion   string `json:"suggestion"`
}