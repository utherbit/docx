package docx_models

import (
	"docx/create_docx"
	"docx/docx_format"
	"docx/docx_format/border_line"
	"docx/docx_format/rgb"
	"docx/docx_xml"
	"fmt"
	"strconv"
)

func (d Document) Save(dir string) (err error) {
	xml := d.GenerateTag().GenerateXml()
	modFiles := map[string][]byte{"word/document.xml": []byte(xml)}
	err = create_docx.WriteToFile(dir, modFiles)

	return
}

func (d Document) GenerateTag() docx_xml.Tag {
	doc := docx_xml.GetTagDocument()
	for _, elem := range d.elements {

		tag := elem.generateTag()
		if tag != nil {
			doc.AddContentTags(tag...)
		}
	}

	return doc
}

//func (e elementModel) parseElement() (tag docx_xml.Tag, exist bool) {
//	exist = false
//
//	switch e.element.(type) {
//	case *Paragraph:
//		tag = e.element.(*Paragraph).generateTag()
//		exist = true
//	case *Table:
//		tag = e.element.(*Table).generateTag()
//		exist = true
//	}
//	return
//}

//func (e elementModel) parseElement() (tag docx_xml.Tag, exist bool) {
//	exist = false
//
//	switch e.elemType {
//	case elementTypeParagraph:
//		tag = e.elemParagraph.generateTag()
//		exist = true
//	case elementTypeTable:
//		tag = e.elemTable.generateTag()
//		exist = true
//	}
//	return
//}

