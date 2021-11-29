package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
//import "C"

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	//return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}

func Lib_add(a, b int) {
	lib := syscall.NewLazyDLL("DllTestDef.dll")
	fmt.Println("dll:", lib.Name)
	add := lib.NewProc("add")
	fmt.Println("+++++++NewProc:", add, "+++++++")

	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil {
		fmt.Println("Lib_add:lib.dll运算结果为:", ret)
	}
}

func DllTestDef_add(a, b int) {
	DllTestDef, _ := syscall.LoadLibrary("DllTestDef.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", DllTestDef, "+++++++")
	defer syscall.FreeLibrary(DllTestDef)
	add, err := syscall.GetProcAddress(DllTestDef, "add")
	fmt.Println("GetProcAddress", add)

	ret, _, err := syscall.Syscall(add,
		2,
		IntPtr(a),
		IntPtr(b),
		0)
	if err != nil {
		fmt.Println("DllTestDef_add:DllTestDef.dll运算结果为:", ret)
	}
}

func DllTestDef_add2(a, b int) {
	DllTestDef := syscall.MustLoadDLL("DllTestDef.dll")
	add := DllTestDef.MustFindProc("add")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil {
		fmt.Println("DllTestDef_add2:DllTestDef的运算结果为:", ret)
	}
}

func Lib_addstr(a, b string) {
	lib := syscall.NewLazyDLL("DllTestDef.dll")
	fmt.Println("dll:", lib.Name)
	addstr2 := lib.NewProc("addstr2")
	fmt.Println("+++++++NewProc:", addstr2, "+++++++")

	ret, _, err := addstr2.Call(StrPtr(a), StrPtr(b))
	if err != nil {
		fmt.Println("Lib_add:lib.dll运算结果为:", ret)
	}
}

func DllTestDef_addstr2(a, b string) {
	DllTestDef := syscall.MustLoadDLL("DllTestDef.dll")
	add := DllTestDef.MustFindProc("addstr")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(StrPtr(a), StrPtr(b))
	if err != nil {
		//defer C.free(unsafe.Pointer(ret))
		p := (*byte)(unsafe.Pointer(ret))
		data := make([]byte, 0)
		for *p != 0 {
			data = append(data, *p)
			ret += unsafe.Sizeof(byte(0))
			p = (*byte)(unsafe.Pointer(ret))
		}
		s := string(data)
		fmt.Println("DllTestDef_add2:DllTestDef的运算结果为:", s)
	}
}

func Lib_strcpy(a, b string) {
	lib := syscall.NewLazyDLL("DllTestDef.dll")
	fmt.Println("dll:", lib.Name)
	addstr2 := lib.NewProc("my_strcpy")
	fmt.Println("+++++++NewProc:", addstr2, "+++++++")
	dest := StrPtr(a)
	ret, _, err := addstr2.Call(dest, StrPtr(b))
	if err != nil {
		fmt.Println("Lib_strcpy:lib.dll运算结果为:", ret)
		fmt.Println("Lib_strcpy:lib.dll运算结果为:", a)
		fmt.Println("Lib_strcpy:lib.dll运算结果为:", b)
		fmt.Println("Lib_strcpy:lib.dll运算结果为:", dest)
		p := (*byte)(unsafe.Pointer(dest))
		data := make([]byte, 0)
		for *p != 0 {
			data = append(data, *p)
			dest += unsafe.Sizeof(byte(0))
			p = (*byte)(unsafe.Pointer(dest))
		}
		s := string(data)
		fmt.Println("Lib_strcpy:DllTestDef的运算结果为:", s)
	}
}

func test() {
	fmt.Println(strings.Fields("hello widuu golang")) //out  [hello widuu golang]
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)

	splitArr := strings.Fields(command)
	strLen := len(splitArr)
	if len(splitArr) > 0 {
		fmt.Println(len(splitArr[strLen-1]))
		//log.Println("command", command)
	} else {
		fmt.Println(0)
	}
}

func test2() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please input string")
	scanner.Scan()
	s1 := strings.ToLower(scanner.Text())
	fmt.Println("please input ch")
	scanner.Scan()
	s2 := strings.ToLower(scanner.Text())

	count := 0
	for _,chr := range s1 {
		if chr == int32(s2[0]) {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	test()
	test2()
	Lib_add(4, 5)
	DllTestDef_add(3, 5)
	DllTestDef_add2(2, 5)
	DllTestDef_addstr2("4567", "123")
	Lib_addstr("4567", "123")
	Lib_strcpy("", "12389545454")
}
