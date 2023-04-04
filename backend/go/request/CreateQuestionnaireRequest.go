package request

type CreateQuestionnaireRequest struct {
	CourseId     uint   `json:"courseId"`
	StudentId    uint   `json:"studentId"`
	Satisfaction int    `json:"satisfaction"`
	Comment      string `json:"comment"`
	Suggestion   string `json:"suggestion"`
}