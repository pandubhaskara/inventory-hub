package validation

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HttpRequestValidationErrorJson(err error) (int, map[string]interface{}) {
	var verr validator.ValidationErrors
	errs := make(map[string]string)

	if errors.As(err, &verr) {
		for _, f := range verr {
			e := f.ActualTag()
			if f.Param() != "" {
				e = fmt.Sprintf("%s=%s", e, f.Param())
			}
			errs[f.Field()] = e
		}
		return http.StatusUnprocessableEntity, map[string]interface{}{"errors": errs, "message": validationErrorMessage}
	}
	return http.StatusBadRequest, map[string]interface{}{"message": badRequestMessage}

}
