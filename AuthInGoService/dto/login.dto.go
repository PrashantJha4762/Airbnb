package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"` //this json thing is called as json marshalling. basically it is a process of converting a go struct into json format. and the validate thing is called as validation tag. it is used to validate the input data. for example, in this case, we are validating that the email field is required and it should be a valid email address.
	Password string `json:"password" validate:"required"`
}