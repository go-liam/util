package encrypt

import (
	"crypto/sha1"
	"fmt"
)

const Size = 20

// Sha1 :
func Sha1(str string) string {
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(str))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	return fmt.Sprintf("%x", bs)
}

func Sum(data []byte) [Size]byte {
	return sha1.Sum(data)
}
