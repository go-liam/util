package uuid

import (
	"github.com/go-liam/util/uuid/sonyflake"
	"log"
)

// AutoInt64ID :
func AutoInt64ID() int64 {
	v,err :=  sonyflake.GetID()
	if err != nil{
		log.Printf("[ERROR] AutoInt64ID %+v\n",err)
	}
	return int64(v)
}

func AutoUInt64ID() uint64 {
	v,err :=  sonyflake.GetID()
	if err != nil{
		log.Printf("[ERROR] AutoInt64ID %+v\n",err)
	}
	return v
}
