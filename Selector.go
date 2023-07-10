package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"github.com/charmbracelet/lipgloss"
)

func main(){
	
	//Here you Can config The Desing of the Header
	var headerStyle =lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()). 
	Background(lipgloss.Color("#D13D73")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Align(lipgloss.Center).BorderStyle(lipgloss.RoundedBorder()).Blink(true).Bold(true). 
	Italic(true).MarginLeft(10)

	//Here you Can config The Desing of the Body
	var bodyStyle = lipgloss.NewStyle().BorderTop(false).BorderBottom(false).BorderLeft(true).BorderRight(true).
	//Background(lipgloss.Color("#D13D73")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Align(lipgloss.Center).BorderStyle(lipgloss.DoubleBorder()).Italic(true).Width(50)
	
	var Topp = lipgloss.NewStyle().BorderTop(true).BorderBottom(false).BorderLeft(true).BorderRight(true).
	Align(lipgloss.Center).BorderStyle(lipgloss.DoubleBorder()).Width(50)

	var Bot = lipgloss.NewStyle().BorderTop(false).BorderBottom(true).BorderLeft(true).BorderRight(true).
	Align(lipgloss.Center).BorderStyle(lipgloss.DoubleBorder()).Width(50)



	//Functionality

	currentUser, err := user.Current()
    if err != nil {
        panic(err)
    }

	absolutePath:= currentUser.HomeDir+ "/.config/gopsalm/Display/"
	directory, err := os.ReadDir(absolutePath)

	if err != nil {
      panic("Error opening the directory")
	}
	
	amountOfFiles :=len(directory)
	fileToOpen := rand.Int31n(int32(amountOfFiles))
	
	fileName := directory[fileToOpen].Name()
	finalPath:= absolutePath+fileName

	file, err := os.Open(finalPath)
    if err != nil {
		panic("I couldn't open the file :,c")
    }
	
	defer file.Close()

	reader := bufio.NewScanner(file)
	
	reader.Scan()
	fmt.Println(headerStyle.Render(reader.Text()))
	fmt.Println(Topp.Render(""))
	for reader.Scan(){
		fmt.Println(bodyStyle.Render(reader.Text()))
	}
	fmt.Println(Bot.Render(""))
	



	
}
