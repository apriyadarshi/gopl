package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

//Indexing
func BuildIndexes() {
	//Build primary index
	//Build full text index
}

//Builds and stores map[num][string] string is filename
func BuildPrimaryIndex() {
	// In our case automatically done by fileName
	//1 --> 1.json
}

func GetComic(comicNo int) (Comics, error) {

	fNameFmt := os.Getenv("GOPATH") + "/src/gopl/ch4/xkcd/data/%d.json"
	fName := fmt.Sprintf(fNameFmt, comicNo)

	f, err := os.Open(fName)
	if err != nil {
		return Comics{}, err
	}

	var comic Comics
	errDecode := json.NewDecoder(f).Decode(&comic)
	if errDecode != nil {
		fmt.Printf("Error while decoding json from file %s: \n", fName)
		log.Fatal(errDecode)
		f.Close()
	}
	f.Close()
	return comic, err
}

//Builds inverted index and saves it in file
//Diff function will load it in memory each time the program starts
//Index only for full words
//Fallback should be there to search whole file in loop if no match n index
//Build Index only on 3 column data
func BuildFTSIndex() {

	baseDir := os.Getenv("GOPATH") + "/src/gopl/ch4/xkcd"

	//Get all comics
	files, _ := ioutil.ReadDir(baseDir + "/data")

	index := make(map[string]map[int]bool) //Map: word -> list of docs in which it appears
	for _, fInfo := range files {
		fName := baseDir + "/data/" + fInfo.Name()
		fmt.Printf("Building index from %s\n", fName)
		//comicNo, _ := strconv.Atoi(strings.Split(fInfo.Name(), ".")[0]) //1.go -> 1
		f, err := os.Open(fName)
		if err != nil {
			log.Fatal(err)
			f.Close()
		}

		var comic Comics
		errDecode := json.NewDecoder(f).Decode(&comic)
		if errDecode != nil {
			fmt.Printf("Error while decoding json from file %s: \n", fName)
			log.Fatal(errDecode)
			f.Close()
		}
		fmt.Println(comic.Title)
		f.Close()
		words := strings.Fields(strings.Join([]string{comic.Title, comic.SafeTitle, comic.Alt, comic.Transcript}, " "))

		for _, word := range words {
			word = getStem(word)

			if word != "" {
				if _, ok := index[word]; !ok {
					//when no key for word initialize empty map
					index[word] = make(map[int]bool)
					index[word][comic.Num] = true
				} else {
					index[word][comic.Num] = true
				}
			}
		}
	}

	//Save index in a file
	bytes, err := json.Marshal(index)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(baseDir+"/index/fts.index", bytes, 0644) //WriteFile automatically truncates file if it exists
}

func getStem(word string) string {

	//Lowercase
	word = strings.ToLower(word)

	//Remove punctuation marks
	word = strings.Trim(word, "\" !.,;:\t\r\n'`$()@#&-_[]{}%+/*")

	//Remove stop words
	if word == "the" || word == "a" || word == "and" || word == "or" || word == "he" || word == "she" || word == "they" || word == "are" || word == "is" {
		word = ""
	}

	//Set empty if word is a decimal no integer/float
	if _, err := strconv.ParseFloat(word, 64); err == nil {
		word = ""
	}
	//
	return word
}
