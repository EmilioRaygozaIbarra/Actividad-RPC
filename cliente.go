package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
	"time"
)

type AuxAlumno struct {
	ANombre       string
	AMateria      string
	ACalificacion float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	var nombre, materia string
	var calificacion, reply float64

	scanner := bufio.NewScanner(os.Stdin)
	for {
		//Limpiar consola
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("1.- Agregar Calificacion De Materia")
		fmt.Println("2.- Mostrar Promedio De Alumno")
		fmt.Println("3.- Mostrar Promedio General")
		fmt.Println("4.- Mostrar Promedio De Materia")
		fmt.Println("0.- Salir")
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Println("Nombre Del Alumno: ")
			scanner.Scan()
			nombre = scanner.Text()
			fmt.Println("Nombre De La Materia: ")
			scanner.Scan()
			materia = scanner.Text()
			fmt.Println("Calificacion: ")
			fmt.Scanln(&calificacion)
			al := &AuxAlumno{ANombre: nombre, AMateria: materia, ACalificacion: calificacion}
			error := c.Call("Server.AgregarCalificacion", al, &reply)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("Alumno Agregado De Manera Satisfactoria")
			}
			time.Sleep(time.Millisecond * 1500)
		case 2:
			fmt.Println("Nombre Del Alumno: ")
			scanner.Scan()
			nombre = scanner.Text()
			error := c.Call("Server.MostrarPromedioAlumno", nombre, &reply)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("El Promedio De ", nombre, " Es: ", reply)
			}
			time.Sleep(time.Millisecond * 1500)
		case 3:
			error := c.Call("Server.MostrarPromedioGeneral", nombre, &reply)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("El promedio General Es: ", reply)
			}
			time.Sleep(time.Millisecond * 1500)
		case 4:
			fmt.Println("Nombre De La Materia: ")
			scanner.Scan()
			materia = scanner.Text()
			error := c.Call("Server.MostrarPromedioMateria", materia, &reply)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("El Promedio De La Materia Es: ", reply)
			}
			time.Sleep(time.Millisecond * 1500)
		case 0:
			return
		}
	}
}

func main() {
	client()

}
