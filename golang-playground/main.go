package main

import (
	"errors"
	"fmt"
)

func main() {
	temp := ""
	fmt.Println([]byte(temp))
	// fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05-0700"))
	//  ISO8601 를 써야한다.
	//fmt.Println(time.Now().UTC())
	//temp := "2022-06-17T06:33:35Z"
	//parse, err := time.Parse("2006-01-02T15:04:05-0700", temp)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(parse)
	// June 17th 2022, 16:06:55.549
}

func myfunc() error {
	return errors.New("test")
}

//func myfunc2() error {
//
//}
