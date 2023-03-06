package docx_format

type ParagraphFormat struct { // ppr
	Align   Alignment // Выравнивание
	Spacing Spacing   // Отступы

	Numbering Numbering // маркеры, нумерация, списки todo кастомизация стилей для списков
}

func NewParagraphFormat() ParagraphFormat {
	return ParagraphFormat{
		Align:   AlignLeft,
		Spacing: Spacing{FirstLine: 0, Hanging: 0, Left: 0, Right: 0, Before: 0, After: 8},
	}
}
func (f ParagraphFormat) SetAlign(align Alignment) ParagraphFormat {
	f.Align = align
	return f
}
func (f ParagraphFormat) SetSpacing(spacing Spacing) ParagraphFormat {
	f.Spacing = spacing
	return f
}
func (f ParagraphFormat) SetNumbering(numbering Numbering) ParagraphFormat {
	f.Numbering = numbering
	return f
}

// Alignment Выравнивание текста
type Alignment int

const (
	AlignLeft   Alignment = iota // Слева
	AlignCenter                  // По центру
	AlignRight                   // Справа
	AlignBoth                    // По ширине
)

// Spacing отступы и интервалы
type Spacing struct {
	FirstLine int // Отступ первой строки
	Hanging   int // Выступ первой строки

	Left   int // Отступ слева
	Right  int // Отступ справа
	Before int // Интервал перед
	After  int // Интервал после
}

type Numbering struct {
	ILvl  int // Уровень глубины вложенности 1-9
	NumId int // todo ID стиля из файла numbering
}
