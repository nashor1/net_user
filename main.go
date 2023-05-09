package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const NERR_Success = 0

var (
	netapi32 = syscall.NewLazyDLL("netapi32.dll")

	procNetUserAdd = netapi32.NewProc("NetUserAdd")
)

type USER_INFO_1 struct {
	Name        *uint16
	Password    *uint16
	PasswordAge uint32
	Priv        uint32
	HomeDir     *uint16
	Comment     *uint16
	Flags       uint32
	ScriptPath  *uint16
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: main.exe name=xxx passwd=123456")
		return
	}
	username := args[1][5:]
	password := args[2][7:]
	comment := "testcomment"
	homedir := "C:\\Users\\" + username
	err := CreateUser(username, password, comment, homedir)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User created successfully")
	}
}

func CreateUser(username string, password string, comment string, homedir string) error {
	user := &USER_INFO_1{
		Name:     syscall.StringToUTF16Ptr(username),
		Password: syscall.StringToUTF16Ptr(password),
		Priv:     1,
		HomeDir:  syscall.StringToUTF16Ptr(homedir),
		Comment:  syscall.StringToUTF16Ptr(comment),
		Flags:    0x10000,
	}
	ret, _, err := procNetUserAdd.Call(
		0,
		uintptr(1),
		uintptr(unsafe.Pointer(user)),
		uintptr(0),
	)
	if ret != uintptr(NERR_Success) {
		return fmt.Errorf("Error creating user: %s", err)
	}
	return nil
}
