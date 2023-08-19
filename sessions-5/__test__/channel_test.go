package test

import (
	"fmt"
	"testing"
	"time"
)

func introduce(s string, ch chan string)  {
	result := fmt.Sprintf("hallo %s\n", s)
	ch <- result
}

func TestImplementChannel(t *testing.T)  {
	// make a channel
	ch := make(chan string)

	go introduce("momo", ch)
	go introduce("sana", ch)
	go introduce("dahyun", ch)
	go introduce("mina", ch)
	go introduce("nayeon", ch)
	go introduce("chaeyoung", ch)

	// menerima data dari channel
	data1 := <- ch
	fmt.Println("data 1 :", data1)

	data2 := <- ch
	fmt.Println("data 2 :", data2)

	data3 := <- ch
	fmt.Println("data 3 : ", data3)

	data4 := <- ch
	fmt.Println("data 4 : ", data4)

	data5 := <- ch
	fmt.Println("data 5 : ", data5)

	data6 := <- ch
	fmt.Println("data 6 : ", data6)

	close(ch)
}

func print(ch chan string)  {
	fmt.Println(<-ch)
}

func TestAnonyomuse(t *testing.T) {
	ch := make(chan string)

	for _, v := range []string{"momo", "sana", "mina", "chaeyoung"} {
		go func(s string) {
			result := fmt.Sprintf("hallo %s\n", s)
			ch <- result
		}(v)
	}

	for i := 0; i < 4; i++ {
		print(ch)
	}

	close(ch)
}

func send(s string, ch chan <- string)  {
	result := fmt.Sprintf("hallo %s\n", s)
	ch <- result
}

func receive(ch <- chan string)  {
	fmt.Println(<- ch)
}

func TestChannelDirections(t *testing.T) {
	ch := make(chan string)

	for _, v := range []string{"momo", "sana", "mina", "chaeyoung"} {
		go send(v, ch)
	}

	for i := 0; i < 4; i++ {
		receive(ch)
	}

	close(ch)
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 5)

	go func(c chan int) {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d start send data to channel\n", i)
			c <- i
			fmt.Printf("%d after send data to channel\n", i)
		}
		close(c)
	}(ch)

	fmt.Println("sleep 2 second")
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("received value from channel", v)
	}
}

func TestUnbufferedChannel(t *testing.T) {
	ch := make(chan int)

	go func(ch chan int) {
		fmt.Println("before sending data")
		ch <- 10
		fmt.Println("after sending data")
	}(ch)

	fmt.Println("sleeps")
	time.Sleep(2 * time.Second)

	fmt.Println("start receive data")
	data := <- ch
	fmt.Println("receive data", data)

	close(ch)
	time.Sleep(2 * time.Second)
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "momo"
	}()

	go func() {
		ch2 <- "dahyun"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- ch1:
			fmt.Println("message 1 : ", msg1)
		case msg2 := <- ch2:
			fmt.Println("message 2 : ", msg2)
		}
	}

	close(ch1)
	close(ch2)
}