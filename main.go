package main

import (
  "html/template"
  "net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, nil)
}

func main() {  
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", HomePage)
  
  http.ListenAndServe(":8000", nil)
}