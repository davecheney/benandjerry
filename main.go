package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	name string
}

func (b *Ben) Hello() {
	name := b.name
	if name == "Ben" {
		return
	}
	fmt.Printf("%T says, \"Hello my name is %s\"\n", b, name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	name := j.name
	if name == "Jerry" {
		return
	}
	fmt.Printf("%T says, \"Hello my name is %s\"\n", j, name)
}

type Container struct {
	IceCreamMaker
	_ uintptr
}

func main() {
	var ben = &Ben{"Ben"}
	var jerry = &Jerry{"Jerry"}
	var makers = make([]Container, 18)

	allbens := func() {
		for i := range makers {
			makers[i].IceCreamMaker = ben
		}
	}

	alljerrys := func() {
		for i := range makers {
			makers[i].IceCreamMaker = jerry
		}
	}

	// set all the makers to ben
	allbens()

	var loop0, loop1, hello func()
	loop0 = func() {
		allbens()
		go loop1()
	}

	loop1 = func() {
		alljerrys()
		go loop0()
	}

	hello = func() {
		for i := range makers {
			makers[i].Hello()
		}
		go hello()
	}

	go loop0()
	go hello()

	select {}
}
