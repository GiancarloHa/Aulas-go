package main

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz!"))
}


func main() {
	// HTTP é um protocolo de comunicação - Base da comunicação WEB
	// Cliente (Faz a requisição) - Servidor (Processa requisição e envia resposta)
	// Request - Response
	// Rotas 
		// URI - Indentificador do Recurso
		// Método - GET, POST, PUT, DELETE
	http.HandleFunc("/home", home)
	
	http.HandleFunc("/usuarios", usuarios)
	http.HandleFunc("/", raiz)  
	log.Fatal(http.ListenAndServer(":5000", nil))

}