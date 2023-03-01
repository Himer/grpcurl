package main

import "fmt"

var chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func encode(num int64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

var (
	red   = string([]byte{27, 91, 51, 49, 109})
	reset = string([]byte{27, 91, 48, 109})
)

func printRed(str string) string {
	if str == "" {
		str = "nil"
	}
	return fmt.Sprintf("%s%s%s", red, str, reset)
}
