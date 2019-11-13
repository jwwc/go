/*
字符串的底层定义
type StringHeader struct{
	Data uintptr
	Len int
}*/
package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var str = "hello,world"
	p := str2bytes(str)
	fmt.Println(p)
	str2 := bytes2str(p)
	fmt.Println(str2)
	re := str2runes(str)
	fmt.Println(re)
}
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}
func str2runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size:= utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		fmt.Println(s)
		s = s[size:]
		fmt.Println(s)
	}
	return []rune(p)
}
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
