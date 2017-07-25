package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main()  {
	http.HandleFunc("/",hey)
	http.ListenAndServe(":8081",nil)
}

const tpl  =`
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<from method="post" action="/">
		Username: <input type="text" name="uname">
		Passwotd: <input type="password" name="pwd">
		<button type="submit">Submit</button>
		</from>
	</body>
</html>
`


func hey(w http.ResponseWriter, r *http.Request)  {
	if r.Method=="GET"{
	t:=template.New("hey")
		t.Parse(tpl)
		t.Execute(w,nil)
	}else{
		fmt.Print(r.FormValue("uname"))
	}
}