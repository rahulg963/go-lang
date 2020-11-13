package main

import (
	"fmt"

	"github.com/rahulg963/go-lang/models"
)

func main() {
	fmt.Println("Hello from a module, Gophers!")
	// variables()
	// pointers1()
	// pointers2()
	// constants()
	// iotaOrConstantExpression()
	// collections()
	// structExample()

	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}

func variables() {
	// variable declarations
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}

func pointers1() {
	var firstName *string = new(string)
	*firstName = "Arthur"
	fmt.Println(firstName)
	fmt.Println(*firstName)
}

func pointers2() {
	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}

func constants() {
	// value should be available at compile time
	const pi = 3.1415
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)
	fmt.Println(c + 1.2)
}

func iotaOrConstantExpression() {
	// iota is initialized with 0 and then increamented by 1
	const (
		first = iota + 6
		// this expression will take default above constant expression, here iota will be 1, so value will be 1 + 6 = 7
		second
		third
	)

	const (
		// iota will reinitialize in new constant block, so it will be 0
		fourth = iota
		fifth
	)
	// 6 7 8 0 1
	fmt.Println(first, second, third, fourth, fifth)
}
func collections() {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr1[:]
	fmt.Println(arr1, slice)

	// see changes in slice changes value of array and vice versa
	arr1[1] = 42
	slice[2] = 27
	fmt.Println(arr1, slice)

	// see size is not mentioned
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 27)
	fmt.Println(slice2)

	s2 := slice2[1:4]
	fmt.Println(s2)

	//maps
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}

func structExample() {
	type user struct {
		// default will 0
		ID int
		// default will be ""
		FirstName  string
		SecondName string
	}

	var u1 user
	u1.ID = 1
	u1.FirstName = "Arthur"
	u1.SecondName = "Dent"
	fmt.Println(u1)

	u2 := user{
		ID:         1,
		FirstName:  "Arthur",
		SecondName: "Dent",
	}
	fmt.Println(u2)

	u3 := user{
		ID:         1,
		SecondName: "Dent",
	}
	fmt.Println(u3)
}
