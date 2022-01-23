package helper

import (
	"fmt"
	"os"
	"strconv"
)

var path = "app/healthy.txt"

func ResolvePath(host string, port string) string {
	h := os.Getenv(host)
	p, _ := strconv.Atoi(os.Getenv(port))
	return fmt.Sprintf("%s:%d", h, p)
}

func CreateHealthFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}

func DeleteHealthFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
