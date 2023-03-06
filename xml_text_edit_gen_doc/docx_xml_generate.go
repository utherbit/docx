package docx

import (
	"fmt"
	"strconv"
)

func (xml XmlText) XmlTextGenerate() string {
	rpr := rprGenerate(xml)
	ppr := pprGenerate(xml, rpr)
	return `<w:p>` + ppr + `<w:r>` + rpr + `<w:t>` + xml.Text + `</w:t></w:r>` + `</w:p>`
}
func (array XmlTextA) XmlTextGenerate() string {
	xml := ""
	for i, text := range array {
		rpr := rprGenerate(text)
		xml += `<w:r>` + rpr + `<w:t>` + text.Text + `</w:t></w:r>`
		if i != len(array)-1 {
			xml += `<w:r>` + rpr + `<w:t xml:space="preserve"> </w:t></w:r>`
		}
	}
	ppr := pprGenerate(array[0], rprGenerate(array[0]))
	return "<w:p>" + ppr + xml + "</w:p>"
}

//<w:spacing w:before="20" w:after="20"/>
//<w:ind w:left="567" w:right="567"/>
func (spacing TextSpacing) getSpacing() string {
	xml := ``
	if spacing.Left != 0 || spacing.Right != 0 {
		xml += `<w:ind`
		if spacing.Left != 0 {
			xml += ` w:left="` + strconv.Itoa(spacing.Left) + `"`
		}
		if spacing.Right != 0 {
			xml += ` w:right="` + strconv.Itoa(spacing.Right) + `"`
		}
		xml += `/>`
	}
	if spacing.Before != 0 || spacing.After != 8 {
		xml += `<w:spacing`
		if spacing.Before != 0 {
			xml += ` w:before="` + strconv.Itoa(spacing.Before) + `"`
		}
		if spacing.After != 8 {
			xml += ` w:after="` + strconv.Itoa(spacing.After) + `"`
		}
		xml += `/>`
	}
	return xml
}
func (style TextStyle) getStyle() string {
	xml := ""
	if style.TextUnderline {
		xml += `<w:u w:val="single"/>`
	}
	if style.TextCursive {
		xml += `<w:i/>`
	}
	if style.TextBold {
		xml += `<w:b/>`
	}

	return xml
}

func (align Alignment) getAlign() string {
	switch align {
	case AlignLeft:
		return ""
	case AlignRight:
		return `<w:jc w:val="right"/>`
	case AlignCenter:
		return `<w:jc w:val="center"/>`
	case AlignBoth:
		return `<w:jc w:val="both"/>`
	}
	return ""
}
func (font Fonts) getFount() string {
	switch font {
	case FontCalibri:
		return ""
	case FontTimesNewRoman:
		return `<w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman"/>`
	case FontWingdings2:
		return `<w:rFonts w:ascii="Wingdings 2" w:hAnsi="Wingdings 2"/>`
	}
	return ""
}
func getSize(size int) string {
	sizeStr := strconv.Itoa(size)
	return `<w:sz w:val="` + sizeStr + `"/><w:szCs w:val="` + sizeStr + `"/>`
}

func rprGenerate(text XmlText) string {
	return `<w:rPr>` + text.fount.getFount() + text.style.getStyle() + getSize(text.size) + `</w:rPr>`
}
func pprGenerate(xml XmlText, rpr string) string {
	r := `<w:pPr>` + xml.spacing.getSpacing() + xml.align.getAlign() + rpr + `</w:pPr>`
	return r
}
func textGenerate(text XmlText) string {
	rpr := rprGenerate(text)
	ppr := pprGenerate(text, rpr)
	return ppr + `<w:r>` + rpr + `<w:t>` + text.Text + `</w:t></w:r>`
}

func GetXmlTextA(text []XmlText) XmlTextA {
	return append(XmlTextA{}, text...)
}

