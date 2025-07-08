package validator

import (
	"univ-academic/model/web"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IEnrollmentValidator interface {
	ValidateEnrollToCourse(req web.EnrollToCourseRequest) error
}

type enrollemntValidator struct{}

func NewEnrollmentValidator() IEnrollmentValidator {
	return &enrollemntValidator{}
}

func (sv *enrollemntValidator) ValidateEnrollToCourse(req web.EnrollToCourseRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.CourseID,
			validation.Required.Error("course_id is required"),
			is.UUID.Error("course_id must be a valid UUID"),
		),
	)
}
