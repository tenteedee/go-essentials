package struct_practice

import (
	"fmt"

	"github.com/tenteedee/go-essentials/utils"
)

type Saver interface {
	SaveToFile() error
}

// type Displayer interface {
// 	String() string
// }

type Outputtable interface {
	Saver
	Display()
}

func Main() {
	title := utils.GetInputString("Note Title: ")
	content := utils.GetInputString("Note Content: ")

	note, err := NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)
}

func saveData(s Saver) error {
	err := s.SaveToFile()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}

func outputData(data Outputtable) {
	data.Display()
	err := saveData(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
