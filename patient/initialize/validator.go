package initialize

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"winterchen.com/patient-go/patient/global"

	"reflect"
	"strings"
)

// init translation
func InitTrans(locale string) {
	color.Red("start translation")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //zh translation
		enT := en.New() //en translation

		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			panic(fmt.Errorf("uni.GetTranslator(%s)", locale))
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		}

		return
	}
	return
}
