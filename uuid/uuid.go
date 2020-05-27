package uuid

import (
	"fmt"
	"log"

	"github.com/go-liam/util/uuid/sonyflake"
)

// AutoInt64ID :
func AutoInt64ID() int64 {
	v, err := sonyflake.GetID()
	if err != nil {
		log.Printf("[ERROR] AutoInt64ID %+v\n", err)
	}
	return int64(v)
}

func AutoUInt64ID() uint64 {
	v, err := sonyflake.GetID()
	if err != nil {
		log.Printf("[ERROR] AutoInt64ID %+v\n", err)
	}
	return v
}

// UUID :
func UUID() string {
	v, err := sonyflake.GetID()
	if err != nil {
		log.Printf("[ERROR] AutoInt64ID %+v\n", err)
	}
	return fmt.Sprintf("%d", v)
}
