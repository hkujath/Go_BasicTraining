package advanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hkujath/Go_BasicTraining/forms"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func LearnInterface() {
	myUser := user{name: "Holger"}
	output(myUser)

	var s fmt.Stringer
	s = user{name: "Hermann"}
	upperOutput(s)

	myCircle := forms.Circle{Radius: 5, XPos: 1, YPos: 2}
	myRect := forms.Rectangle{Length: 2, Width: 4, XPos: 4, YPos: 5}

	fmt.Println("Forms have same area?", forms.SameArea(myCircle, myRect))

	// Important to set a concrete type to an interface to use it
	var w io.Writer
	b := &bytes.Buffer{}
	w = b
	w.Write([]byte("moin"))
	fmt.Println(w)
	fmt.Println(b)

	// From interface to type
	var r io.Reader
	checkInterface(r)
	r = bytes.NewBufferString("Hello interface")
	//r = strings.NewReader("Hello newreader")
	checkInterface(r)

	// Using type assertion
	var r2 io.Reader
	r2 = bytes.NewBufferString("Hello interface")
	b2 := r2.(*bytes.Buffer)
	fmt.Println("Data with type assertion", b2.Bytes())

	var r3 io.Reader
	r3 = bytes.NewBufferString("Hello interface")
	f, ok := r3.(*os.File)
	if !ok {
		fmt.Printf("*os.File expected. Got %T\n", r3)
	}
	defer f.Close()

	// Struct with some interfaces
	test := myType{
		Reader: strings.NewReader("mytype"),
		Writer: &bytes.Buffer{},
	}
	fmt.Printf("%v\n", test)
	fmt.Printf("%+v\n", test)
	fmt.Printf("%#v\n", test)
}

// Function to test interface as function parameter
func sumFile(inFileName, outFileName string) error {
	in, err := os.Open(inFileName)
	defer in.Close()
	if err != nil {
		return fmt.Errorf("sumFile: open inFile: %w", err)
	}

	out, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer out.Close()
	if err != nil {
		return fmt.Errorf("sumFile: creating outfile: %w", err)
	}
	return sum(in, out)
}

func sum(r io.Reader, w io.Writer) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("sum(): reading data: %w", err)
	}

	var numbers []int
	err = json.Unmarshal(b, &numbers)
	if err != nil {
		return fmt.Errorf("sum(): unmarshal numbers: %w", err)
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	_, err = w.Write([]byte(fmt.Sprintf("%d", sum)))
	return err
}

// Interfaces in structs

type myType struct {
	io.Reader
	io.Writer
}

// Interface of interfaces

type ReadWriter interface {
	io.Reader
	io.Writer
}

type ReadWriteCloser interface {
	ReadWriter
	Close()
}

// Type check of an interface
func checkInterface(r io.Reader) {
	switch v := r.(type) {
	case nil:
		fmt.Println("nil pointer")
	case *bytes.Buffer:
		fmt.Println("bytes.Buffer", v.Bytes())
	case *strings.Reader:
		fmt.Println("string.reader", v.Len())
	default:
		fmt.Println("Type unknown")

	}
}

// Empty interface which an accept anything
func PrintLn(a ...interface{}) (n int, err error) {

	return 0, nil
}

// User interface

type user struct {
	name string
}

func (u user) String() string {
	return "Name: " + u.name
}

// Interface for functions

type MyInterfacer interface {
	String() string
}

func output(s MyInterfacer) {
	fmt.Println(s.String())
}

func upperOutput(s MyInterfacer) {
	upper := strings.ToUpper(s.String())
	fmt.Println(upper)
}
