package bolster
import("bytes"; "fmt"; "strings"; "golang.org/x/term";)

// Prints an horizontal row the width of the terminal
func Printhr() {
	if !term.IsTerminal(0) {
		fmt.Println(strings.Repeat("=", 80))
		return
	}
	width, _, _:=term.GetSize(0)
	fmt.Println(strings.Repeat("=", width))
}

// Joins multiple string values
func Join(values ...string) string {
	var b bytes.Buffer
	for _, v:=range values {
		b.WriteString(v)
	}
	return b.String()
}
