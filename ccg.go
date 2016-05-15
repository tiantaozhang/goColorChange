/*
ccg package provides functions to change the color of console text.

Under windows platform, the Console API is used. Under other systems, ANSI text mode is used.
*/
package ccg

// Color is the type of color to be set.
type Color int

const (
	// No change of color
	None = Color(iota)
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
type Style int

const (
	Default   = Style(iota)
	HighLight = Style(1)
	Underline = Style(4)
	Flicker   = Style(5)
	AntiWhite = Style(7)
	Invisible = Style(8)
)

// ResetColor resets the foreground and background to original colors
func ResetColor() {
	resetColor()
}

// ChangeColor sets the foreground and background colors. If the value of the color is None,
// the corresponding color keeps unchanged.
// If fgBright or bgBright is set true, corresponding color use bright color. bgBright may be
// ignored in some OS environment.
func ChangeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	changeColor(fg, fgBright, bg, bgBright)
}

// Foreground changes the foreground color.
func Foreground(cl Color, bright bool) {
	ChangeColor(cl, bright, None, false)
}

// Background changes the background color.
func Background(cl Color, bright bool) {
	ChangeColor(None, false, cl, bright)
}

func ChangeColorAndStyle(style Style, fg Color, bg Color) {
	changeColorAndStyle(style, fg, bg)
}
