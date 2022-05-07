package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	mfdCloexec  = 0x0001
	memfdCreate = 319
)

func main() {
	testFilename := "./memfd_write_test"
	data, err := ioutil.ReadFile(testFilename)
	if err != nil {
		fmt.Printf("read file(%v) err is %v", testFilename, err)
	}

	fdName := "memfd_write_test"
	fdm, r2, errNo := unix.Syscall(memfdCreate, uintptr(unsafe.Pointer(&fdName)), uintptr(mfdCloexec), 0)
	if errNo != 0 {
		fmt.Printf("Syscall memfdCreate err: %v, r2: %v\n", errNo, r2)
		return
	}

	n, err := syscall.Write(int(fdm), data)
	if err != nil {
		fmt.Printf("syscall Write err: %v, len of writing byte\n", err, n)
		return
	}

	cmdline := fmt.Sprintf("/proc/self/fd/%d", fdm)
	cmd := exec.Command(cmdline)
	if err := cmd.Start(); err != nil {
		fmt.Printf("start %v err is %v", cmdline, err)
		return
	}
	cmd.Process.Release()

	fmt.Printf("cmd is %+v", cmd)
	time.Sleep(time.Second * 60)
}
