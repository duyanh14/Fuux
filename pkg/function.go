package pkg

import (
	"sync"
)

func IsMapContainNil(s map[string]interface{}, omitKey []string) bool {
	for k, v := range s {
		if contains(omitKey, k) == false {
			if v == "" {
				return true
			}
		}
	}
	return false
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func AllKeyRequired(s map[string]string, requiredKey []string) bool {
	for _, v := range requiredKey {
		_, ok := s[v]
		// If the key exists
		if ok != true {
			return false
		}
	}
	return true
}

type ObjSync struct {
	mu sync.Mutex

	prvTimestamp int64
	instanceId   int
}
