package request

type CreateCheckInRecordRequest struct {
	StudentName   string   `json:"studentName"`
	CourseName    string   `json:"courseName"`
}