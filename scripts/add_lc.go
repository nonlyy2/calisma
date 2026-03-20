// this script was created to ease the way of archiving solutions and
// make data-stats for obsidian notes by making md lines for note

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	// path to obsidian markdown file
	obsidianFilePath = "/Users/assylkhan/Library/Mobile Documents/com~apple~CloudDocs/Obsidian/main/03_LeetCode/LeetCode Archive.md"

	// path to solutions
	goFilesBaseDir = "/Users/assylkhan/Desktop/neetcode-go/leetcode"
)

// parser
type LCResponse struct {
	Data struct {
		Question struct {
			QuestionFrontendId string `json:"questionFrontendId"`
			Title              string `json:"title"`
			Difficulty         string `json:"difficulty"`
			TopicTags          []struct {
				Name string `json:"name"`
			} `json:"topicTags"`
		} `json:"question"`
	} `json:"data"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run add_lc.go <slug> (e.g. two-sum)")
		os.Exit(1)
	}
	slug := os.Args[1]

	// query to GraphQL API LeetCode
	query := `{"query":"query questionData($titleSlug: String!) { question(titleSlug: $titleSlug) { questionFrontendId title difficulty topicTags { name } } }","variables":{"titleSlug":"` + slug + `"}}`

	resp, err := http.Post("https://leetcode.com/graphql", "application/json", bytes.NewBuffer([]byte(query)))
	if err != nil {
		fmt.Printf("network error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var lcData LCResponse
	json.Unmarshal(body, &lcData)

	q := lcData.Data.Question
	if q.Title == "" {
		fmt.Println("problem not found, make sure to use the slug (from the url), not the number")
		os.Exit(1)
	}

	// tags for dataview plugin
	difficultyTag := "#" + strings.ToLower(q.Difficulty)
	var tags []string
	tags = append(tags, difficultyTag)
	for _, t := range q.TopicTags {
		// space to underscore
		formattedTag := "#" + strings.ToLower(strings.ReplaceAll(t.Name, " ", "_"))

		// remove parentheses for obsidian
		formattedTag = strings.ReplaceAll(formattedTag, "(", "")
		formattedTag = strings.ReplaceAll(formattedTag, ")", "")

		tags = append(tags, formattedTag)
	}
	tagsStr := strings.Join(tags, " ")

	// formating path to exact file
	fileName := fmt.Sprintf("%s_%s.go", q.QuestionFrontendId, strings.ReplaceAll(slug, "-", "_"))
	localFilePath := fmt.Sprintf("%s/%s", goFilesBaseDir, fileName)

	// making md
	newLine := fmt.Sprintf("- [%s. %s](vscode://file%s) %s\n", q.QuestionFrontendId, q.Title, localFilePath, tagsStr)

	// write at the end
	f, err := os.OpenFile(obsidianFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("failed to open obsidian file: %v\ncheck obsidianFilePath constant\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(newLine); err != nil {
		fmt.Printf("write error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ added to obsidian:\n%s\n", newLine)
}
