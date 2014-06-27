package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	name string
}

func (b *Ben) Hello() {
	if b.name == "Ben" {
		return
	}
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	if j.name == "Jerry" {
		return
	}
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
}

func main() {
	var ben = &Ben{"Ben"}
	var jerry = &Jerry{"Jerry"}
	var makers [18]IceCreamMaker

	allbens := func() {
		for i := range makers {
			makers[i] = ben
		}
	}

	alljerrys := func() {
		for i := range makers {
			makers[i] = jerry
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
