package main

import (
	"fmt"
	"time"
)

func main() {

	nowTime := time.Now().Add(time.Duration(20*time.Second - 8*60*60*time.Second))

	fmt.Println(nowTime.Format("Mon, 02 Jan 2006 15:04:05 GMT"))

}
