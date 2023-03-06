package docx

import "fmt"

// GetXmlText Default
func GetXmlText(text string) XmlText {
	return XmlText{
		Text:  text,
		size:  22,
		fount: FontCalibri,
		align: AlignLeft,
		style: TextStyle{
			TextBold:      false,
			TextCursive:   false,
			TextUnderline: false,
		},
		spacing: TextSpacing{
			After:  0,
			Before: 0,
			Left:   0,
			Right:  0,
		},
	}
}

func (xml XmlText) SetTextXmlText(text string) XmlText {
	xml.Text = text
	return xml
}
func (xml XmlText) SetSizeXmlText(size int) XmlText {
	xml.size = size
	return xml
}
func (xml XmlText) SetFontXmlText(font Fonts) XmlText {
	xml.fount = font
	return xml
}
func (xml XmlText) SetAlignXmlText(align Alignment) XmlText {
	xml.align = align
	return xml
}
func (xml XmlText) SetStyleXmlText(style TextStyle) XmlText {
	xml.style = style
	return xml
}
func (xml XmlText) SetSpacingXmlText(spacing TextSpacing) XmlText {
	if xml.spacing.Left != spacing.Left {
		xml.spacing.Left = spacing.Left
	}
	if xml.spacing.Right != spacing.Right {
		xml.spacing.Right = spacing.Right
	}
	if xml.spacing.Before != spacing.Before {
		xml.spacing.Before = spacing.Before
	}
	if xml.spacing.After != spacing.After {
		xml.spacing.After = spacing.After
	}
	return xml
}

type XmlTextA []XmlText // Смешанный текст (с разными стилями)
type XmlText struct {
	Text    string      // Текст
	size    int         // Размер шрифта
	fount   Fonts       // Шрифт
	align   Alignment   // Выравнивание
	style   TextStyle   // Стили
	spacing TextSpacing // Отступы
}

type TextSpacing struct {
	Left   int // Отступ слева
	Right  int // Отступ справа
	Before int // Интервал перед
	After  int // Интервал после
}

type TextStyle struct {
	TextBold      bool // Жирный
	TextCursive   bool // Курсив
	TextUnderline bool // Подчеркнутый
}

type Fonts int

const (
	FontCalibri       Fonts = iota // Calibri (Основной текст)
	FontTimesNewRoman              // TimesNewRoman
	FontWingdings2                 // Wingdings2
)

type Alignment int

const (
	AlignLeft   Alignment = iota // Слева
	AlignCenter                  // По центру
	AlignRight                   // Справа
	AlignBoth                    // По ширине
)

// GetXmlTable default
func GetXmlTable(x, y int) XmlTable {

	var vh []int
	var t [][]XmlText
	var vw [][]int
	var tcB [][]XmlTableTcBorders
	for yi := 0; yi < y; yi++ {
		var t1 []XmlText
		var vw1 []int
		var tcB1 []XmlTableTcBorders
		for xi := 0; xi < x; xi++ {
			t1 = append(t1, GetXmlText(""))
			vw1 = append(vw1, 500)
			tcB1 = append(tcB1, XmlTableTcBorders{})
		}

		t = append(t, t1)
		vw = append(vw, vw1)
		vh = append(vh, 50)
		tcB = append(tcB, tcB1)
	}
	return XmlTable{x: x, y: y, Data: t, CWidth: vw, CHeight: vh, MergeCells: [][3]int{}, TcBorders: tcB, TcpPr: XmlTableTcpPr{}}
}

func (table *XmlTable) SetWidth(colums []int) {
	for i := range table.CWidth {
		for iw, width := range colums {
			table.CWidth[i][iw] = width
		}
	}
}
func (table *XmlTable) SetHeight(lines []int) {
	table.CHeight = lines
}

func (table *XmlTable) SetWidthColum(colum int, width int) {
	for i := range table.CWidth {
		table.CWidth[i][colum] = width
	}
}

func (table *XmlTable) AddMergeCells(line int, c1, c2 int) {
	table.MergeCells = append(table.MergeCells, [3]int{line, c1, c2})
}

type XmlTable struct {
	x, y       int                   // Размер таблицы: x - кол-во колонок; y - кол-во линий
	Data       [][]XmlText           // Содержание полей [y][x]
	CWidth     [][]int               // Ширина колонок
	CHeight    []int                 // Высота линий
	TcBorders  [][]XmlTableTcBorders // Задать стиль границам
	TcpPr      XmlTableTcpPr
	MergeCells [][3]int // Объединение полей todo (косяк при объединение нескольких, не работает в высоту)
}
type XmlTableTcpPr struct {
	TblPX int
}

type XmlTableTcBorders struct {
	Top    TcBordersParams
	Left   TcBordersParams
	Bottom TcBordersParams
	Right  TcBordersParams
}

type TcBordersParams struct {
	Size  int
	Space int
	Color ColorRGB
}
type ColorRGB struct {
	R uint8
	G uint8
	B uint8
}

func (c ColorRGB) ToHEX() string {
	return fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
}

//
//func TestRgb() {
//	var c ColorRGB
//	c.R = 255
//	c.G = 255
//	c.B = 255
//	fmt.Printf("%v", c.ToHEX())
//}
