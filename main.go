package main

import (
	"fmt"
	"embed"
	"io/ioutil"
	"io/fs"

)

//go:embed version.txt
var version string

//go:embed gambar.png
var gambar []byte

//go:embed files/*.txt
var path embed.FS

func main(){
	fmt.Println(version)
	err  :=ioutil.WriteFile("new_gambar.png",gambar, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/"+ entry.Name())
			fmt.Println(string(file))
		}
	}

}