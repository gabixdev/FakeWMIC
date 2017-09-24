package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func main() {
	//wmicConfig := Dir() + string(os.PathSeparator) + "fakewmic.ini"
	logFilename := Dir() + string(os.PathSeparator) + "aegis-wmic.log"

	if _, err := os.Stat(logFilename); os.IsNotExist(err) {
		_, _ = os.Create(logFilename);
	}

	f, err := os.OpenFile(logFilename, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("FakeWMIC: Can't create log file")
		return
	}

	defer f.Close()

	text := "PID " + strconv.Itoa(os.Getppid()) + " executed: " + strings.Join(os.Args, " ") + "\n"

	if _, err = f.WriteString(text); err != nil {
		fmt.Println("FakeWMIC: Can't write file")
	}

	if len(os.Args) == 4 {
		if strings.ToLower(os.Args[1]) == "diskdrive" {
			if strings.ToLower(os.Args[2]) == "get" {
				if strings.ToLower(os.Args[3]) == "serialnumber" {
					fmt.Println("SerialNumber")
					fmt.Println(RandString(16))
				}
			}
		}
	}
}

func Dir() string {
	usr, err := user.Current()
	var homeDir string
	if err == nil {
		homeDir = usr.HomeDir
	} else {
		homeDir = os.Getenv("HOME")
	}
	return homeDir
}

const letterBytes = "12345678901234567890ABCDEFABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
