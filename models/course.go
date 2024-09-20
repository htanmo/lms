package models

type Course struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	InstructorID int    `json:"instructor_id"`
}
