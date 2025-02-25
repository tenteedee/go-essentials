package struct_practice

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content are required fields")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf(`
Note Details
Title: %s
Content: %s
Created at: %s
`, n.Title, n.Content, n.CreatedAt.Format(time.RFC1123))
}

func (note Note) SaveToFile() error {
	filename := "struct_practice/" + strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	jsonData, err := json.MarshalIndent(note, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
