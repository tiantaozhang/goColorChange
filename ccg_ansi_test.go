// +build !windows

package goColorChange

import (
	"fmt"
	"testing"
)

func TestAnsiText(t *testing.T) {

	// fmt.Printf("none %s", ansiText(None, false, None, false))
	// fmt.Printf("red %s", ansiText(Red, false, None, false))
	// fmt.Printf("red %s", ansiText(Red, true, None, false))
	// fmt.Printf("none%s", ansiText(None, false, Green, false))
	// fmt.Printf("red %s", ansiText(Red, false, Green, false))
	// fmt.Printf("red %s", ansiText(Red, true, Green, false))
	// fmt.Println("")

	Foreground(Green, false)
	fmt.Println("Green text starts here...")
	ChangeColor(Red, true, None, false)
	fmt.Println("red foreground")
	ResetColor()

}

func TestAnsiText2(t *testing.T) {
	ResetColor()

	ChangeColorAndStyle(Underline, Green, None)
	fmt.Println("underline green white")
	LogToFile("", "underline %v %v with lines,files", "green", "white")
	LogToFile("1.txt", "Flicker %v %v wieh lines,files", "blue", "none")
	ResetColor()
	ChangeColorAndStyle(Flicker, Blue, None)
	fmt.Println("Flicker Blue none")
	LogToFile("", "Flicker %v %v wieh lines,files", "blue", "none")
	ResetColor()
}
