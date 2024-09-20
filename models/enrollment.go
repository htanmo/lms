package models

type Enrollment struct {
	ID        int `json:"id"`
	CourseID  int `json:"course_id"`
	StudentID int `json:"student_id"`
	Progress  int `json:"progress"`
}
