// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type DatedIssues struct {
	MonthAgo       []*Issue
	ThreeMonthsAgo []*Issue
	SixMonthsAgo   []*Issue
	YearAgo        []*Issue
	OlderThanYear  []*Issue
}

//!+
// Exercise 4.10
func main() {
	datedIssues := new(DatedIssues)

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	t := time.Now()
	fmt.Println(t)
	t2 := t.AddDate(0, -1, 0)
	fmt.Println(t)
	fmt.Println(t2)
	currentTime := time.Now()
	lastMonth := currentTime.AddDate(0, -1, 0)
	threeMonthsAgo := currentTime.AddDate(0, -3, 0)
	sixMonthsAgo := currentTime.AddDate(0, -6, 0)
	yearAgo := currentTime.AddDate(-1, 0, 0)

	// I know what you are thinking, and I think the very same thing.
	// This looks like an awful piece of code which could be refactored as a helper function
	// And you are right, but it's just an exercise.
	// TODO: Refactor all those loops and if statements into a function.

	// It seems to me after consideration that Map might be better suited for
	// this kind of data than a struct.
	// TODO: Try replacing the struct with a map.
	for _, item := range result.Items {
		if item.CreatedAt.After(lastMonth) {
			datedIssues.MonthAgo = append(datedIssues.MonthAgo, item)
		} else if item.CreatedAt.After(threeMonthsAgo) {
			datedIssues.ThreeMonthsAgo = append(datedIssues.MonthAgo, item)
		} else if item.CreatedAt.After(sixMonthsAgo) {
			datedIssues.SixMonthsAgo = append(datedIssues.MonthAgo, item)
		} else if item.CreatedAt.After(yearAgo) {
			datedIssues.YearAgo = append(datedIssues.MonthAgo, item)
		} else {
			datedIssues.OlderThanYear = append(datedIssues.OlderThanYear, item)
		}
	}

	fmt.Printf("\n\n\n# Issues created less than a month ago #\n\n\n")
	for _, item := range datedIssues.MonthAgo {
		fmt.Printf("#%-6d %9.9s %-55.45s %-20s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\n\n\n# Issues created less than three months ago #\n\n\n")
	for _, item := range datedIssues.ThreeMonthsAgo {
		fmt.Printf("#%-6d %9.9s %-55.45s %-20s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\n\n\n# Issues created less than six months ago #\n\n\n")
	for _, item := range datedIssues.SixMonthsAgo {
		fmt.Printf("#%-6d %9.9s %-55.45s %-20s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\n\n\n# Issues created less than year ago #\n\n\n")
	for _, item := range datedIssues.YearAgo {
		fmt.Printf("#%-6d %9.9s %-55.45s %-20s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\n\n\n# Issues created more than a year ago #\n\n\n")
	for _, item := range datedIssues.OlderThanYear {
		fmt.Printf("#%-6d %9.9s %-55.45s %-20s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func datedCategoriesPrinter(issues *DatedIssues, category string) {
	fmt.Printf("\n\n\n# Issues created more than a year ago #\n\n\n")
	for _, item := range issues.OlderThanYear {
		fmt.Printf("#%-6d %9.9s %-55.45s %-20s\n",
			item.Number, item.User.Login, item.Title)
	}
}
