package controllers

import "net/http"

// CriarUsuarios cria um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando Usuario!"))
}

// BuscarUsuarios busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuario!"))
}

// BuscarUsuario busca um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um  Usuario!"))
}

// AtualizarUsuario atualiza um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuarios"))
}

//DeletarUsuario deleta um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar um usuarios"))
}
