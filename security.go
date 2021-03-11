package gg_common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/a316523235/gg-common/conf"
)

func Md5(str string) string {
	if str == "" {
		return ""
	}
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodePwd(pwd string) string {
	return Md5(conf.PasswordHalt + pwd)
}

func EncodeToken(username, encodePwd string) string {
	str := fmt.Sprintf("%s.%s.%s", conf.TokenHalt, username, encodePwd)
	return Md5(str)
}

