package main

import (
	"fmt"
	"gopl/ch4/xkcd"
	"os"
)

func main() {
	//xkcd.DownloadAll()
	//xkcd.BuildFTSIndex()
	term := os.Args[1]
	comics := xkcd.Search(term)
	if len(comics) == 0 {
		fmt.Println("No matching comics found")
	} else {
		fmt.Printf("%d matching results :\n", len(comics))
		for i, comic := range comics {
			fmt.Printf("%d.\t", i)
			fmt.Printf("Number:\t%d\n \tTitle:\t%s\n \tURL:\t%s\n \tTranscript:\t%s\n\n", comic.Num, comic.Title, comic.Link, comic.Transcript)
		}
	}
}

/*result, err := github.GetIssue(os.Args[1], os.Args[2], os.Args[3])
if err != nil {
	log.Fatal(err)
}
fmt.Println("IssueNo\tTitle\tStatus")
fmt.Printf("%d\t%s\t%s\n", result.Number, result.Title, result.State)
*/
/*issue := github.Issue{Title: os.Args[3]}
//issue.Title = os.Args[3]
result, err := github.CreateIssue(os.Args[1], os.Args[2], issue)
if err != nil {
	log.Fatal(err)
}
fmt.Println("IssueNo\tTitle\tStatus")
fmt.Printf("%d\t%s\t%s\n", result.Number, result.Title, result.State)
/*for _, v := range *result {
	fmt.Printf("%d\t%s\t%s\n", v.Number, v.Title, v.State)
}*/
/*
	//Test issue update
	owner := os.Args[1]
	repo := os.Args[2]
	issueNum := os.Args[3]
	newTitle := os.Args[4]

	//Get Issue
	issue, err := github.GetIssue(os.Args[1], os.Args[2], issueNum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updating Issue with Issue No: %d and title %s", issue.Number, issue.Title)

	issue.Title = newTitle

	result, err := github.EditIssue(owner, repo, issue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IssueNo\tNew Title\tStatus")
	fmt.Printf("%d\t%s\t%s\n", result.Number, result.Title, result.State)

}*/

/*result, err := github.SearchIssues(os.Args[1:])
if err != nil {
	log.Fatal(err)
}
fmt.Printf("%d issues:\n", result.TotalCount)
issuesByAge := make(map[string][]string)
for _, item := range result.Items {
	//fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, daysGone(item.CreatedAt))
	issuesByAge[daysGone(item.CreatedAt)] = append(issuesByAge[daysGone(item.CreatedAt)], fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title))
}

for age, issues := range issuesByAge {
	fmt.Printf(" Age of Issues : %v \n", age)
	fmt.Println("-----------------------")
	for i := range issues {
		fmt.Println(i)
	}
	fmt.Println("-----------------------")
	fmt.Println("")
}*/
/*fmt.Printf("Input= %d; Output=%d", 34000, ch3.Comma2("34000"))
//fmt.Printf("Input= %d; Output=%d", 34000.5467, ch3.Comma2("34000.5467"))
reader := bufio.NewReader(os.Stdin)
var text string
for {
	fmt.Print("Enter strings seperated by commas whose sha256 hash are to be compared : ")
	text, _ = reader.ReadString('\n')
	words := strings.Split(text[:len(text)-1], ",")

	fmt.Println(ch4.DiffBits(sha256.Sum256([]byte(words[0])), sha256.Sum256([]byte(words[1]))))
}*/
/*wordfreq := ch4.WordFreq("/home/skywalker/Documents/go/src/gopl/ch4/inputforwordfreq.txt")
fmt.Println("Word\tCount")
for k, v := range wordfreq {
	fmt.Printf("%v\t%d\n", k, v)
}*/

/*
func daysGone(t time.Time) string {
	days := int(time.Since(t).Hours() / 24)

	switch {
	case days <= 30:
		return "<1m"
	case days > 30 && days <= 365:
		return "<1y"
	case days > 365:
		return ">1y"
	}
	return ""
}*/
