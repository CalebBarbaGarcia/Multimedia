package multimedia

import (
	"fmt"
)

type ContenidoWeb struct {
	Contenido []Multimedia
}

func (cw *ContenidoWeb) Agregar(m Multimedia){
	cw.Contenido = append(cw.Contenido,m)
}

func (cw *ContenidoWeb) Mostrar(){
	for _, f := range cw.Contenido {
		f.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo string
	Formato string
	Canales int64
}

func (i *Imagen) Mostrar() {
	fmt.Println("Titulo: ", i.Titulo," Formato: ", i.Formato, " Canales: ", i.Canales)
}

type Audio struct {
	Titulo string
	Formato string
	Duracion float64
}

func (a *Audio) Mostrar() {
	fmt.Println("Titulo: ", a.Titulo," Formato: ", a.Formato, " Duracion: ", a.Duracion)
}

type Video struct {
	Titulo string
	Formato string
	Frames int64
}

func (v *Video) Mostrar() {
	fmt.Println("Titulo: ", v.Titulo," Formato: ", v.Formato, " Frames: ", v.Frames)
}