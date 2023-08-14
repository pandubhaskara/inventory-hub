package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"invhub/database"
	"invhub/helper"

	"gorm.io/gorm"
)

type UniqueField struct {
	ReqField string
	DbField  string
	Model    interface{}
	Struct   interface{}
}

type UniqueFieldExcept struct {
	ReqField      string
	DbField       string
	Model         interface{}
	Struct        interface{}
	ExceptDbField string
	ExceptValue   interface{}
}

func Unique(requests []UniqueField) (map[string]interface{}, bool) {
	errs := make(map[string]string)

	for _, element := range requests {
		if field, ok := reflect.TypeOf(element.Struct).Elem().FieldByName(element.ReqField); ok {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name != "-" {
				val := fmt.Sprintf("%s", reflect.ValueOf(element.Struct).Elem().FieldByName(element.ReqField))

				result := map[string]interface{}{}
				if err := database.Db.Where(fmt.Sprintf("%s = ?", element.DbField), val).Model(element.Model).First(&result).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						helper.Logger.Error(err)
					}
				} else {
					//handle record found
					errs[name] = "duplicate"
				}
			}
		} else {
			helper.Logger.Warnf("Field %s not found in struct %s", element.ReqField, reflect.TypeOf(element.Struct).Elem())
		}
	}

	if len(errs) > 0 {
		return map[string]interface{}{"errors": errs, "message": validationErrorMessage}, true
	} else {
		return nil, false
	}

}

func UniqueExcept(requests []UniqueFieldExcept) (map[string]interface{}, bool) {
	errs := make(map[string]string)

	for _, element := range requests {
		if field, ok := reflect.TypeOf(element.Struct).Elem().FieldByName(element.ReqField); ok {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name != "-" {
				val := fmt.Sprintf("%s", reflect.ValueOf(element.Struct).Elem().FieldByName(element.ReqField))

				result := map[string]interface{}{}
				if err := database.Db.Where(fmt.Sprintf("%s = ? AND %s <> ?", element.DbField, element.ExceptDbField), val, element.ExceptValue).Model(element.Model).First(&result).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						helper.Logger.Error(err)
					}
				} else {
					//handle record found
					errs[name] = "duplicate"
				}
			}
		} else {
			helper.Logger.Warnf("Field %s not found in struct %s", element.ReqField, reflect.TypeOf(element.Struct).Elem())
		}
	}

	if len(errs) > 0 {
		return map[string]interface{}{"errors": errs, "message": validationErrorMessage}, true
	} else {
		return nil, false
	}

}
