package gx
import("io/ioutil";)

// Load file to string
func ReadFile(filename string) string {
	bytes, err:=ioutil.ReadFile(filename)
	Fatal(err)
	return string(bytes)
}
