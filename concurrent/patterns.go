package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func pattern1() {
	// normal channels (not buffered) will block until there is a msg in channel,
	//producer and consumer statements will block
	c := make(chan string)
	go boring1("Boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring1(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func pattern2() {
	c := boring2("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring2(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func pattern2Dash() {
	// this is channels as a handle on a service
	// both are blocking each other, not a good design if not required explicitly
	joe := boring2("Joe!")
	ann := boring2("Ann!")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("You say: %q\n", <-ann)
	}
	fmt.Println("You're boring; I'm leaving")
}

func pattern3(){
	// none of the channels are blocking each other
	// sync is lost between joe and ann
	c := fanIn(boring2("Joe"), boring2("Ann"))
	for i:= 0; i < 10; i++{
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {for { c <- <-input1 } }()
	go func() {for { c <- <-input2 } }()
	return c
}


type Message struct {
	str string
	wait chan bool
}

func pattern4(){
	// channels within channels.. making sync again
	// sync is restored between joe and ann
	c := fanIn2(boring3("Joe"), boring3("Ann"))
	for i:= 0; i < 5; i++{
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring; I'm leaving")
}

func fanIn2(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {for { c <- <-input1 } }()
	go func() {for { c <- <-input2 } }()
	return c
}

func boring3(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				str: fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			// waiting for pattern4 to add true, while other one is running
			<-waitForIt
		}
	}()
	return c
}

func TestGoConcurrencyPatterns() {
	//pattern1()
	//pattern2()
	//pattern2Dash()
	//pattern3()
	pattern4()
}
