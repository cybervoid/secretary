package main

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func notion(token string) {
	// Create Notion client
	client := notionapi.NewClient(notionapi.Token(token))

	// Fetch a page (meta info, but not content)
	page, err := client.Page.Get(context.Background(), notionapi.PageID("101d941266e7808ea147c6f2b07d4b07"))
	if err != nil {
		fmt.Println("Error occurred while fetching page:", err)
		return
	}

	// Fetch page content (blocks)
	blocks, err := client.Block.GetChildren(context.Background(), notionapi.BlockID(page.ID), nil)
	if err != nil {
		fmt.Println("Error occurred while fetching blocks:", err)
		return
	}

	// Iterate over blocks and print their content
	for _, block := range blocks.Results {
		switch b := block.(type) {
		case *notionapi.ParagraphBlock:
			fmt.Println("Paragraph:", getText(b.Paragraph.RichText))
		case *notionapi.Heading1Block:
			fmt.Println("Heading 1:", getText(b.Heading1.RichText))
		case *notionapi.Heading2Block:
			fmt.Println("Heading 2:", getText(b.Heading2.RichText))
		case *notionapi.Heading3Block:
			fmt.Println("Heading 3:", getText(b.Heading3.RichText))
		default:
			fmt.Printf("Other block type: %T\n", b)
		}
	}
}

// Helper function to extract text from RichText array
func getText(richTexts []notionapi.RichText) string {
	var result string
	for _, rt := range richTexts {
		result += rt.Text.Content
	}
	return result
}
