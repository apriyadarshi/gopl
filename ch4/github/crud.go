package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const GithubAPIURL = "https://api.github.com"
const OAUTH_TOKEN = "4f0f8cd8e53b002f0731b04e5dd4cab14d3f23e1"

//Gets all issues for a particular repository
//State: 0 open; 1 closed
func GetAllIssues(owner string, repo string, state bool) (*[]*Issue, error) {

	//Make a get request
	url := strings.Join([]string{GithubAPIURL, "repos", owner, repo, "issues"}, "/")
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//Redirects not handled
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request for issuesfailed with status: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

//Gets the issue for Issue number
func GetIssue(owner string, repo string, num string) (*Issue, error) {

	//Make a get request
	url := strings.Join([]string{GithubAPIURL, "repos", owner, repo, "issues", num}, "/")
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//Redirects not handled
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request for issue no %d failed with status: %s", num, resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

//Creates an issue and responds whether the issue was successfully created or not. This is a sync call
/*Pending -- Authentication*/
func CreateIssue(owner string, repo string, issue Issue) (*Issue, error) {
	url := strings.Join([]string{GithubAPIURL, "repos", owner, repo, "issues"}, "/")

	reader := new(bytes.Buffer)
	json.NewEncoder(reader).Encode(issue)

	urlWithAuth := url + "?access_token=4f0f8cd8e53b002f0731b04e5dd4cab14d3f23e1"
	fmt.Println(urlWithAuth)
	resp, err := http.Post(urlWithAuth, "application/json; charset=utf-8", reader)

	if err != nil {
		return nil, err
	}

	//Redirects not handled
	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("Request for update failed with status: %s and response %s", resp.Status, body)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func EditIssue(owner string, repo string, issue *Issue) (*Issue, error) {

	numStr := strconv.Itoa(issue.Number)
	url := strings.Join([]string{GithubAPIURL, "repos", owner, repo, "issues", numStr}, "/")
	url += "?access_token=4f0f8cd8e53b002f0731b04e5dd4cab14d3f23e1"

	reader := new(bytes.Buffer)
	json.NewEncoder(reader).Encode(issue)

	//New Http Client required as PATCH method is not supported directly
	//Also OAuth token that was generated manually will be sent in header for better safety
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPatch, url, reader)

	req.Header.Add("Authorization", "token "+OAUTH_TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("Method \"Delete Github Issue \" failed:\n \twith status : %s \n\tmessage : %v", resp.Status, body)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil

}

//For delete simply get and update status = Closed
