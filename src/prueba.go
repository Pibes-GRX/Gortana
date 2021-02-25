package main

import "log"
import "os"
import "fmt"

func main() {

	file, err := os.OpenFile("logPibes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    	if err != nil {
		log.Fatal(err)
    	}

        fmt.Println("Hola mundo")
	log.Println("Registrada acción en el log")

	log.SetOutput(file)
	log.Println("Se ha realizado una ejecución satisfatoria")
}


