package golib
import("fmt"; "testing";)

func Test_strings(t *testing.T){
	fmt.Println(Join("Hello", ",", " ", "World", "!"))
	Printhr()
	Trace()
	Tracef("%v: %d", "Some message", 10)
}
