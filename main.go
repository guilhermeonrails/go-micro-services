package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Aluno struct {
	ID     int    `json:"id,string"`
	Codigo string `json:"-"`
	Nome   string `json:"nome,omitempty"`
}

func main() {
	port := 8080

	http.HandleFunc("/alunos", AlunosHandle)

	log.Printf("Iniciando sevidor na porta %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func AlunosHandle(w http.ResponseWriter, r *http.Request) {
	listaDeAlunos := []Aluno{
		{1, "001", "Guilherme"},
		{2, "002", "Ana"},
		{3, "003", ""},
		{14, "014", "Bono"},
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&listaDeAlunos)
}
