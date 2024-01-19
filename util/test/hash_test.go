package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hash := md5.New()
	hash.Write([]byte("123456"))
	bytes := hash.Sum(nil)
	encodeToString := hex.EncodeToString(bytes)
	fmt.Println(encodeToString)
	//digest := util.Md5("123456")
	//if digest != "e10adc3949ba59abbe56e057f20f883e" {
	//	fmt.Println(digest)
	//	t.Fail()
	//}
}

// go test -v .\util\test\ -run=^TestHash$ -count=1
