package main

type User struct {
	Id int
	name string
}
type Manager struct {
	User
	group string
}

func main(){
	m :=Manager{User{1,"Jack"},"IT"}
    println(m.Id,m.name,m.group)

	m.name="rose"
	println(m.name,m.Id,m.group)
}
