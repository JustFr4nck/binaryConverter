package main

import (
	"fmt"
	//"unicode"
)

func main() {

	var (
		val int = 0
		scr string = "";
		n int = 0
	)

	fmt.Print("\n___MENU___")
	fmt.Print("\nInserire i valori richiesti per eseguire le seguenti operazioni")
	fmt.Print("\n1) Per trasformare una parola in codice binario")
	fmt.Print("\n2) Per trasformare un valore decimale in binario")
	fmt.Print("\n3) Per trasformare un codice binario in un valore numerico")
	fmt.Print("\n4) Per trasformare un codice binario in una parola")
	fmt.Print("\nEXIT) Inserire qualsiasi altro valore per uscire dal programma\n")

	fmt.Scanln(&n)

	for {

		switch n {

		case 1:
			fmt.Print("\nInserisci parola da convertire: ")
			fmt.Scanln(&scr)

			for _, char := range scr{

				charInt := int(char);

				bin := fmt.Sprintf("%08b", charInt)

				fmt.Print(bin, " ")
			}
			

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
