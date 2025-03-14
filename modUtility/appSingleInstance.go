package modUtility

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

func EnsureSingleInstance() bool {
	lockFile := "/tmp/" + C_APPID + G_AppToken + ".pid"
	if _, err := os.Stat(lockFile); err == nil {
		file, err := os.Open(lockFile)
		if err != nil {
			log.Fatalf("Error opening PID file: %v", err)
		}
		defer file.Close()

		var pid int
		_, err = fmt.Fscanf(file, "%d", &pid)
		if err != nil {
			log.Fatalf("Error reading PID from file: %v", err)
		}

		if isProcessRunning(pid) {
			fmt.Println("Another instance of the program is already running.")
			return false
		}
	}

	file, err := os.Create(lockFile)
	if err != nil {
		log.Fatalf("Error creating PID file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		log.Fatalf("Error writing PID to file: %v", err)
	}

	return true
}

func isProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}
