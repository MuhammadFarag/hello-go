package main

func main() {
	var s Speaker
	s = Person{"Muhammad"}
	s.speak()
}

type Speaker interface {
	speak()
}

var _ Speaker = Person{}

type Person struct {
	name string
}

func (p Person) speak() {
	println("My name is", p.name)
}

type SoccerPlayer interface {
	Play()
}

type SoccerSpeaker interface {
	Speaker
	SoccerPlayer
}
