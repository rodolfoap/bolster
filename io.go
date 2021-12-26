package golib
import("io/ioutil")

func FileToString(filename string)string{
	bytes, err := ioutil.ReadFile(filename)
	Fatal(err)
	return string(bytes)
}
