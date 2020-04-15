package unit

import "regexp"

func VerifyPhone(phone string) bool {
	return regexp.MustCompile( `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`).MatchString(phone)
}
