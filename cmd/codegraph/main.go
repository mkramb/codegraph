package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

type MatchInfo struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	FilePath string `json:"filepath"`
	Position struct {
		Row    int `json:"row"`
		Column int `json:"column"`
	} `json:"position"`
}

var queryStr = `
(
	(comment) @comment
	(#match? @comment "// @code-graph-(label|link)/.+")
)
(
	(comment) @comment
	(#match? @comment "/\\*([\\s\\S]*?)@code-graph-(label|link)/[\\s\\S]*?\\*/")
)
`

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <path_to_js_file1> <path_to_js_file2> ...")
	}

	var allResults []MatchInfo

	for _, filePath := range os.Args[1:] {
		results, err := parseFile(filePath)

		if err != nil {
			log.Printf("Error parsing file %s: %s", filePath, err)
			continue
		}

		allResults = append(allResults, results...)
	}

	if err := printJSONOutput(allResults); err != nil {
		log.Fatalf("Error printing JSON output: %s", err)
	}
}

func parseFile(filePath string) ([]MatchInfo, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	ctx := context.Background()
	lang := javascript.GetLanguage()

	rootNode, err := sitter.ParseCtx(ctx, content, lang)

	if err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}

	results, err := extractMatches(rootNode, content, filePath, lang)

	if err != nil {
		return nil, fmt.Errorf("failed to extract matches: %w", err)
	}

	return results, nil
}

func extractMatches(rootNode *sitter.Node, content []byte, filePath string, lang *sitter.Language) ([]MatchInfo, error) {
	q, err := sitter.NewQuery([]byte(queryStr), lang)

	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(q, rootNode)

	var results []MatchInfo
	re := regexp.MustCompile(`@code-graph-(label|link)/([^\s*/]+)`)

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, capture := range match.Captures {
			commentContent := capture.Node.Content(content)
			startPosition := capture.Node.StartPoint()

			lines := strings.Split(commentContent, "\n")

			for _, line := range lines {
				line = strings.TrimSpace(line)
				matches := re.FindStringSubmatch(line)

				if len(matches) < 3 {
					continue
				}

				matchInfo := MatchInfo{
					Type:     matches[1],
					Value:    matches[2],
					FilePath: filePath,
				}

				matchInfo.Position.Row = int(startPosition.Row + 1)
				matchInfo.Position.Column = int(startPosition.Column + 1)

				results = append(results, matchInfo)

				startPosition.Row++
				startPosition.Column = 0
			}
		}
	}

	return results, nil
}

func printJSONOutput(results []MatchInfo) error {
	jsonOutput, err := json.Marshal(results)

	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	fmt.Println(string(jsonOutput))

	return nil
}
