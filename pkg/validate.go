package pkg

import (
	"net/mail"
	"reflect"
	"regexp"
)

func IsEmailValid(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

func IsPhoneNumberValid(e string) bool {
	emailRegex := regexp.MustCompile(`(84|0[3|5|7|8|9])+([0-9]{8})\b`)
	return emailRegex.MatchString(e)
}

func IsStructContainNil(s interface{}) bool {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() != reflect.String && val.Field(i).Kind() != reflect.Struct {
			if val.Field(i).IsNil() {
				return true
			}
		}
	}
	return false
}
