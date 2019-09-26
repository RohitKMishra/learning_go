package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []person

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {

	p1 := person{"James", 32}
	p2 := person{"MoneyPenny", 26}
	p3 := person{"Q", 64}
	p4 := person{"M", 34}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
}
