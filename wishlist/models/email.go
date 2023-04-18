package models

// type emailForm struct {
// 	Email string `form:"email" json:"email" binding:"required"`
// 	Name  string `form:"name" json:"name" binding:"required"`
// }

type Email struct {
	// From    emailForm `form:"from" json:"from" binding:"required"`
	// Subject string    `form:"subject" json:"subject" binding:"required"`
	Body string `form:"body" json:"body" binding:"required"`
}

type EmailService interface {
	Send(email *Email) error
}
