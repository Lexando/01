package main

import "fmt"

type dog struct {
	kind     string
	name     string
	weight   float32
	foodCons int
	foodStep int
}

func (d dog) animalKind() string {
	return d.kind
}

func (d dog) animalName() string {
	return d.name
}

func (d dog) animalCons() int {
	return d.foodCons
}

func (d dog) animalWeight() float32 {
	return d.weight
}

func (d dog) animalStep() int {
	return d.foodStep
}

type cat struct {
	kind     string
	name     string
	weight   float32
	foodCons int
	foodStep int
}

func (c cat) animalKind() string {
	return c.kind
}

func (c cat) animalName() string {
	return c.name
}

func (c cat) animalCons() int {
	return c.foodCons
}

func (c cat) animalWeight() float32 {
	return c.weight
}

func (c cat) animalStep() int {
	return c.foodStep
}

type cow struct {
	kind     string
	name     string
	weight   float32
	foodCons int
	foodStep int
}

func (co cow) animalKind() string {
	return co.kind
}

func (co cow) animalName() string {
	return co.name
}

func (co cow) animalCons() int {
	return co.foodCons
}

func (co cow) animalWeight() float32 {
	return co.weight
}

func (co cow) animalStep() int {
	return co.foodStep
}

type animal interface {
	getKind
	getName
	getCons
	getWeight
	getStep
}

type getKind interface {
	animalKind() string
}

type getName interface {
	animalName() string
}

type getCons interface {
	animalCons() int
}

type getWeight interface {
	animalWeight() float32
}

type getStep interface {
	animalStep() int
}

func calcAnimalCons(aCons getCons, aWeight getWeight, aStep getStep) float32 {
	return aWeight.animalWeight() * float32(aCons.animalCons()) / float32(aStep.animalStep())
}

func main() {
	oneDog := dog{kind: "Dog", name: "Toozique", weight: 13.4, foodCons: 10, foodStep: 5}
	oneCat := cat{kind: "Cat", name: "Moorzique", weight: 4.1, foodCons: 7, foodStep: 1}
	oneCow := cow{kind: "Cow", name: "Boorionqa", weight: 372, foodCons: 25, foodStep: 1}
	fmt.Println(oneDog.animalName(), oneCat, oneCow)

	var (
		aKind      getKind
		aName      getName
		aCons      getCons
		aWeight    getWeight
		aStep      getStep
		foodAmount float32
	)
	animals := []animal{
		oneDog,
		oneCat,
		oneCow,
	}

	for _, a := range animals {
		aCons = a
		aWeight = a
		aStep = a
		aName = a
		aKind = a
		foodConsumption := calcAnimalCons(aCons, aWeight, aStep)
		fmt.Printf("\n%s %s consumpts %.2f kg of food a month", aKind.animalKind(), aName.animalName(), foodConsumption)
		foodAmount += foodConsumption
	}
	fmt.Printf("\nAmount of food to be supplied to the farm every month is %.2f kg", foodAmount)
}
