package main

import "fmt"
import "time"

func main() {
	i := 2

	fmt.Print("Write", i, "as ")
	switch i {
	case 1 :
		fmt.Println("one")
	case 2 :
		fmt.Println("two")
	case 3 :
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default :
		fmt.Println("It's a weekday")
	}

	whatAmi := func (i interface{}) {
		switch t := i.(type){
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmi(true)
	whatAmi(1)
	whatAmi("hey")
}