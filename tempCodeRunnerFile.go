package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify()  {
	fmt.Println(&u.name)
	fmt.Printf("sending email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeMail(email string)  {
	fmt.Println(&u.email)
	u.email = email
}

func main()  {
	bill := user{"Bill", "bill@mail.com"}
	// fmt.Println(&bill.name)
	// bill.notify()
	lisa := &user{"Lisa", "lisa@mail.com"}
	// fmt.Println(&lisa.name)
	// lisa.notify()
	fmt.Println(&bill.email)
	bill.changeMail("bill@163.com")
	bill.notify()

	lisa.changeMail("lisa@163.com")
	lisa.notify()
}