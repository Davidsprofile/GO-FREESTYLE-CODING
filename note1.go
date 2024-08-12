package main

import (
	"bufio"
	"fmt"
	"go_course/note"
	"os"
	"strings"
)


func main () {
	title, content := getNoteData()
	userNote,err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return 
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("save failed")
		return
	}
	fmt.Println("Save succedded")
}


func getNoteData () (string, string) {
	title := getUserInput("Title: ")
	content := getUserInput("Content: ")
	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)  // GO dont want to accept input normaly!!!!!!!!!!!!!
	text,err := reader.ReadString('\n') // identifier when it should stop reading
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n") // we tell to remove \n
	text = strings.TrimSuffix(text, "\r")
	return text
}
