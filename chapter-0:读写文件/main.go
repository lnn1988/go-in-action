package main

import (
	_ "io/ioutil"
	"fmt"
	"io/ioutil"
	"os"
)

//全量读取
func Read_1()  {
	dat, err := ioutil.ReadFile("./main.go")
	fmt.Println(err)
	fmt.Println(string(dat))
}

//带缓冲区读取
func Read_2()  {
	f, _ := os.Open("./main.go")
	defer f.Close()
	//只读取前 16 字节
	buf := make([]byte, 16)
	f.Read(buf)

	fmt.Println(string(buf))
}

//任意位置读取 f.Seek + f.Read
func Read_3()  {
	f, _ := os.Open("/main.go")
	defer f.Close()

	b1 := make([]byte, 2)
	f.Seek(5, 0)
	f.Read(b1)
	fmt.Println(string(b1))
}

//任意位置读 f.Readat
func Read_4()  {
	f, _ := os.Open("./main.go")
	defer f.Close()

	b1 := make([]byte, 2)
	f.ReadAt(b1, 5)
	fmt.Println(string(b1))
}



func main()  {
	Read_5()
}
