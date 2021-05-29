package darkSide

import (
	"fmt"
	"time"
)

func Func1() {
	for i := 1; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 5);
}
