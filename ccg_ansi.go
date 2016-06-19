// +build !windows

package goColorChange

import (
	"fmt"
	"os"
	"strconv"
)

func isDumbTerm() bool {
	return os.Getenv("TERM") == "dumb"
}

func resetColor() {
	if isDumbTerm() {
		return
	}
	fmt.Print("\x1b[0m")
}

// style Style
func ansiText2(style Style, fg Color, bg Color) string {
	fmt.Println("=====================")
	if fg == None && bg == None {
		return ""
	}
	s := []byte("\x1b[")
	if style != Default {
		s = strconv.AppendUint(s, (uint64)(style-Default), 10)
	} else {
		s = strconv.AppendUint(s, (uint64)(Default), 10)
	}
	// fmt.Printf("s:%v\n", string(s))
	if bg != None {
		s = strconv.AppendUint(append(s, ";"...), 40+(uint64)(bg-Black), 10)
	}
	// fmt.Printf("s bg:%v\n", string(s))
	if fg != None {
		s = strconv.AppendUint(append(s, ";"...), 30+(uint64)(fg-Black), 10)
	}
	s = append(s, "m"...)
	// fmt.Printf("s fg:%v\n", string(s))
	return string(s)
}

func ansiText(fg Color, fgBright bool, bg Color, bgBright bool) string {
	if fg == None && bg == None {
		return ""
	}
	s := []byte("\x1b[0")
	// fmt.Printf("s:%v\n", string(s))
	if fg != None {
		s = strconv.AppendUint(append(s, ";"...), 30+(uint64)(fg-Black), 10)
		if fgBright {
			s = append(s, ";1"...)
		}
	}
	// fmt.Printf("s fg:%v\n", string(s))
	if bg != None {
		s = strconv.AppendUint(append(s, ";"...), 40+(uint64)(bg-Black), 10)
	}
	s = append(s, "m"...)
	// fmt.Printf("s bg:%v\n", string(s))
	return string(s)
}

func changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	if isDumbTerm() {
		return
	}
	if fg == None && bg == None {
		return
	}

	n, _ := fmt.Print(ansiText(fg, fgBright, bg, bgBright))
	fmt.Printf("n:%v\n", n)
}

func changeColorAndStyle(style Style, fg Color, bg Color) {
	if isDumbTerm() {
		return
	}
	if fg == None && bg == None {
		return
	}

	fmt.Print(ansiText2(style, fg, bg))
}
