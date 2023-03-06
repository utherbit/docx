package docx_text

type DocxText struct {
	Text string // Текст

	size  int   // Размер шрифта
	fount Fonts // Шрифт

	align   Alignment   // Выравнивание
	style   TextStyle   // Стили
	spacing TextSpacing // Отступы
}

func GetDocxText(text string) DocxText {
	return DocxText{
		Text:  text,
		size:  22,
		fount: FontCalibri,
		align: AlignLeft,
		style: TextStyle{
			TextBold: false, TextCursive: false, TextUnderline: false,
		},
		spacing: TextSpacing{
			After: 0, Before: 0, Left: 0, Right: 0,
		},
	}
}
