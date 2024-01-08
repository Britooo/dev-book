package controller

import (
	"net/http"
)

// Create register a user in API
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuários"))
}

// List retrives all users in API
func List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários"))
}

// ById serch for a user by id
func ById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário por ID"))
}

// Update update's a user in API by ID
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// Delete a user in API by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removendo usuário por id"))
}