func (t Table) generateTag() []docx_xml.Tag {

	// Подготовка данных с учётом merge
	for _, row := range t.rows {
		for _, cell := range row.cells {
			cell.format.Width = t.columns[cell.indexX].format.Width
			if cell.merge != nil {
				if cell.merge.general {
					continue // Общая ячейка пропускается
				}
				// Добавление контента присоединённых ячеек, в общую ячейку
				cell.merge.fixed.content = append(cell.merge.fixed.content, cell.content...)
				// Задаётся новая ширина ячеек, с учётом соединения
				if cell.indexX != cell.merge.fixed.indexX {
					row.cells[cell.merge.fixed.indexX].format.Width += cell.format.Width
				}
			}
		}
	}

	table := docx_xml.GetTag("w:tbl")

	//w:tblPr
	tblPr := docx_xml.GetTag("w:tblPr")
	tblPr.AddContentTags(tableBorders())
	//tblPr.SetContentText("")

	//borders := tableBorders()
	table.AddContentTags(tblPr)

	//w:tblGrid
	tblGrid := docx_xml.GetTag("w:tblGrid")
	for _, c := range t.columns {

		gridCol := docx_xml.GetTag("w:gridCol")
		{ // Add tag <w:gridCol w:w="Width">
			if c.format.Width != 0 { // Если ширина не указана
				gridCol.AddAttribute("w:w", strconv.Itoa(c.format.Width))
			}
			//else {
			//	gridCol.AddAttribute("w:w", strconv.Itoa(c.format.Width))
			//	gridCol.AddAttribute("w:type", "dxa")
			//}
		}
		tblGrid.AddContentTags(gridCol)
	}
	//tblGrid.SetContentText("")
	table.AddContentTags(tblGrid)

	for _, r := range t.rows {
		row := docx_xml.GetTag("w:tr")

		//w:trPr
		trPr := docx_xml.GetTag("w:trPr")
		{ // Add tag <w:trPr w:trHeight="Height">
			trHeight := docx_xml.GetTag("w:trHeight")

			trHeight.AddAttribute("w:hRule", r.format.HeightRule.String())
			trHeight.AddAttribute("w:val", strconv.Itoa(r.format.Height))

			trPr.AddContentTags(trHeight)
		}
		row.AddContentTags(trPr)

		for _, c := range r.cells {
			cell := docx_xml.GetTag("w:tc")
			//cellColumn := t.Column(c.indexX)

			// генерация формата ячейки
			{
				tcPr := docx_xml.GetTag("w:tcPr")
				// добавляется объединения ячеек таблицы
				if c.merge != nil {

					var general *TableCell // левая верхняя ячейка в объединении
					if c.merge.general {
						general = c
					} else {
						general = c.merge.fixed
					}
					// количество горизонтально объединенных ячеек - 1
					horizon := general.merge.fixed.indexX - general.indexX
					// количество вертикально объединенных ячеек - 1
					vertical := general.merge.fixed.indexY - general.indexY
					// Если есть горизонтально объединённые ячейки
					if horizon > 0 {
						// Пропустить горизонтально присоединённые ячейки
						if general.indexX != c.indexX {
							continue
						}
						// Задаётся размер горизонтальной сетки ячейки
						gridSpan := docx_xml.GetTag("w:gridSpan")
						gridSpan.AddAttribute("w:val", strconv.Itoa(horizon+1))
						tcPr.AddContentTags(gridSpan)
					}
					// Если есть вертикально объединенные ячейки
					if vertical > 0 {
						vMerge := docx_xml.GetTag("w:vMerge")
						if c.merge.general {
							vMerge.AddAttribute("w:val", "restart")
						}
						tcPr.AddContentTags(vMerge)
					}
				}
				// Задать ширину ячейки
				{
					tcW := docx_xml.GetTag("w:tcW")
					if c.format.Width == 0 { // Если ширина не указана
						tcW.AddAttribute("w:type", "auto")
					} else {
						tcW.AddAttribute("w:w", strconv.Itoa(c.format.Width))
						tcW.AddAttribute("w:type", "dxa")
					}
					tcPr.AddContentTags(tcW)
				}
				// Задать цвет выделения ячейки
				{
					if !compareRGB(c.format.Shadow, rgb.ColorRGB{R: 0, G: 0, B: 0}) {
						highlight := docx_xml.GetTag("w:shd")
						highlight.AddAttribute("w:fill", c.format.Shadow.ToHEX())
						tcPr.AddContentTags(highlight)
					}
				}
				// Задать границы ячейки
				{
					borders := getTagTcBorders(c.format.Border)
					tcPr.AddContentTags(borders)
				}
				cell.AddContentTags(tcPr)
			}
			// Генерация содержимого ячейки
			{
				if len(c.content) != 0 {
					for _, element := range c.content {
						tag := element.generateTag()
						cell.AddContentTags(tag...)
					}
				} else { // Добавить параграф если ячейка пуста
					tag := docx_xml.GetTag("w:p")
					cell.AddContentTags(tag)
				}
			}

			row.AddContentTags(cell)
		}

		table.AddContentTags(row)
	}

	return []docx_xml.Tag{table}
}
func tableBorders() docx_xml.Tag {
	borders := docx_xml.GetTag("w:tblBorders")
	{
		top := docx_xml.GetTag("w:top")
		top.AddAttribute("w:val", "single")
		top.AddAttribute("w:sz", "4")
		top.AddAttribute("w:space", "0")
		top.AddAttribute("w:color", "auto")
		borders.AddContentTags(top)
	}
	{
		left := docx_xml.GetTag("w:left")
		left.AddAttribute("w:val", "single")
		left.AddAttribute("w:sz", "4")
		left.AddAttribute("w:space", "0")
		left.AddAttribute("w:color", "auto")
		borders.AddContentTags(left)
	}
	{
		bottom := docx_xml.GetTag("w:bottom")
		bottom.AddAttribute("w:val", "single")
		bottom.AddAttribute("w:sz", "4")
		bottom.AddAttribute("w:space", "0")
		bottom.AddAttribute("w:color", "auto")
		borders.AddContentTags(bottom)
	}
	{
		right := docx_xml.GetTag("w:right")
		right.AddAttribute("w:val", "single")
		right.AddAttribute("w:sz", "4")
		right.AddAttribute("w:space", "0")
		right.AddAttribute("w:color", "auto")
		borders.AddContentTags(right)
	}
	{
		insideH := docx_xml.GetTag("w:insideH")
		insideH.AddAttribute("w:val", "single")
		insideH.AddAttribute("w:sz", "4")
		insideH.AddAttribute("w:space", "0")
		insideH.AddAttribute("w:color", "auto")
		borders.AddContentTags(insideH)
	}
	{
		insideV := docx_xml.GetTag("w:insideV")
		insideV.AddAttribute("w:val", "single")
		insideV.AddAttribute("w:sz", "4")
		insideV.AddAttribute("w:space", "0")
		insideV.AddAttribute("w:color", "auto")
		borders.AddContentTags(insideV)
	}
	return borders
}

//<w:tblBorders>
//<w:top w:val="single" w:sz="4" w:space="0" w:color="auto"/>
//<w:left w:val="single" w:sz="4" w:space="0" w:color="auto"/>
//<w:bottom w:val="single" w:sz="4" w:space="0" w:color="auto"/>
//<w:right w:val="single" w:sz="4" w:space="0" w:color="auto"/>
//<w:insideH w:val="single" w:sz="4" w:space="0" w:color="auto"/>
//<w:insideV w:val="single" w:sz="4" w:space="0" w:color="auto"/>
//</w:tblBorders>

