package sonyflake

// 分布式ID
import (
	"encoding/binary"
	"fmt"
	sony "github.com/sony/sonyflake"
	"math/rand"
	"os"
	"time"
)

var sf *sony.Sonyflake

func init() {
	t, _ := time.Parse("2006-01-02", "2019-01-01")
	settings := sony.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf = sony.NewSonyflake(settings)
	if sf == nil {
		panic("New Sonyflake fail!")
	}
}

func checkMachineID(machineID uint16) bool {
	fmt.Println("This Machine Hostname Code(uint16): ", machineID)
	return true
}

func getMachineID4() (uint16, error) {
	return 1, nil
}

func getMachineID() (uint16, error) {
	var machineID uint16
	ht, err := os.Hostname()
	rand.Seed(time.Now().UnixNano())
	if err != nil {
		fmt.Println("This Machine Get Hostname Fail: ", err)
		machineID = uint16(rand.Intn(1000000))
	} else {
		rd := rand.Intn(1000)
		machineID = uint16(int(binary.BigEndian.Uint16([]byte(ht))) + rd)
		fmt.Println("This Machine Hostname Code(string): ", ht, rd)
	}
	return machineID, nil
}

// GetID 获取分布式ID
func GetID() (uint64, error) {
	id, err := sf.NextID()
	return id, err
}
