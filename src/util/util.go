package util

type Animaler interface {
	Speak() string
}

type Cat struct {
	name string
}

type Pig struct {
	name string
}

func (cat *Cat) Speak() string {
	return "my name is " + cat.name
}

func (pig *Pig) Speak() string {
	return "my name is " + pig.name
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (pig *Pig) SetName(name string) {
	pig.name = name
}
