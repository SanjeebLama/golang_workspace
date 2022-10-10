package main

import "fmt"

type Speak interface {
	Speak() string
}

type Parrot struct {
	parrotSpeak string
}

type Cow struct {
	cowMoo string
}

type Cat struct {
	catMeow string
}

func (p Parrot) Speak() string {
	return p.parrotSpeak
}

func (c Cow) Speak() string {
	return c.cowMoo
}

func (c Cat) Speak() string {
	return c.catMeow
}

func main() {
	parrot := Parrot{"parrotSpeak"}
	cow := Cow{"Cow Moo"}
	cat := Cat{"Cat Meow"}

	generateSpeak(parrot)
	generateSpeak(cow)
	generateSpeak(cat)

}

func generateSpeak(s Speak) {
	fmt.Println(s.Speak())
}
