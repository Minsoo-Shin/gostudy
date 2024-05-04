package main

type Vehicle struct {
	speed int
}

func (r *Vehicle) isEmpty() bool {
	return r.speed == 0
}

type Car struct {
	Vehicle
}

func (r *Vehicle) move() {

}

func main() {
	car := Car{}
	println(car.isEmpty())
}
