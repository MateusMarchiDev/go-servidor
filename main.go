package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil) // O parâmetro nil indica que o servidor deve usar o gerenciador de certificados padrão do sistema.
}
