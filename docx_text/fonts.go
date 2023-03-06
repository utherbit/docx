package docx_text

import "docx/docx_xml"

// Fonts Шрифты
type Fonts int

const (
	FontCalibri       Fonts = iota // Calibri (Основной текст)
	FontTimesNewRoman              // TimesNewRoman
	FontWingdings2                 // Wingdings2
)

func (font Fonts) isDefault() bool {
	return font == FontCalibri
}

func (font Fonts) getFountName() string {
	switch font {
	case FontCalibri:
		return `Calibri`
	case FontTimesNewRoman:
		return `Times New Roman`
	case FontWingdings2:
		return `Wingdings 2`
	}
	return ""
}

func getTagFont(font Fonts) docx_xml.Tag {
	fontName := font.getFountName()
	rFonts := docx_xml.GetTag("w:rFonts")

	rFonts.AddAttribute("w:ascii", fontName)
	rFonts.AddAttribute("w:hAnsi", fontName)

	return rFonts
}
