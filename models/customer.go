package models

type Customer struct {
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Company       string `json:"company,omitempty"`
	PostCode      string `json:"post_code,omitempty"`
	TermsAccepted *bool  `json:"terms_accepted" validate:"required"`
	Date          string `json:"date"`
}
