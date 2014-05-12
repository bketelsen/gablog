package eject

import (
	"errors"
	"syscall"
	"unsafe"
)

var winmm = syscall.MustLoadDLL("winmm.dll")
var mciSendStringProc = winmm.MustFindProc("mciSendStringW")
var mciGetErrorStringProc = winmm.MustFindProc("mciGetErrorStringW")

func mciGetErrorString(mcierr int) string {
	var b [256]uint16
	mciGetErrorStringProc.Call(uintptr(mcierr), uintptr(unsafe.Pointer(&b[0])), uintptr(256))
	return syscall.UTF16ToString(b[:])
}

func mciSendString(cmd string) int {
	r1, _, _ := mciSendStringProc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(cmd))), 0, 0, 0)
	return int(r1)
}

func Eject() error {
	r := mciSendString("capability cdaudio can eject")
	if r == 0 {
		r = mciSendString("set cdaudio door open")
	} else {
		r = mciSendString("set cdaudio door close")
	}
	if r != 0 {
		return errors.New(mciGetErrorString(r))
	}
	return nil
}
