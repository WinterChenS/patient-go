package helpers

import (
	"strings"

	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// handle validator error
func HandleValidatorError(c *gin.Context, err error) {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		responses.Error(c, 500, 500, "params is validate", err.Error())
	}
	msg := removeTopStruct(errs.Translate(global.Trans))
	responses.Error(c, 400, 400, "params is validate", msg)
	return
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
