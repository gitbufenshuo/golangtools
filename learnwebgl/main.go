package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	thefilename := os.Getenv("FILENAME")
	fmt.Println(thefilename)
	segs := strings.Split(thefilename, ".")
	frontName := segs[0]
	suffixName := segs[1]
	var fileout string
	var url string
	if suffixName == "vert" || suffixName == "frag" {
		os.Mkdir("shaders", 0755)
		fileout = path.Join("shaders", thefilename)
		url = "http://learnwebgl.brown37.net/lib/shaders/" + thefilename
	} else {
		os.Mkdir("models", 0755)
		fileout = path.Join("models", thefilename)
		url = "http://learnwebgl.brown37.net/lib/models/" + thefilename
	}
	//////////////////////////////////////////////////
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(fileout, data, 0644)
}
