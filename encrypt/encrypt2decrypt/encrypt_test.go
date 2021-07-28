package encrypt2decrypt

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEncryptedID(t *testing.T) {
	st := EncryptedID(1234567890, "")
	log.Printf("id= %s", st)
	//assert.NotNil(t, err)
	assert.Equal(t, st, "45bce62b71d42784ec39f07d13b8fe35")
}

func TestDecryptedID(t *testing.T) {
	st := DecryptedID("45bce62b71d42784ec39f07d13b8fe35", "")
	log.Printf("id= %s", st)
	//assert.NotNil(t, err)
	assert.Equal(t, st, "1234567890")
}

//func TestName(t *testing.T) {
//	st := EncryptedID(123,env.ConstEncryptKey)
//	log.Printf("id= %s", st) // 09a064b3c415ac95
//}
