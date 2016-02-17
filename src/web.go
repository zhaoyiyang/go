package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	w.Write([]byte("go web! master"))
=======
	w.Write([]byte("go web! dev!"))
>>>>>>> dev
}