//func TextGetCheckBox() {
//	fmt.Printf("%v", getCheckBoxTable([]checkBox{{text: "Соответствует", value: true}, {text: "Не соответствует", value: false}}).XmlTableGenerate())
//}

func (table XmlTable) XmlTableGenerate() string {
	tblpPr := ``
	if table.TcpPr.TblPX != 0 {
		tblpPr = `<w:tblpPr w:vertAnchor="text" w:horzAnchor="page" w:tblpX="` + strconv.Itoa(table.TcpPr.TblPX) + `"/>`
	}
	tblStart := `<w:tbl><w:tblPr>` + tblpPr +
		`<w:tblW w:w="0" w:type="auto"/><w:tblBorders><w:top w:val="single" w:sz="4" w:space="0" w:color="auto"/><w:left w:val="single" w:sz="4" w:space="0" w:color="auto"/><w:bottom w:val="single" w:sz="4" w:space="0" w:color="auto"/><w:right w:val="single" w:sz="4" w:space="0" w:color="auto"/><w:insideH w:val="single" w:sz="4" w:space="0" w:color="auto"/><w:insideV w:val="single" w:sz="4" w:space="0" w:color="auto"/></w:tblBorders><w:tblLook w:val="04A0" w:firstRow="1" w:lastRow="0" w:firstColumn="1" w:lastColumn="0" w:noHBand="0" w:noVBand="1"/></w:tblPr>`
	tblEnd := `</w:tbl>`

	setGrid := func() string {
		gridStart := `<w:tblGrid>`
		gridEnd := `</w:tblGrid>`
		s := ""
		for _, cw := range table.CWidth[0] {
			s += fmt.Sprintf(`<w:gridCol w:w="%v"/>`, cw)
		}
		return gridStart + s + gridEnd
	}

	setLine := func(s string) string {
		lineStart := `<w:tr w:rsidR="000463C7" w14:paraId="14D34361" w14:textId="77777777" w:rsidTr="000463C7">`
		lineEnd := `</w:tr>`
		return lineStart + s + lineEnd
	}

	setLinePr := func(h int) string {
		if h != 0 {
			return `<w:trPr><w:trHeight w:val="` + strconv.Itoa(h) + `"/></w:trPr>`
		}
		return ``
	}

	setColumn := func(s string, w int, sp bool, br XmlTableTcBorders) string {
		width := fmt.Sprintf("%v", w)
		gr := ""
		if sp {
			gr = `<w:gridSpan w:val="2"/>`
		}

		setBorders := func(br XmlTableTcBorders) string {
			def := XmlTableTcBorders{} // По умолчанию
			//fmt.Printf("\n%v", br)
			if br == def {
				return `<w:tcBorders>` +
					`<w:top w:val="single" w:sz="4" w:space="0" w:color="auto"/>` +
					`<w:left w:val="single" w:sz="4" w:space="0" w:color="auto"/>` +
					`<w:bottom w:val="single" w:sz="4" w:space="0" w:color="auto"/>` +
					`<w:right w:val="single" w:sz="4" w:space="0" w:color="auto"/>` +
					`</w:tcBorders>`
			} else {
				fmt.Printf("\nxmlRes:%v", br)
				xmlRes := `<w:tcBorders>`
				if br.Top != def.Top {
					xmlRes += fmt.Sprintf(`<w:top w:val="single" w:sz="%v" w:space="%v" w:color="%v"/>`, br.Top.Size, br.Top.Space, br.Top.Color.ToHEX())
				}
				if br.Left != def.Left {
					xmlRes += fmt.Sprintf(`<w:left w:val="single" w:sz="%v" w:space="%v" w:color="%v"/>`, br.Left.Size, br.Left.Space, br.Left.Color.ToHEX())

				}
				if br.Bottom != def.Bottom {
					xmlRes += fmt.Sprintf(`<w:bottom w:val="single" w:sz="%v" w:space="%v" w:color="%v"/>`, br.Bottom.Size, br.Bottom.Space, br.Bottom.Color.ToHEX())

				}
				if br.Right != def.Right {
					xmlRes += fmt.Sprintf(`<w:right w:val="single" w:sz="%v" w:space="%v" w:color="%v"/>`, br.Right.Size, br.Right.Space, br.Right.Color.ToHEX())
				}
				xmlRes += `</w:tcBorders>`
				//fmt.Printf("\nxmlRes:%v", xmlRes)
				return xmlRes
			}

		}

		columnStart := `<w:tc><w:tcPr><w:tcW w:w="` + width + `" w:type="dxa"/>` + gr + setBorders(br) + `</w:tcPr>`
		columnEnd := `</w:tc>`
		return columnStart + s + columnEnd
	}

	xml := tblStart

	xml += setGrid()

	if len(table.MergeCells) != 0 {
		for _, v := range table.MergeCells {
			width := table.CWidth[v[0]][v[1]] + table.CWidth[v[0]][v[2]]
			newCWidth := append(table.CWidth[v[0]][:v[1]], width)
			newCWidth = append(newCWidth, table.CWidth[v[0]][v[2]+1:]...)
			table.CWidth[v[0]] = newCWidth

			newDate := append(table.Data[v[0]][:v[1]], table.Data[v[0]][v[1]])
			newDate = append(newDate, table.Data[v[0]][v[2]+1:]...)
			table.Data[v[0]] = newDate
		}
	}

	for yi, yd := range table.Data {

		lineXml := ""
		lineXml += setLinePr(table.CHeight[yi])

		mergeLen := 0
		for xi, xd := range yd {
			ifSpan := false
			for _, v := range table.MergeCells {
				if yi == v[0] && xi == v[1]-mergeLen {
					mergeLen += 1
					ifSpan = true
				}
			}
			//fmt.Printf("\n%v, %v", xi, yi)
			lineXml += setColumn(xd.XmlTextGenerate(), table.CWidth[yi][xi], ifSpan, table.TcBorders[yi][xi])
		}
		xml += setLine(lineXml)

	}

	xml += tblEnd

	return xml
}

