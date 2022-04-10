package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	path := "/home/work/memfd_write_test.txt"
	cnt := 0
	for {
		err := ioutil.WriteFile(path, []byte("test--"+strconv.Itoa(cnt)+"\n"), 0666) //写入文件(字节数组)
		if err != nil {
			fmt.Printf("write err is %v", err.Error())
		}

		time.Sleep(time.Second * 5)
		cnt++
		if cnt > 100 {
			break
		}
	}
	return
}
