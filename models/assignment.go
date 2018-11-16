package models

// Assignment .
type Assignment struct {
	ProjectID UUID `json:"project_id"`
	MemberID  UUID `json:"member_id"`
}
