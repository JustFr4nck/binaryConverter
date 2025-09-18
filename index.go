package main

import (
	"fmt"
	//"unicode"
)

func main() {

	var (
		val int = 0
		n   int = 0
	)

	fmt.Print("\n___MENU___")
	fmt.Print("Inserire i valori richiesti per eseguire le seguenti operazioni")
	fmt.Print("1) Per trasformare un testo in un codice binario")
	fmt.Print("2) Per trasformare un valore decimale in binario")
	fmt.Print("3) Per trasformare un codice binario in un valore numerico o in un testo")
	fmt.Print("EXIT) Inserire qualsiasi altro valore per uscire dal programma")

	fmt.Scanln(&n)

	for {

		switch n {

		case 1:
			

		case 2:

			fmt.Print("\nInserisci valore da convertire: ")
			fmt.Scanln(&val)

			/*if unicode.IsDigit(val){

				fmt.Print("ERRORE: INSERIRE UN VALORE NUMERICO ")
			} else{
				break;
			}
			*/

			fmt.Printf("%b", val)

			break

		case 3:
			break
		
		case 4:
			break

		default:
			break

		}


		break

	}

}
