package golib
import("fmt"; "os"; "testing";)

func Test_strings(t *testing.T){
	fmt.Println(Join("Hello", ",", " ", "World", "!"))
	Printhr()
	Trace()
	Tracef("%v: %d", "Some message", 10)
	_, e:=os.Open("go.mod")
	Fatal(e)

	// RunCommand(... string) string
	s:=RunCommand("ls", "-l")
	fmt.Println(s)

	// FileToString(string) string
	s=FileToString("go.mod")
	fmt.Println(s)
}
