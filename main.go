package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Data struct {
	Games  []Game
	Amount int
}

var slice = make([]Game, 0)

func RenderForms(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		// if len(slice) != 0 {
		// 	mergeSort(slice, false, "name")
		// 	fmt.Println(slice)
		// }
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := Data{
			Games:  slice,
			Amount: 1,
		}
		tmpl.Execute(w, data)
	} else {
		r.ParseForm()

		flag := strings.Join(r.Form["flag"], "")
		reverse, _ := strconv.Atoi(strings.Join(r.Form["reverse"], ""))
		amount, _ := strconv.Atoi(strings.Join(r.Form["amount"], ""))

		flagReverse := false
		if reverse == 1 {
			flagReverse = true
		}

		gamesList := extractData()
		gamesList = mergeSort(gamesList, flagReverse, flag)
		tmpl := template.Must(template.ParseFiles("templates/list.html"))
		data := Data{
			Games:  gamesList[:amount],
			Amount: amount,
		}
		tmpl.Execute(w, data)
	}
}

func RenderList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Amount:", r.Form["amount"])
	amount, _ := strconv.Atoi(strings.Join(r.Form["amount"], ""))
	fmt.Println(amount)
	gamesList := extractData()
	gamesList = mergeSort(gamesList, false, "name")
	tmpl := template.Must(template.ParseFiles("templates/list.html"))
	data := Data{
		Games:  gamesList,
		Amount: amount,
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", RenderForms)
	http.HandleFunc("/result", RenderList)
	http.ListenAndServe(":8080", nil)
}
