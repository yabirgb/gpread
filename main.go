package main
import (
	"fmt"
	"os"
)

func readOrCreate(string path){
	var _, err = os.Stat(path)

	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if isError(err){
			fmt.Println("Couldn't create database file at %s", path)
		}
		defer file.Close()
	}

	fmt.Println("Database file created successfully at %s", path)
}

func main(){
	data, err :=
}
