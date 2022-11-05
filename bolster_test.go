package bolster
import("fmt"; "os"; "testing";)

func Test_strings(t *testing.T) {
	fmt.Println(Join("Hello", ",", " ", "World", "!"))
	Printhr()
	Trace()
	Trace("Some message", 10)
	_, e:=os.Open("go.mod")
	Error(e)

	// ReadFile(string) string
	s:=ReadFile("go.mod")
	fmt.Println(s)
}
