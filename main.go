package main

import (
	"docx/create_docx"
	"docx/docx_format"
	dm "docx/docx_models"
	"docx/docx_text"
	"docx/docx_xml"
	document "docx/xml_text_edit_gen_doc"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var outputDir = "./output"

var (
	// Стили для параграфов
	pf0 = docx_format.NewParagraphFormat().SetSpacing(docx_format.Spacing{FirstLine: 260, After: 200, Before: 200})
	pf1 = docx_format.NewParagraphFormat().SetAlign(docx_format.AlignBoth).SetSpacing(docx_format.Spacing{FirstLine: 260, After: 200, Before: 200})
	pf2 = docx_format.NewParagraphFormat().SetAlign(docx_format.AlignCenter).SetSpacing(docx_format.Spacing{Before: 24})
	pf3 = docx_format.NewParagraphFormat().SetAlign(docx_format.AlignLeft).SetNumbering(docx_format.Numbering{ILvl: 1, NumId: 1})
	// Стили для текстов
	rf0 = docx_format.NewTextFormat()
	rf1 = docx_format.NewTextFormat().SetFont(docx_format.FontTimesNewRoman).SetSize(20)
	rf2 = docx_format.NewTextFormat().SetFont(docx_format.FontTimesNewRoman).SetSize(24).SetStyle(docx_format.TextStyleUnderline)
	rf3 = docx_format.NewTextFormat().SetFont(docx_format.FontTimesNewRoman).SetSize(28).SetStyle(docx_format.TextStyleBold)
)

func newStyle() {
	format := docx_format.NewParagraphFormat()
	format.SetAlign(docx_format.AlignLeft)
}

func main() {
	SampleOne()
}

func main2() {

	var doc dm.Document
	p := doc.AddParagraph().Format(pf2)
	p.AddText("Заголовок документа").Format(rf3)

	doc.AddParagraph().SetText(pf3, rf1, "Список1")
	doc.AddParagraph().SetText(pf3, rf1, "Список2")
	doc.AddParagraph().SetText(pf3, rf1, "Список3")
	doc.AddParagraph().SetText(pf3, rf1, "Список4")

	text1 := "Повседневная практика показывает, что постоянное информационно-пропагандистское обеспечение нашей деятельности позволяет оценить значение модели развития. Таким образом реализация намеченных плановых заданий позволяет выполнять важные задания по разработке направлений прогрессивного развития. Разнообразный и богатый опыт реализация намеченных плановых заданий требуют определения и уточнения направлений прогрессивного развития."
	text2 := "Значимость этих проблем настолько очевидна, что рамки и место обучения кадров играет важную роль в формировании новых предложений. "
	text3 := " Новая модель организационной деятельности представляет собой интересный эксперимент проверки системы обучения кадров, соответствует насущным потребностям. Равным образом укрепление и развитие структуры позволяет оценить значение новых предложений.  рамки и место обучения кадров требуют определения и уточнения модели развития. Разнообразный и богатый опыт новая модель организационной деятельности позволяет оценить значение дальнейших направлений развития. "

	p = doc.AddParagraph().Format(pf1)
	p.AddText(text1).Format(rf0)

	p = doc.AddParagraph().Format(pf1)
	p.AddText(text2).Format(rf0)
	p.AddText(" Товарищи!").Format(rf0.SetStyle(docx_format.TextStyleCursive, docx_format.TextStyleBold))
	p.AddText(text3).Format(rf0)

	p = doc.AddParagraph()
	p.AddText("Для этого текста будут заданы разные стили. ")
	p.AddText("Жирный текст. ").Format(rf0.SetStyle(docx_format.TextStyleBold))
	p.AddText("Текст курсивом. ").Format(rf0.SetStyle(docx_format.TextStyleCursive))
	p.AddText("Подчёркнутый текст. ").Format(rf0.SetStyle(docx_format.TextStyleUnderline))

	newFormat := docx_format.NewTextFormat().SetStyle(docx_format.TextStyleBold, docx_format.TextStyleCursive, docx_format.TextStyleUnderline)
	p.AddText("Всё вместе!").Format(newFormat)

	p = doc.AddParagraph()
	p.AddText("Таблица 1").Format(rf1)

	table := doc.AddTable(1, 3)
	for _, column := range table.Columns() {
		column.Format(docx_format.TableColumnFormat{Width: 2000})
	}

	row := table.Row(0).Format(docx_format.TableRowFormat{Height: 100, HeightRule: docx_format.HeightRuleAtLeast})

	row.Cell(0).SetText(pf0, rf2, "Колонка 1")
	row.Cell(1).SetText(pf0, rf2, "Колонка 2")
	row.Cell(2).SetText(pf0, rf2, "Колонка 3")

	//doc.

	data := [][3]string{{"Игорь", "Иван", "Петя"}, {"Лев", "Виталий", "Андрей"}}
	for _, names := range data {
		row = table.AddRow()
		row.Format(docx_format.TableRowFormat{Height: 500, HeightRule: docx_format.HeightRuleExactly})
		row.Cell(0).AddParagraph().Format(pf0).AddText(names[0]).Format(rf0)
		row.Cell(1).AddParagraph().Format(pf0).AddText(names[1]).Format(rf0)
		row.Cell(2).AddParagraph().Format(pf0).AddText(names[2]).Format(rf0)

	}

	err := table.Merge(table.Cell(0, 0), table.Cell(0, 1))
	PanicIfErr(err)

	//for _, tableRow := range table.Rows() {
	//	println()
	//	tableRow.Format(docx_format.TableRowFormat{Height: 1000})
	//	for _, cell := range tableRow.Cells() {
	//		print("\t|")
	//		mEx, mG := cell.GetMerge()
	//		if mEx && mG {
	//			print(" G ")
	//		} else if mEx && !mG {
	//			print(" m ")
	//		} else {
	//			print(" - ")
	//		}
	//
	//	}
	//}

	err = doc.Save("./output/newDocx33.docx")
	panicIfErr(err)

	//d := doc.GenerateTag().GenerateXml()
	//fmt.Printf("\n%v", d)

	//xml := doc.GenerateTag().GenerateXml()
	//modFiles := map[string][]byte{"word/document.xml": []byte(xml)}
	//outName := outputDir + "/outputTest.docx"

	//err := create_docx.WriteToFile(outName, modFiles)
	//panicIfErr(err)
	fmt.Printf("\nДокумент создан успешно")
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main1() {
	t := time.Now()

	doc := docx_xml.GetTagDocument()
	body := docx_xml.GetTag("w:body")
	name := docx_text.GetTagText(docx_text.GetDocxText("Имя"))
	body.AddContentTags(name)
	pr := docx_xml.GetTagSectPr()
	body.AddContentTags(pr)
	doc.AddContentTags(body)
	xml := docx_xml.Prolog + doc.GenerateXml()

	modFiles := map[string][]byte{"word/document.xml": []byte(xml)}
	outName := outputDir + "/outputTest.docx"

	err := create_docx.WriteToFile(outName, modFiles)
	panicIfErr(err)

	fmt.Printf("\nДокумент %v создан успешно", outName)
	fmt.Printf("\nВремя выполнения: %v", time.Since(t))
}

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Это говно надо переписать
func writeToFileEditNameIfExist(doc *document.Docx, filename string) string {
	dir := outputDir + "/" + filename
	file, err := os.Stat(dir)
	if err != nil {
		err = doc.WriteToFile(dir)
		fmt.Printf("\nWriteToFile %v", dir)
		return filename
	} else {
		//fmt.Printf("\n%v is exist, skiped", dir)
		for i := 1; true; i++ {

			newFileName := strings.Replace(filename, ".docx", "("+strconv.Itoa(i)+").docx", 1)
			newDir := strings.Replace(dir, file.Name(), newFileName, 1)
			_, e := os.Stat(newDir)
			if e != nil {
				err = doc.WriteToFile(newDir)
				fmt.Printf("\nWriteToFile %v", newDir)
				return newFileName
			} else {
				//fmt.Printf("\n%v is exist, skiped", newDir)
			}
		}
	}
	return ""
}

// TraceErr

//type TstDoc struct {
//	err DocumentError
//}
//
//func main() {
//	doc := TstDoc{}
//
//	doc.dge()
//
//	panicIfErr(doc.err)
//}
//func (d *TstDoc) dge() {
//	teErr(d)
//}
//
//func teErr(doc *TstDoc) {
//	err := DocumentError{err: "Test error DocumentError", stack: debug.Stack()}
//	//stack := debug.Stack() // debug.Stack()
//
//	doc.err = err
//}
//
//type DocumentError struct {
//	err   string
//	stack []byte
//}
//
//func (d DocumentError) Error() string {
//	return d.err + string(d.stack)
//}
