package main

import (
	"net/http"
	//"log"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("http://www.mcleods.com/")
	//if err != nil {
	//	log.Fatal(err)
	//}
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("%s", page)
}
