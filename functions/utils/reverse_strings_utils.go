package utils

func Reverse(s string) string {

	var newStr string

	for i := len(s) - 1; i >= 0; i-- {

		newStr += string(s[i])

	}

	return newStr
}
