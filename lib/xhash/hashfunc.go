package xhash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
)

var (
	Md5Inst  hash.Hash
	Sha1Inst hash.Hash
)

func init() {
	Md5Inst = md5.New()
	Sha1Inst = sha1.New()
}

func GenMD5(s interface{}, salt ...interface{}) []byte {
	value := fmt.Sprint(s, salt)
	Md5Inst.Reset()
	Md5Inst.Write([]byte(value))
	return Md5Inst.Sum(nil)
}

func GenSHA1(s interface{}, salt ...interface{}) []byte {
	value := fmt.Sprint(s, salt)
	Sha1Inst.Reset()
	Sha1Inst.Write([]byte(value))
	return Sha1Inst.Sum(nil)
}
