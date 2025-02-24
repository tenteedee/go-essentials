package struct_practice

import (
	"fmt"

	"github.com/tenteedee/go-essentials/utils"
)

func Main() {
	title := utils.GetInputString("Note Title: ")
	content := utils.GetInputString("Note Content: ")

	note, err := NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(note.String())

	err = note.SaveToFile()
	if err != nil {
		fmt.Println(err)
		return
	}
}
