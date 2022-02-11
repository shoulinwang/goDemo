package main

import (
	"fmt"
	"goDemo/interface/mock"
	"goDemo/interface/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

/**
duck typing
鸭子类型
*/
func main() {
	var r Retriever
	r = mock.Retriever{Contents: "------------------实现接口-------------------"}
	//fmt.Println(download(r))
	//fmt.Printf("%T %v\n",r,r)
	inspect(r)

	r = &real.Retriever{
		Timeout: time.Duration(time.March),
	}
	//fmt.Printf("%T %v",r,r)
	//fmt.Println(download(r))
	inspect(r)

	//Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	//Type switch
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("Contents: ", v.Contents)

	}
}
