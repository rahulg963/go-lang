package darkSide

import (
	"fmt"
	"time"
)

func DarkSide(){
	//func1()
	func2()
}

func func1() {
	// this prints incorrect as variable i is shared across and other go routines may not run in that time for loop is running.
	for i := 1; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 5)
}

func func2() {
	// this prints incorrect as variable i is shared across and other go routines may not run in that time for loop is running.
	for i := 1; i < 10; i++ {
		i :=  i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 5)
}
