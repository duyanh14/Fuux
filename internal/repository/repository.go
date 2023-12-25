package repository

import (
	"encoding/json"
	"reflect"
	"strings"
)

func UpdateField(save any) map[string]interface{} {
	updateField := map[string]interface{}{}

	val := reflect.ValueOf(save).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		value := field.Interface()
		tag := val.Type().Field(i).Tag.Get("gorm")

		if field.Kind() != reflect.Ptr {
			continue
		}

		if field.IsNil() {
			continue
		}

		var cl string
		vl := value

		z := strings.Split(tag, ";")
		if len(z) > 0 {
			for _, v := range z {
				zz := strings.Split(v, ":")
				if len(zz) == 0 {
					continue
				}

				switch zz[0] {
				case "column":
					cl = zz[1]
				case "serializer":
					if zz[1] == "json" {
						vlz, err := json.Marshal(value)
						if err == nil {
							vl = string(vlz)
						}
					}
				}
			}
		} else {
			cl = strings.Split(tag, ":")[1]
		}

		updateField[cl] = vl
	}

	return updateField
}
