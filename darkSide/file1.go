package darkSide

import (
	"fmt"
	"runtime"
	"time"
)

func DarkSide(){
	//func1()
	//func2()
	func3()
}
func func3(){
	go func() {
		for i := 0 ; i <= 255; i++{

		}
	}()
	fmt.Println("Dropping mic")
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done!")
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

func func1() {
	// this prints incorrect as variable i is shared across and other go routines may not run in that time for loop is running.
	for i := 1; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 5)
}