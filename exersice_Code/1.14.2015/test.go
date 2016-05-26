package main

type User struct {
	Id   int
	name string
}

func main() {
	u := User{100, "lily" }
	println(u.Id,u.name)

	var u2*User=new(User)
	*u2=u
	println(u2.name,u2.Id)

}

