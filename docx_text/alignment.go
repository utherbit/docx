package docx_text

import "docx/docx_xml"

// Alignment Выравнивание текста
type Alignment int

const (
	AlignLeft   Alignment = iota // Слева
	AlignCenter                  // По центру
	AlignRight                   // Справа
	AlignBoth                    // По ширине
)

func (a Alignment) isDefault() bool {
	return a == AlignLeft
}

func getTagAlignment(alignment Alignment) docx_xml.Tag {
	jc := docx_xml.GetTag("w:jc")
	switch alignment {
	case AlignRight:
		jc.AddAttribute("w:val", "right")
	case AlignCenter:
		jc.AddAttribute("w:val", "center")
	case AlignBoth:
		jc.AddAttribute("w:val", "both")
	}
	return jc
}
