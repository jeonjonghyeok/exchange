package main

import (
	. "fmt"
)

func main(){
	var _pipe = function(){}
	var pipe = _pipe(function(a){return a+1}, function(a){return a*a})
	
	Println(pipe(1))

}
