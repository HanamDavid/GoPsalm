package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"github.com/charmbracelet/lipgloss"
)

func main(){
	currentUser, err := user.Current()
    if err != nil {
        panic(err)
    }

	absolutePath:= currentUser.HomeDir+ "/.config/gopsalm/"
	directory, err := os.ReadDir(absolutePath)

	if err != nil {
      panic(err)
	}
	
	amountOfFiles :=len(directory)
	fileToOpen := rand.Int31n(int32(amountOfFiles))
	
	fileName := directory[fileToOpen].Name()
	finalPath:= absolutePath+fileName
	

	var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(2).
    PaddingLeft(4).
    Width(22)

	fmt.Println(style.Render(finalPath))



}
