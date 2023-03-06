package docx_format

import "docx/docx_format/rgb"

type TextFormat struct { // rpr
	Size  TextSize  // Размер шрифта
	Font  Fonts     // Шрифт
	Style TextStyle // Стили

	Color  rgb.ColorRGB
	Shadow rgb.ColorRGB
}

func NewTextFormat() TextFormat {
	return TextFormat{
		Size:   24,
		Font:   FontCalibri,
		Style:  TextStyle{},
		Color:  rgb.ColorRGB{0, 0, 0},
		Shadow: rgb.ColorRGB{255, 255, 255},
	}
}
func (f TextFormat) SetSize(size TextSize) TextFormat {
	f.Size = size
	return f
}
func (f TextFormat) SetFont(font Fonts) TextFormat {
	f.Font = font
	return f
}
func (f TextFormat) SetStyle(s ...Style) TextFormat {
	f.Style = getStyles(s)
	return f
}
func (f TextFormat) SetColor(color rgb.ColorRGB) TextFormat {
	f.Color = color
	return f
}
func (f TextFormat) SetShadow(color rgb.ColorRGB) TextFormat {
	f.Shadow = color
	return f
}

func getStyles(s []Style) TextStyle {
	var style TextStyle
	for _, v := range s {
		switch v {
		case TextStyleBold:
			style.TextBold = true
		case TextStyleCursive:
			style.TextCursive = true
		case TextStyleUnderline:
			style.TextUnderline = true
		}
	}
	return style
}

// TextSize Размер текста
type TextSize int

// Fonts Шрифты
type Fonts int

// TextStyle Стиль текста
type TextStyle struct {
	TextBold      bool // Жирный
	TextCursive   bool // Курсив
	TextUnderline bool // Подчеркнутый
}

type Style int

const (
	TextStyleBold Style = iota
	TextStyleCursive
	TextStyleUnderline
)

const (
	FontCalibri       Fonts = iota // Calibri (Основной текст)
	FontTimesNewRoman              // TimesNewRoman
	FontWingdings2                 // Wingdings2
)
