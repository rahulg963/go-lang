package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func pattern1(){
	// normal channels (not buffered) will block until there is a msg in channel,
	//producer and consumer statements will block
	c := make(chan string)
	go boring1("Boring!", c)
	for i := 0; i < 5; i++{
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring1(msg string, c chan string) {
	for i := 0; ; i++{
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func pattern2(){
	c := boring2("boring!")
	for i :=0; i < 5; i++{
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring2(msg string) <-chan string{
	c := make(chan string)
	go func() {
		for i := 0; ; i++{
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}


func pattern2Dash(){
	// this is channels as a handle on a service
	// both are blocking each other, not a good design if not required explicitly
	joe := boring2("Joe!")
	ann := boring2("Ann!")

	for i :=0; i < 5; i++{
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("You say: %q\n", <-ann)
	}
	fmt.Println("You're boring; I'm leaving")
}

func TestGoConcurrencyPatterns(){
	//pattern1()
	//pattern2()
	pattern2Dash()
}