package main

import (
	"fmt"
)

func main() {
	fmt.Printf("--------------- VARIABLES MODIFIED BY VALUE ---------------\n")
	//Defines the variable by value, not by pointer
	var name string = "Walter"
	AddSurname(name, "Tobias")
	//It does not return the surname because is not passed by reference.
	//So it makes a copy of the variable and gives it to the method
	fmt.Println("MAIN -- 'name' variable value", name)
	//'&' returns the memory direction
	fmt.Println("MAIN -- 'name' variable memory direction", &name)

	fmt.Printf("\n\n--------------- VARIABLES MODIFIED BY REFERENCE ---------------\n")
	AddSurnameWithPointer(&name, "Tobias")
	//It returns the value, because of the pointer
	fmt.Println("MAIN WITH POINTER -- 'name' variable value", name)
	fmt.Println("MAIN WITH POINTER -- 'name' variable memory direction", &name)

	fmt.Printf("\n\n--------------- HOW TO INITIALIZE POINTERS ---------------\n")
	HowToInitializePointers()

	fmt.Printf("\n\n--------------- POINTERS AND STRUCTURES ---------------\n")
	PointersAndStructures()
}

func AddSurname(name string, surname string) {
	name = name + " " + surname
	//Returns surname too, is a method local variable because of a copy
	fmt.Println("ADD SURNAME -- 'name' variable value", name)
	fmt.Println("ADD SURNAME -- 'name' variable memory direction", &name)
}

//The method receives a pointer, you shoud get the reference of the variable with *
func AddSurnameWithPointer(name *string, surname string) {
	//To concat the pointer with value referenced variables such as " " and surname, you have to
	//de-reference it, with same operator '*' as used to define a pointer
	*name = *name + " " + surname
	//Returns surname too, is a method local variable because of a copy
	fmt.Println("ADD SURNAME WITH POINTER -- 'name' variable value", *name)
	fmt.Println("ADD SURNAME WITH POINTER -- 'name' variable memory direction", name)
}

func HowToInitializePointers() {
	//declares the variable by value, and assigns this value to a pointer by reference using '&'
	name := "Meison"
	fmt.Printf("name variable by value pointer address: [%v]\n", &name)
	pointer_one := &name
	fmt.Printf("pointer_one address: [%v]\n", pointer_one)
	//declares the pointer directly
	var pointer_two *string
	fmt.Printf("pointer_two address: [%v]\n", pointer_two)
	pointer_two = &name
	//declares a variable and instantiates the pointer with the method new(<pointer_type>)
	var pointer_three = new(string)
	fmt.Printf("pointer_two address: [%v]\n", pointer_three)
	pointer_three = &name
	//Same address to all pointers, because of the assignation to the variable
	fmt.Printf("pointer_one is type: [%T] with value [%v] and the pointer is in address [%v]\n", pointer_one, *pointer_one, pointer_one)
	fmt.Printf("pointer_two is type: [%T] with value [%v] and the pointer is in address [%v]\n", pointer_two, *pointer_two, pointer_two)
	fmt.Printf("pointer_three is type: [%T] with value [%v] and the pointer is in address [%v]\n", pointer_three, *pointer_three, pointer_three)
}

type Cat struct {
	Name         *string
	Nickname     string
	IsLongHaired bool
}

//Function to update nick by value
func UpdateNicknameByValue(cat Cat) {
	cat.Nickname = "Lobi"
	fmt.Println("UPDATE NICKNAME BY VALUE Cat changed nickname in function by value: ", cat)
}

//Method to update nick by reference
func (cat *Cat) UpdateNicknameByReference() {
	(*cat).Nickname = "Leisi" //same as cat.Nickname
	fmt.Println("UPDATE NICKNAME BY REFERENCE Cat changed nickname in method by reference: ", cat)
}

//Method to update nick by reference
func FunctionUpdateNameByReference(cat *Cat) {
	name := "Mincli"
	//Changes the pointer to point to the new string referencing with &
	cat.Name = &name
	//Gets the value of the pointer with *
	fmt.Println("UPDATE NAME BY REFERENCE Cat changed name in function by reference: ", *cat.Name)
}

func PointersAndStructures() {
	name := "Rayo"
	cat := Cat{
		Name:         &name,
		Nickname:     "Miti",
		IsLongHaired: true,
	}

	fmt.Printf("POINTERS AND STRUCTURES Cat base: %v and name %v\n", cat, *cat.Name)
	UpdateNicknameByValue(cat)
	fmt.Printf("POINTERS AND STRUCTURES Cat before nickname changed by value: %v and name %v\n", cat, *cat.Name)
	cat.UpdateNicknameByReference()
	fmt.Printf("POINTERS AND STRUCTURES Cat before nickname changed by reference: %v and name %v\n", cat, *cat.Name)
	FunctionUpdateNameByReference(&cat)
	fmt.Printf("POINTERS AND STRUCTURES Cat before name changed by reference: %v and name %v\n", cat, *cat.Name)
}
