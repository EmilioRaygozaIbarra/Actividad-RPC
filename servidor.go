package main

import (
	"container/list"
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

//var alumno = make(map[string]int)
var listaAlumno = list.New()

//Alumno Comnetario
type Alumno struct {
	ANombre       string
	AMateria      string
	ACalificacion float64
}

type Server struct{}

//AgregarCalificacion Comentario
func (this *Server) AgregarCalificacion(rec Alumno, reply *float64) error {
	if listaAlumno.Len() > 0 {
		if rec.ANombre == "" || rec.AMateria == "" || rec.ACalificacion == 0 {
			return errors.New("Error: Favor De Ingresar Todos Los Datos Pedidos")
		}
		for e := listaAlumno.Front(); e != nil; e = e.Next() {
			al := e.Value.(*Alumno)
			if rec.ANombre == al.ANombre {
				if rec.AMateria == al.AMateria {
					return errors.New("Error: Ya Existe Una Calificacion De Esta Materia Para El Alumno Ingresado")
				}
			}
		}
		al := &Alumno{ANombre: rec.ANombre, AMateria: rec.AMateria, ACalificacion: rec.ACalificacion}
		listaAlumno.PushBack(al)
		return nil
	} else {
		if rec.ANombre == "" || rec.AMateria == "" || rec.ACalificacion == 0 {
			return errors.New("Error: Favor De Ingresar Todos Los Datos Pedidos")
		}
		al := &Alumno{ANombre: rec.ANombre, AMateria: rec.AMateria, ACalificacion: rec.ACalificacion}
		listaAlumno.PushBack(al)
		return nil
	}

}

func (this *Server) MostrarPromedioAlumno(nombre string, reply *float64) error {
	if nombre == "" {
		return errors.New("Error: Favor De Ingresar Todos Los Datos Pedidos")
	}
	if listaAlumno.Len() > 0 {
		var aux float64
		var contador float64
		contador = 0
		aux = 0
		for e := listaAlumno.Front(); e != nil; e = e.Next() {
			al := e.Value.(*Alumno)
			if nombre == al.ANombre {
				aux += al.ACalificacion
				contador++
			}
		}
		if contador == 0 {
			return errors.New("Error: No Se Encontro Al Alumno")
		} else {
			*reply = aux / contador
			return nil
		}
	} else {
		return errors.New("Error: No Se Han Ingresado Alumnos")
	}
}

func (this *Server) MostrarPromedioGeneral(s string, reply *float64) error {
	if listaAlumno.Len() > 0 {
		var aux float64
		var contador float64
		contador = 0
		aux = 0
		for e := listaAlumno.Front(); e != nil; e = e.Next() {
			al := e.Value.(*Alumno)
			aux += al.ACalificacion
			contador++
		}
		*reply = aux / contador
		return nil
	} else {
		return errors.New("Error: No Se Han Ingresado Alumnos")
	}
}

func (this *Server) MostrarPromedioMateria(materia string, reply *float64) error {
	if materia == "" {
		return errors.New("Error: Favor De Ingresar Todos Los Datos Pedidos")
	}
	if listaAlumno.Len() > 0 {
		var aux float64
		var contador float64
		contador = 0
		aux = 0
		for e := listaAlumno.Front(); e != nil; e = e.Next() {
			al := e.Value.(*Alumno)
			if materia == al.AMateria {
				aux += al.ACalificacion
				contador++
			}
		}
		if contador == 0 {
			return errors.New("Error: No Se Encontro La Materia")
		} else {
			*reply = aux / contador
			return nil
		}
	} else {
		return errors.New("Error: No Se Han Ingresado Alumnos")
	}
}

func server() {
	rpc.Register(new(Server))
	ln, error := net.Listen("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
	}
	for {
		c, error := ln.Accept()
		if error != nil {
			fmt.Println(error)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var pausa string
	fmt.Scanln(&pausa)
}