func GetCheckBox(b bool) string {
	getV1ifTrueElseV2 := func(b bool, v1 string, v2 string) string {
		if b {
			return v1
		} else {
			return v2
		}
	}
	return `<w:p w:rsidR="00307666" w:rsidRPr="003326AF" w:rsidRDefault="00307666" w:rsidP="00B1791E"><w:pPr><w:spacing w:after="0"/><w:rPr><w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:eastAsia="ru-RU"/></w:rPr></w:pPr><w:sdt><w:sdtPr><w:rPr><w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman"/><w:sz w:val="36"/><w:szCs w:val="24"/><w:lang w:eastAsia="ru-RU"/></w:rPr>` +
		`<w:id w:val="-1380620035"/>` +
		`<w15:color w:val="3366FF"/><w14:checkbox><w14:checked w14:val="` +
		getV1ifTrueElseV2(b, "1", "0") +
		//func(b bool) string {
		//	if b {
		//		return `1`
		//	} else {
		//		return `0`
		//	}
		//}(b) +
		`"/><w14:checkedState w14:val="0052" w14:font="Wingdings 2"/><w14:uncheckedState w14:val="00A3" w14:font="Wingdings 2"/></w14:checkbox></w:sdtPr><w:sdtContent><w:r><w:rPr><w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman"/><w:sz w:val="36"/><w:szCs w:val="24"/><w:lang w:eastAsia="ru-RU"/></w:rPr>` +
		getV1ifTrueElseV2(b, `<w:sym w:font="Wingdings 2" w:char="F052"/>`, `<w:sym w:font="Wingdings 2" w:char="F0A3"/>`) +
		//`<w:sym w:font="Wingdings 2" w:char="F052"/>` +
		`</w:r></w:sdtContent></w:sdt></w:p>`

}
