package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func Validate(c *gin.Context, v interface{}) (errs ValidErrors) {
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, ok := v.(ut.Translator)
		if !ok {
			return
		}
		validationErrors, ok := err.(val.ValidationErrors)
		if !ok {
			return
		}

		for key, value := range validationErrors.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return
	}

	return
}
