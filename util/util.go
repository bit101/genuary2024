// Package util is utils
package util

import cairo "github.com/bit101/blcairo"

func Stampit(context *cairo.Context, text string) {
	h := context.Height
	context.Save()
	context.SetFontSize(14)
	context.SetSourceBlack()
	context.FillText(text, 5, h-5)
	context.SetSourceGray(0.75)
	context.FillText(text, 4, h-6)
	context.Restore()

}
