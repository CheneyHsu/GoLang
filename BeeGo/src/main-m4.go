package main

import (
	"net/http"
	"net/http/httptest"
)

type ModifieMiddleware struct {
	handler http.Handler
}

func (this *ModifieMiddleware) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	rec := httptest.NewRecorder()
	this.handler.ServeHTTP(rec,r)

	for k,v:=range rec.Header(){
		w.Header()[k]=v
	}
	w.Header().Set("go-web-foundation","vip")
	w.WriteHeader(418)
	w.Write([]byte("Hey,this is middlerware!"))
	w.Write(rec.Body.Bytes())
}

func myHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello world!"))
}

func main() {

	mid:=&ModifieMiddleware{http.HandlerFunc(myHandler)}
	http.ListenAndServe(":8080",mid)

}


//rec 真正的记录访问中发生的事情。自定义响应内容。