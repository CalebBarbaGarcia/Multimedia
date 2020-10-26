package main

import (
	"fmt"
	"./multimedia"
	"bufio"
	"os"
)



func main()  {
	cw01 := multimedia.ContenidoWeb{}
	var opcion int64
	var titulo string
	var formato string
	var canales int64
	var frames int64
	var duracion float64

	hacerMenu := true

	for (hacerMenu){
		fmt.Println("Menu: ")
		fmt.Println("1) Agregar imagen")
		fmt.Println("2) Agregar Audio")
		fmt.Println("3) Agregar video")
		fmt.Println("4) Mostrar")
		fmt.Println("5) Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&opcion)
		fmt.Println("")

		if (opcion == 1 || opcion == 2 || opcion == 3){
			fmt.Print("Escriba el titulo: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			titulo = scanner.Text()
			fmt.Print("Escriba el formato: ")
			fmt.Scanln(&formato)

			if (opcion == 1){
				fmt.Print("Escriba el numero de los canales: ")
				fmt.Scanln(&canales)

				cw01.Agregar(&multimedia.Imagen{Titulo:titulo,Formato:formato,Canales:canales})

			}else if (opcion == 2){
				fmt.Print("Escriba el numero de la duracion: ")
				fmt.Scanln(&duracion)

				cw01.Agregar(&multimedia.Audio{Titulo:titulo,Formato:formato,Duracion:duracion})
			} else{
				fmt.Print("Escriba el numero de frames: ")
				fmt.Scanln(&frames)

				cw01.Agregar(&multimedia.Video{Titulo:titulo,Formato:formato,Frames:frames})
			}
		} else if (opcion == 4){
			cw01.Mostrar()
		} else if (opcion == 5){
			hacerMenu = false
		} else{
			fmt.Println("Opcion incorrecta")
		}
		fmt.Println("")
	}
} 