func (p Paragraph) generateTag() []docx_xml.Tag {
	paragraph := docx_xml.GetTag("w:p")

	//rpr := GetTagRpr(data.style, data.fount, data.size)
	ppr := getTagPpr(p.format)
	paragraph.AddContentTags(ppr)

	//ppr.AddContentTags(rpr)
	for _, text := range p.texts {
		paragraph.AddContentTags(text.generateTag())
	}

	return []docx_xml.Tag{paragraph}
}
func (t TextModel) generateTag() docx_xml.Tag {
	tagR := docx_xml.GetTag("w:r")

	tagT := docx_xml.GetTag("w:t")
	if extremeSpace(t.Text) {
		tagT.AddAttribute("xml:space", "preserve")
	}
	tagT.SetContentText(t.Text)

	rpr := getTagRpr(t.format)
	tagR.AddContentTags(rpr, tagT)
	return tagR
}

func extremeSpace(text string) bool {
	lenT := len(text)
	if lenT > 0 {
		return text[0] == ' ' || text[lenT-1] == ' '
	}
	return false
}

func getTagRpr(format docx_format.TextFormat) docx_xml.Tag {
	rpr := docx_xml.GetTag("w:rPr")

	// Задать шрифт
	{
		if format.Font != docx_format.FontCalibri {
			rFonts := docx_xml.GetTag("w:rFonts")
			fontName := getFountName(format.Font)
			rFonts.AddAttribute("w:ascii", fontName)
			rFonts.AddAttribute("w:hAnsi", fontName)
			rpr.AddContentTags(rFonts)
		}
	}
	// Задать стиль выделения
	{
		if format.Style.TextUnderline {
			u := docx_xml.GetTag("w:u")
			u.AddAttribute("w:val", "single")
			rpr.AddContentTags(u)
		}
		if format.Style.TextCursive {
			i := docx_xml.GetTag("w:i")
			rpr.AddContentTags(i)
		}
		if format.Style.TextBold {
			b := docx_xml.GetTag("w:b")
			rpr.AddContentTags(b)
		}
	}

	// Задать цвет текста
	{
		if !compareRGB(format.Color, rgb.ColorRGB{R: 0, G: 0, B: 0}) {
			fmt.Printf("\nSetTextColor: %v", format.Color.ToHEX())
			textColor := docx_xml.GetTag("w:color")
			textColor.AddAttribute("w:val", format.Color.ToHEX())
			rpr.AddContentTags(textColor)
		}
	}
	// Задать цвет выделения
	{
		if !compareRGB(format.Shadow, rgb.ColorRGB{R: 255, G: 255, B: 255}) {
			highlight := docx_xml.GetTag("w:shd")
			highlight.AddAttribute("w:fill", format.Shadow.ToHEX())
			rpr.AddContentTags(highlight)
		}
	}
	// Задать размер текста
	{
		if format.Size != 22 {
			textSize := strconv.Itoa(int(format.Size))

			sz := docx_xml.GetTag("w:sz")
			sz.AddAttribute("w:val", textSize)
			rpr.AddContentTags(sz)
		}
	}
	return rpr
}

const (
	defaultSpacingAfter     = 8
	defaultSpacingBefore    = 0
	defaultSpacingRight     = 0
	defaultSpacingLeft      = 0
	defaultSpacingHanging   = 0
	defaultSpacingFirstLine = 0
)

func getTagPpr(format docx_format.ParagraphFormat) docx_xml.Tag {
	ppr := docx_xml.GetTag("w:pPr")

	if format.Numbering.ILvl != 0 {
		numPr := docx_xml.GetTag("w:numPr")
		iLvl := docx_xml.GetTag("w:ilvl")
		iLvl.AddAttribute("w:val", strconv.Itoa(format.Numbering.ILvl))
		numId := docx_xml.GetTag("w:numId")
		numId.AddAttribute("w:val", strconv.Itoa(format.Numbering.NumId))
		numPr.AddContentTags(iLvl, numId)
		ppr.AddContentTags(numPr)
	}

	if format.Spacing.Left != defaultSpacingLeft || format.Spacing.Right != defaultSpacingRight ||
		format.Spacing.FirstLine != defaultSpacingFirstLine || format.Spacing.Hanging != defaultSpacingHanging {
		ind := docx_xml.GetTag("w:ind")
		if format.Spacing.Left != defaultSpacingLeft {
			ind.AddAttribute("w:left", strconv.Itoa(format.Spacing.Left))
		}
		if format.Spacing.Right != defaultSpacingRight {
			ind.AddAttribute("w:right", strconv.Itoa(format.Spacing.Right))
		}
		if format.Spacing.FirstLine != defaultSpacingFirstLine {
			ind.AddAttribute("w:firstLine", strconv.Itoa(format.Spacing.FirstLine))
		}
		if format.Spacing.Hanging != defaultSpacingHanging {
			ind.AddAttribute("w:hanging", strconv.Itoa(format.Spacing.Hanging))
		}
		ppr.AddContentTags(ind)
	}

	if format.Spacing.Before != defaultSpacingBefore || format.Spacing.After != defaultSpacingAfter {
		spa := docx_xml.GetTag("w:spacing")
		if format.Spacing.Before != defaultSpacingBefore {
			spa.AddAttribute("w:before", strconv.Itoa(format.Spacing.Before))
		}
		if format.Spacing.After != defaultSpacingAfter {
			spa.AddAttribute("w:after", strconv.Itoa(format.Spacing.After))
		}

		ppr.AddContentTags(spa)
	}

	if format.Align != docx_format.AlignLeft {
		jc := docx_xml.GetTag("w:jc")
		jc.AddAttribute("w:val", getAlignmentString(format.Align))
		ppr.AddContentTags(jc)
	}

	return ppr
}
func getAlignmentString(alignment docx_format.Alignment) string {
	switch alignment {
	default:
		return "left"
	case docx_format.AlignRight:
		return "right"
	case docx_format.AlignCenter:
		return "center"
	case docx_format.AlignBoth:
		return "both"
	}
}
func compareRGB(color1, color2 rgb.ColorRGB) bool {
	return color1.R == color2.R && color1.G == color2.G && color1.B == color2.B
}

