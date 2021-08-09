package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"github.com/rahulg963/go-lang/darkSide"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rahulg963/go-lang/controllers"
	"github.com/rahulg963/go-lang/middleware"
	"github.com/rahulg963/go-lang/model"
)

// go run github.com/rahulg963/go-lang
// go build github.com/rahulg963/go-lang
func main() {
	//fmt.Println("Hello from a module, Gophers!")
	// concurrent testing
	// concurrent.ConcurrentLearningWithoutGoRoutine()
	// concurrent.ConcurrentLearningWithGoRoutine()

	// channels
	// concurrent.ChannelsDemo()

	// logParser()
	// learningSyntax()

	// startingWebServer()
	//db := connectToDatabase()
	//defer db.Close()
	//result, err := model.Login("test@gmail.com", "123")
	//if err != nil {
	//	fmt.Errorf("Error in retrieve query %v", err)
	//}
	//fmt.Print(result)

	// concurrency patterns
	//concurrent.TestGoConcurrencyPatterns()

	//dark side of go
	darkSide.DarkSide()
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://nnichlle:r8cFon2vgxkZ72Li5xGQJm55T21WwXrS@rajje.db.elephantsql.com:5432/nnichlle")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	err = db.Ping()
	if err != nil {
		fmt.Errorf("Unable to connect to database ping method: %v", err)
	}

	model.SetDatabase(db)
	return db
}

func logParser() {
	// go run . -level INFO
	// go run . -help
	path := flag.String("path", "myapp.log", "Path to the log that should be analysed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}

func startingWebServer() {
	controllers.RegisterControllers()
	fmt.Println("Web server starting")
	// without middleware
	// err := http.ListenAndServe(":3000", nil)

	// with middleware
	// err := http.ListenAndServe(":3000", new(middleware.GzipMiddleware))

	// with timeout and middleware
	err := http.ListenAndServe(":3000", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
	if err != nil {
		log.Fatal(err)
	}
}

func learningSyntax() {
	variables()
	pointers1()
	pointers2()
	constants()
	iotaOrConstantExpression()
	collections()
	structExample()
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
