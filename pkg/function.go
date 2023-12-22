package pkg

import (
	"encoding/binary"
	"sync"
	"time"
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

func NewObjSync(instanceId int) *ObjSync {
	obj := ObjSync{
		mu:           sync.Mutex{},
		prvTimestamp: 0,
		instanceId:   instanceId,
	}
	return &obj
}

func (oSync *ObjSync) GenServiceObjID() int64 {
	oSync.mu.Lock()
	defer oSync.mu.Unlock()
	var ret int64 = 0

	binsID := make([]byte, 8)
	baseB := make([]byte, 8)
	instanceB := make([]byte, 4)

	var instanceMod = oSync.instanceId % 256 // max 256 instance

	t := time.Now().UnixMilli()
	if t <= oSync.prvTimestamp {
		ret = oSync.prvTimestamp + 1
	} else {
		ret = t
	}
	oSync.prvTimestamp = ret

	binary.BigEndian.PutUint64(baseB, uint64(ret))
	binary.BigEndian.PutUint32(instanceB, uint32(instanceMod))

	// set first 6byte
	binsID[1] = baseB[2]
	binsID[2] = baseB[3]
	binsID[3] = baseB[4]
	binsID[4] = baseB[5]
	binsID[5] = baseB[6]
	binsID[6] = baseB[7]

	// next 1 byte for instance id
	binsID[7] = instanceB[3]

	ret = int64(binary.BigEndian.Uint64(binsID))

	return ret
}
