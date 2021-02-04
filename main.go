package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Aluno struct {
	ID     int
	Codigo string
	Nome   string
}

func main() {
	port := 8080
	http.HandleFunc("/alunos", AlunosHandler)
	log.Printf("Iniciando Servidor:%v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func AlunosHandler(w http.ResponseWriter, r *http.Request) {
	listaDeAlunos := []Aluno{
		{1, "001", "Guilherme"},
		{2, "002", "Ana"},
		{3, "003", "Maria"},
		{14, "014", "Bono"},
	}
	data, err := json.Marshal(listaDeAlunos)
	if err != nil {
		panic("Ops, algo deu errado")
	}
	fmt.Fprint(w, string(data))
}
