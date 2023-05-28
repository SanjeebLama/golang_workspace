package main

func main() {
	t := triangle{
		base:   10,
		height: 20,
	}

	printArea(t)

	s := square{
		sideLength: 10,
	}

	printArea(s)

}
