// Exercise 4.11 in The Go Programming Language book
// Simple utility CLI program, perfrom CRUD operations on Github issues
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var repoOwner = flag.String("o", " ", "the username of the repo owner")
var issueAction = flag.String("a", " ", "the action to be performed (create, update, read, delete)")
var repoName = flag.String("n", " ", "the repository name")

const APIURL = "https://api.github.com/repos/"

// TODO: Separate the structs to live in a different file.
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	// TODO: Parsing input from the command line
	// TODO: Refactor as a different function
	flag.Parse()
	// TODO: Question: Should there be one instance of Scanner passed around or should there be one spawned each time a function executes?
	// input := bufio.NewScanner(os.Stdin)

	// Adding comments with my thoughts, since it's just play code
	// I am using single Issue pointer because I expect every single action to print the issue at the end of it.
	// It should reduce the amount of code in individual functions and the "side effect" of printing
	// On the other hand it makes the switch ugly, it should just invoke the function
	// Consider moving the code to a helper function
	var result *Issue

	switch *issueAction {
	case "create":
		createIssue(*repoOwner, *repoName)
	case "read":
		var err error
		result, err = readIssue(*repoOwner, *repoName)
		if err != nil {
			fmt.Printf("Process has exited when trying to read the issue: %v", err)
			os.Exit(1)
		}
		fmt.Println(result.CreatedAt.Clock())
	case "update":
		updateIssue(*repoOwner, *repoName)
	case "delete":
		deleteIssue(*repoOwner, *repoName)
	}

	fmt.Printf("#%-6d \nLogin: %9.9s \nTitle: %-55s \nBody:\n %s\n",
		result.Number, result.User.Login, result.Title, result.Body)

}

func createIssue(repoOwner, repoName string) (IssuesSearchResult, error) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input the title of the issue you want to create: ")
	input.Scan()
	issueTitle := input.Text()
	fmt.Println(issueTitle)

	fmt.Printf("Please input the body of the issue you want to create: ")
	input.Scan()
	issueBody := input.Text()
	fmt.Println(issueBody)

	return IssuesSearchResult{}, nil
}

func readIssue(repoOwner, repoName string) (*Issue, error) {
	// TODO: Querying the Github API get an issue
	// https://github.com/projectcalico/calicoctl
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input the number of the issue you want to read: ")
	input.Scan()
	// TODO: Question: Is there Go package which validates strings which are supposed to be numbers?
	// Could be done with regex I guess, probably better approach than Parsing an int to verify it is correct
	issueNumber := input.Text()

	formattedUrl := APIURL + repoOwner + "/" + repoName + "/issues/" + issueNumber
	resp, err := http.Get(formattedUrl)
	if err != nil {
		fmt.Printf("Something is wrong: %v", err)
		os.Exit(1)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func updateIssue(repoOwner, repoName string) (IssuesSearchResult, error) {
	return IssuesSearchResult{}, nil
}

func deleteIssue(repoOwner, repoName string) (IssuesSearchResult, error) {
	return IssuesSearchResult{}, nil
}
