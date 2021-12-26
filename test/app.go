package main
import("fmt"; "os"; "github.com/rodolfoap/golib";)

func main() {
	// Join(values... string)
	fmt.Println(golib.Join("Hello", ",", " ", "World", "!"))

	// Printhr()
	golib.Printhr()

	// Debug
	golib.Trace()
	golib.Tracef("%v: %d", "Some message", 10)

	//_, e:=os.Open("file.not.found")
	_, e:=os.Open("go.mod")
	golib.Fatal(e)

	_=golib.RunCommand("ls", "-l")
	_=golib.FileToString("go.mod")
}
