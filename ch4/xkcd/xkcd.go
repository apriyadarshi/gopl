//Downloads and builds indexes locally
//Search the index and print its url and transcript for search term on command line
//1 -> 1833 comics
package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const XkcdURLFmt = "https://xkcd.com/%d/info.0.json"

type Comics struct {
	Month      int `json:",string"`
	Num        int
	Link       string
	Year       int `json:",string"`
	News       string
	SafeTitle  string `json:safe_title`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        int `json:",string"`
}

//Only full word search allowed
func Search(term string) []*Comics {

	//Put index in memory
	ftsIndexFName := os.Getenv("GOPATH") + "/src/gopl/ch4/xkcd" + "/index/fts.index"

	ftsIndexF, err := os.Open(ftsIndexFName)
	if err != nil {
		log.Fatal(err)
	}

	var ftsIndex map[string]map[int]bool
	if err := json.NewDecoder(ftsIndexF).Decode(&ftsIndex); err != nil {
		fmt.Println("Error while unmarshalling Index:")
		log.Fatal(err)
	}

	comicNums, ok := ftsIndex[term]

	var matches []*Comics
	if ok {
		for comicNum := range comicNums {

			comic, errComic := GetComic(comicNum)
			if errComic != nil {
				fmt.Println("Error while getting comic no : " + strconv.Itoa(comicNum) + " with error: ")
				log.Fatal(errComic)
			}

			matches = append(matches, &comic)
		}
		return matches

	} else {
		return []*Comics{}
	}
}

//Downloads the JSON objects for all comics and store it locally in seperate files in folder data
//If already downloaded, doesn't download a file that has already been downloaded
func DownloadAll() {

	fNameFmt := os.Getenv("GOPATH") + "/src/gopl/ch4/xkcd/data/%d.json"

	for i := 1; i < 1835; i++ {

		fName := fmt.Sprintf(fNameFmt, i)

		if _, err := os.Stat(fName); err == nil {
			//If file exists, skip loop
			continue
		}

		url := fmt.Sprintf(XkcdURLFmt, i)

		fmt.Printf("Downloading File %s \n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error Downloading %s : %s", fName, err)
		}

		if resp.StatusCode == http.StatusNotFound {
			continue
		}

		bytes, _ := ioutil.ReadAll(resp.Body)

		errWrite := ioutil.WriteFile(fName, bytes, 0644)
		if errWrite != nil {
			fmt.Printf("Error writing %s : %s", fName, errWrite)
		} else {
			fmt.Printf("Downloaded.")
		}

		resp.Body.Close()
	}
}
