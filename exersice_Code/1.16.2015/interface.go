package main

type Reader interface {
	Read()
}
type Writer interface {
	Write()
}
type ReadWrite interface {
	Read()
	Write()
}
type ReadWriteTest struct {

}

func (*ReadWriteTest) Write(){
	println("Read")
}
func (*ReadWriteTest) Read(){
	println("Write")
}
func main(){
	t:=ReadWriteTest{}
	var rw ReadWrite=&t
	rw.Read()
	rw.Write()
}
