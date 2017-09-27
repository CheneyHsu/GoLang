package main

import "net/http"

type SingleHost struct {

	handler    http.Handler
	allowHost string
}

func(this *SingleHost)ServeHTTP(w http.ResponseWriter,r *http.Request ){
	println(r.Host)
	if r.Host==this.allowHost {
		this.handler.ServeHTTP(w, r)
	}else{
		w.WriteHeader(403)
	}

}

func myHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello world!"))
}


func main(){
	single:=&SingleHost{
		handler: http.HandlerFunc(myHandler),
		allowHost: "localhost:8080",
	}
	http.ListenAndServe(":8080",single)
}