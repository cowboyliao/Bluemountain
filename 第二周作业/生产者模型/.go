package main

import "fmt"

func P(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产数据:", data)
		out <- data // 写入数据
	}
	close(out) //写完关闭管道
}

func C(in <-chan int) {

	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}

}

func main() {
	ch := make(chan int, 5) // 长度为5的管管

	go P(ch) // 生产者
	C(ch)    // 消费者
}
