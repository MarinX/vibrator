// vibrator
// port of C file from android hardware vibrator/vibrator.c
package vibrator

import (
	"fmt"
	"os"
)

const (
	THE_DEVICE = "/sys/class/timed_output/vibrator/enable"
)

func VibratorON(timeout_ms int) (int, error) {
	return sendIt(timeout_ms)
}

func VibratorOFF() (int, error) {
	return sendIt(0)
}

func VibratorExists() bool {
	fd, err := os.OpenFile(THE_DEVICE, os.O_RDWR, 0666)
	if err != nil {
		return false
	}

	fd.Close()

	return true
}

func sendIt(timeout_ms int) (int, error) {
	fd, err := os.OpenFile(THE_DEVICE, os.O_RDWR, 0666)
	if err != nil {
		return -1, err
	}

	defer fd.Close()

	return fd.Write([]byte(fmt.Sprintf("%d", timeout_ms)))
}

