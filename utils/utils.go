package utils

import (
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var numRand *rand.Rand = nil

func getNumRand() *rand.Rand {
	if numRand == nil {
		numRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return numRand
}

func GetRandFloat(len int) string {
	r := getNumRand()
	tmp := "0."
	for i := 0; i < len; i++ {
		tmp += strconv.Itoa(r.Int() % 10)
	}
	return tmp
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}

	return exist
}

func WriteFile(filename string, reader io.ReadCloser) error {
	var f *os.File
	var err error
	if CheckFileIsExist(filename) { //如果文件存在
		//f, err = os.OpenFile(filename, os.O_APPEND, 0777) //打开文件
		errDel := os.Remove(filename)
		if errDel != nil {
			return errDel
		}
	}
	f, err = os.Create(filename) //创建文件

	if err != nil {
		return err
	}
	_, err = io.Copy(f, reader)
	return err
}
