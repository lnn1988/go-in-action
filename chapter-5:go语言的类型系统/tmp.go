package main

import "fmt"

type T struct {
	value int
}

func (t T) StayTheSame ()  {
	t.value = 3
}

func (t *T) Update () {
	t.value = 3
}

func main()  {
	m := T{0}
	fmt.Println(m)

	m.StayTheSame()
	fmt.Println(m)

	m.Update()
	fmt.Println(m)
}