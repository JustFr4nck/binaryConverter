package main

import(

	"fmt"
	//"unicode"
) 

func main() {

	var val int = 0;

	for {

		fmt.Print("\nInserisci valore da convertire: ")
		fmt.Scanln(&val)

		/*if unicode.IsDigit(val){

			fmt.Print("ERRORE: INSERIRE UN VALORE NUMERICO ")
		} else{
			break;
		}
		*/

		break;
	}

	fmt.Printf("%b", val)

}
