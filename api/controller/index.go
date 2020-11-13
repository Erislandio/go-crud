package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	model "github.com/alura-curso-crud-api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := model.List()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na convensão do preço")
		}

		quantidadeConvertida, _ := strconv.Atoi(quantidade)

		model.AdicionarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	model.RemoverProduto(id)
	http.Redirect(w, r, "/", 301)
}
