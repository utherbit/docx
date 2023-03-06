package docx_models

import (
	"docx/docx_format"
)

type TextModel struct {
	Text string // Текст

	format docx_format.TextFormat
}

func (t *TextModel) Format(format docx_format.TextFormat) *TextModel {
	t.format = format
	return t
}

//
//
//
//

func getFountName(font docx_format.Fonts) string {
	switch font {
	case docx_format.FontCalibri:
		return `Calibri`
	case docx_format.FontTimesNewRoman:
		return `Times New Roman`
	case docx_format.FontWingdings2:
		return `Wingdings 2`
	}
	return ""
}
