package main
import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"time"
	"errors"
	"io/ioutil"
)

const FILENAME string = ".gqread.yaml"

type Document struct{
	Url string
	Added int64 // date of the addition
	Read bool // if the url has been read or not
	Hn string // link to hn comments if exists
}

type T struct{
	Articles []Document
}

var noDocuments = errors.New("No se ha encontrado ningun archivo")
var everyThingRead = errors.New("All urls have been read")

func addToDocument(url string, data *T){
	newArticle := Document {
		Url: url,
		Added: time.Now().Unix(),
		Read: false,
		Hn: "",
	}

	data.Articles = append([]Document{newArticle},data.Articles...)
}

func readFromDocument(data *T) (Document, error){
	
	if len(data.Articles) == 0 {
		return Document{}, noDocuments
	}
	// find first not read
	for i, art := range data.Articles{
		if art.Read == false{
			return data.Articles[i], nil
		}
	}

	return Document{}, everyThingRead
}

func markLastAsRead(data *T) error{
	if len(data.Articles) == 0 {
		return noDocuments
	}
	// find first not read
	for i, art := range data.Articles{
		if art.Read == false{
			data.Articles[i].Read = true
			return nil
		}
	}

	return everyThingRead
}

func main(){

	home, homerr := os.UserHomeDir()

	if homerr != nil{
		fmt.Println("Couldn't find user home directory")
		return
	}

	FILEPATH := home + "/" + FILENAME
	
	

	file, errReadingFile := ioutil.ReadFile(FILEPATH)
	
	if errReadingFile != nil{
		fmt.Println("Dabase not found. Creating file at", FILEPATH)
		err := ioutil.WriteFile(FILEPATH, []byte("articles:"), 0755)
		if err != nil{
			fmt.Println(err)
			return
		}
	}

	var data T
	err := yaml.Unmarshal(file, &data)

	if err != nil{
		fmt.Println("Error reading database file")
		return 
	}

	if len(os.Args) == 3 || len(os.Args) == 2{
		// In this case we want to store an url
		if os.Args[1] == "add"{
			addToDocument(os.Args[2], &data)
		}else if os.Args[1] == "pop"{
				art, err := readFromDocument(&data)
 
				if err != nil{
					fmt.Println(err)
				}
				article, err := yaml.Marshal(art)
				fmt.Println(string(article))
		}else if os.Args[1] == "read"{
			modifyReadError := markLastAsRead(&data)
			if modifyReadError != nil{
				fmt.Println(modifyReadError)
			}
		}else{
			fmt.Println("Unknown argument ", os.Args[1])
		}
	} else {
		fmt.Println("Incorrect number of arguments")
	}

	marshaled, _ := yaml.Marshal(&data)
	//fmt.Println(string(marshaled))
	
	werror := ioutil.WriteFile(FILEPATH, marshaled, 0744)

	if werror != nil {
		fmt.Println(werror)
	}
}
