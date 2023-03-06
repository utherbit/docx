package main

import (
	df "docx/docx_format"
	"docx/docx_format/border_line"
	"docx/docx_format/rgb"
	docx "docx/docx_models"
)

func SampleOne() {
	doc := docx.Document{}

	headerFP := df.NewParagraphFormat().SetAlign(df.AlignCenter)
	headerFT := df.NewTextFormat().SetSize(40).SetStyle(df.TextStyleBold).SetFont(df.FontTimesNewRoman)

	simpleTextFP := df.NewParagraphFormat().SetAlign(df.AlignLeft).SetSpacing(df.Spacing{FirstLine: 400, After: 200})
	simpleTextFT := df.NewTextFormat().SetSize(32).SetFont(df.FontTimesNewRoman)

	p := doc.AddParagraph()
	p.SetText(headerFP, headerFT, "Это пример сгенерированного документа")

	p = doc.AddParagraph()
	p.SetText(simpleTextFP, simpleTextFT,
		"Я постарался реализовать самые необходимые, по моему мнению функции word документа")

	p = doc.AddParagraph().Format(simpleTextFP)
	p.AddText("Большая часть документа состоит из параграфов. ").Format(simpleTextFT)
	p.AddText("У параграфа есть два параметра: отступы и выравнивание. ").Format(simpleTextFT)
	p.AddText("Отступы задаются числовым значением, можно задать отступы слева, справа, сверху и снизу. ").Format(simpleTextFT)
	p.AddText("Выравнивание может иметь 4 значения - по левому краю, по правому, по центру или по ширине страницы. ").Format(simpleTextFT)

	p = doc.AddParagraph().Format(simpleTextFP)
	p.AddText("Параграфы могут содержать 1 или больше тегов с текстом. ").Format(simpleTextFT)
	p.AddText("Тексту можно задать множество параметров, например: ").Format(simpleTextFT)
	p.AddText("размер текста, ").Format(simpleTextFT.SetSize(38))
	p.AddText("шрифт, ").Format(simpleTextFT.SetFont(df.FontCalibri))
	p.AddText("цвет текста, ").Format(simpleTextFT.SetColor(rgb.ColorRGB{R: 255, G: 0, B: 0}))
	p.AddText("цвет выделения, ").Format(simpleTextFT.SetShadow(rgb.ColorRGB{R: 0, G: 255, B: 0}))
	p.AddText("стили выделения - ").Format(simpleTextFT)
	p.AddText("жирный текст, ").Format(simpleTextFT.SetStyle(df.TextStyleBold))
	p.AddText("курсив, ").Format(simpleTextFT.SetStyle(df.TextStyleCursive))
	p.AddText("подчёркнутый. ").Format(simpleTextFT.SetStyle(df.TextStyleUnderline))
	p.AddText("Конечно их можно ").Format(simpleTextFT)
	p.AddText("комбинировать. ").Format(simpleTextFT.SetStyle(df.TextStyleCursive, df.TextStyleUnderline, df.TextStyleBold))

	p = doc.AddParagraph().Format(simpleTextFP.SetSpacing(df.Spacing{FirstLine: 400, Before: 600, After: 100}))
	p.AddText("Вам может быть полезно создание таблиц. ").Format(simpleTextFT)
	p.AddText("Я постарался реализовать этот функционал максимально удобным для использования. ").Format(simpleTextFT)

	t := doc.AddTable(3, 3)

	for _, column := range t.Columns() {
		column.Format(df.NewTableColumnFormat().SetWidth(3000))
	}

	tblHeaderFP := df.NewParagraphFormat().SetAlign(df.AlignCenter)
	t.Cell(0, 0).SetText(tblHeaderFP, simpleTextFT, "Колонка 1").Format(df.NewTableCellFormat().SetShadow(rgb.ColorRGB{R: 0, G: 0, B: 255}))
	t.Cell(0, 1).SetText(tblHeaderFP, simpleTextFT, "Колонка 2")
	t.Cell(0, 2).SetText(tblHeaderFP, simpleTextFT, "Колонка 3")

	p = doc.AddParagraph().Format(simpleTextFP.SetSpacing(df.Spacing{FirstLine: 400, Before: 200, After: 200}))
	p.AddText("Можно очень просто объединять ячейки таблицы, ").Format(simpleTextFT)
	p.AddText("любым, допустимым в редакторе способом.").Format(simpleTextFT)

	t = doc.AddTable(3, 5)
	columnF := df.NewTableColumnFormat()
	t.Column(0).Format(columnF.SetWidth(1638))
	t.Column(1).Format(columnF.SetWidth(1242))
	t.Column(2).Format(columnF.SetWidth(2880))
	t.Column(3).Format(columnF.SetWidth(2178))
	t.Column(4).Format(columnF.SetWidth(702))

	testCellFormat := df.NewTableCellFormat().SetBorders(
		border_line.CellBorders{}.SetExternalBorders(
			border_line.NewBorderFormat().SetBorderLine(border_line.BorderLineThinThickThinLargeGap).SetColor(rgb.ColorRGB{R: 255})))
	exampleMergeFP := df.NewParagraphFormat().SetAlign(df.AlignLeft)
	{
		t.Cell(0, 0).SetText(exampleMergeFP, simpleTextFT, "AAA").Format(testCellFormat)
		t.Merge(t.Cell(0, 0), t.Cell(0, 1))
		t.Cell(0, 2).SetText(exampleMergeFP, simpleTextFT, "BBB")
		t.Merge(t.Cell(0, 2), t.Cell(0, 3))
		t.Cell(0, 4).SetText(exampleMergeFP, simpleTextFT, "CCC")
	}
	{
		t.Cell(1, 0).SetText(exampleMergeFP, simpleTextFT, "DDD")
		t.Cell(1, 1).SetText(exampleMergeFP, simpleTextFT, "EEE")
		t.Merge(t.Cell(1, 1), t.Cell(1, 2))
		t.Cell(1, 3).SetText(exampleMergeFP, simpleTextFT, "FFF")
		t.Merge(t.Cell(1, 3), t.Cell(2, 4))
	}
	{
		t.Cell(2, 0).SetText(exampleMergeFP, simpleTextFT, "GGG")
		t.Merge(t.Cell(2, 0), t.Cell(2, 1))
		t.Cell(2, 2).SetText(exampleMergeFP, simpleTextFT, "HHH")
		t.Cell(2, 3).SetText(exampleMergeFP, simpleTextFT, "III")
	}

	err := doc.Save("output/sample-one.docx")
	panicIfErr(err)
}
