package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/is0metry/listman/list"
)

var temp = template.Must(template.ParseFiles("list.html"))
var fl = list.FileList{Root: "root"}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	if _, err := os.Stat(path + ".txt"); os.IsNotExist(err) {
		file, err := os.Create(path + ".txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		file.Close()
	}
	lst, err := fl.GetList(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	temp.Execute(w, lst)
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/add/"):]
	item := r.FormValue("newItem")
	err := fl.AddItem(name, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+name, http.StatusFound)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/root", http.StatusFound)
}
func removeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/remove/"):]
	names := strings.Split(path, "/")
	item, _ := strconv.Atoi(names[1])
	if err := fl.RemoveItem(names[0], item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/remove/", removeHandler)
	http.ListenAndServe(":8080", nil)
}
