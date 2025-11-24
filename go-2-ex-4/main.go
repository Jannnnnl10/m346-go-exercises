package main

import "fmt"

type Student struct {
 FirstName string
 LastName  string
}

type Class []Student

type Modules map[int][]Class

func main() {
 classA := Class{
  {FirstName: "Cristiano", LastName: "Ronaldo"},
  {FirstName: "Jan", LastName: "Lehmann"},
  {FirstName: "Lionel", LastName: "Messi"},
 }
 classB := Class{
  {FirstName: "Steven", LastName: "Gerrard"},
  {FirstName: "Eden", LastName: "Hazard"},
  {FirstName: "Neymar ", LastName: "Junior"},
 }

 modules := Modules{
  104: {classA, classB},
  117: {classA},
  346: {classB},
 }

 fmt.Println("Class A:", classA)
 fmt.Println("Class B:", classB)
 fmt.Println("Modules:", modules)
}
 