func getTagTcBorders(borders border_line.CellBorders) docx_xml.Tag {
	tcBorders := docx_xml.GetTag("w:tcBorders")
	generateTagBorderFormat := func(border *docx_xml.Tag, format border_line.BorderFormat) {
		if compareRGB(format.Color, rgb.ColorRGB{R: 255, G: 255, B: 255}) {
			border.AddAttribute("w:color", format.Color.ToHEX())
		}
		if format.Shadow {
			border.AddAttribute("w:shadow", "true")
		}
		if format.Space != 0 {
			border.AddAttribute("w:space", strconv.Itoa(format.Space))
		}
		if format.Size != 4 {
			border.AddAttribute("w:sz", strconv.Itoa(format.Size))
		}
		border.AddAttribute("w:val", string(format.Value))

	}
	//  Top    BorderFormat
	if !compareBorderFormatIsDefault(borders.Top) {
		border := docx_xml.GetTag("w:top")
		generateTagBorderFormat(&border, borders.Top)
		tcBorders.AddContentTags(border)
	}
	//	Bottom BorderFormat
	if !compareBorderFormatIsDefault(borders.Bottom) {
		border := docx_xml.GetTag("w:bottom")
		generateTagBorderFormat(&border, borders.Bottom)
		tcBorders.AddContentTags(border)

	}
	//	Start  BorderFormat
	if !compareBorderFormatIsDefault(borders.Start) {
		border := docx_xml.GetTag("w:start")
		generateTagBorderFormat(&border, borders.Start)
		tcBorders.AddContentTags(border)

	}
	//	End    BorderFormat
	if !compareBorderFormatIsDefault(borders.End) {
		border := docx_xml.GetTag("w:end")
		generateTagBorderFormat(&border, borders.End)
		tcBorders.AddContentTags(border)

	}

	//	InsideH BorderFormat
	if borders.InsideH.Value != border_line.BorderLineNil {
		border := docx_xml.GetTag("w:insideH")
		generateTagBorderFormat(&border, borders.InsideH)
		tcBorders.AddContentTags(border)

	}
	//	InsideV BorderFormat
	if borders.InsideV.Value != border_line.BorderLineNil {
		border := docx_xml.GetTag("w:insideV")
		generateTagBorderFormat(&border, borders.InsideV)
		tcBorders.AddContentTags(border)

	}
	//	Tl2br   BorderFormat
	if borders.Tl2br.Value != border_line.BorderLineNil {
		border := docx_xml.GetTag("w:tl2br")
		generateTagBorderFormat(&border, borders.Tl2br)
		tcBorders.AddContentTags(border)

	}
	//	Tr2bl   BorderFormat
	if borders.Tr2bl.Value != border_line.BorderLineNil {
		border := docx_xml.GetTag("w:tr2bl")
		generateTagBorderFormat(&border, borders.Tr2bl)
		tcBorders.AddContentTags(border)

	}
	return tcBorders
}

func compareBorderFormatIsDefault(border border_line.BorderFormat) bool {
	def := border_line.BorderFormat{
		Color:  rgb.ColorRGB{R: 255, G: 255, B: 255},
		Shadow: false,
		Space:  0,
		Size:   4,
		Value:  border_line.BorderLineSingle,
	}
	return border == def
}
