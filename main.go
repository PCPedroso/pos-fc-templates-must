package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Seja DevGolang", 40}
	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
