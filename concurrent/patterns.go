package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func pattern1(){
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


func TestGoConcurrencyPatterns(){
	pattern1()
}