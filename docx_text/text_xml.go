package docx_text

import (
	"docx/docx_xml"
	"strconv"
)

func GetTagText(data DocxText) docx_xml.Tag {
	p := docx_xml.GetTag("w:p")

	rpr := GetTagRpr(data.style, data.fount, data.size)
	ppr := GetTagPpr(data.spacing, data.align)

	ppr.AddContentTags(rpr)

	r := docx_xml.GetTag("w:r")
	t := docx_xml.GetTag("w:t")
	t.SetContentText(data.Text)

	r.AddContentTags(rpr, t)

	p.AddContentTags(ppr, r)

	return p
}

func GetTagPpr(spacing TextSpacing, alignment Alignment) docx_xml.Tag {
	ppr := docx_xml.GetTag("w:pPr")
	if !spacing.isDefaultLeft() || !spacing.isDefaultRight() {
		ind := docx_xml.GetTag("w:ind")
		if !spacing.isDefaultLeft() {
			ind.AddAttribute("w:left", strconv.Itoa(spacing.Left))
		}
		if !spacing.isDefaultRight() {
			ind.AddAttribute("w:right", strconv.Itoa(spacing.Right))
		}
		ppr.AddContentTags(ind)
	}

	if !spacing.isDefaultBefore() || !spacing.isDefaultAfter() {
		spa := docx_xml.GetTag("w:spacing")
		if !spacing.isDefaultBefore() {
			spa.AddAttribute("w:before", strconv.Itoa(spacing.Before))
		}
		if !spacing.isDefaultAfter() {
			spa.AddAttribute("w:after", strconv.Itoa(spacing.After))
		}
		ppr.AddContentTags(spa)
	}

	if alignment != AlignLeft {
		jc := getTagAlignment(alignment)
		ppr.AddContentTags(jc)
	}

	return ppr
}

func GetTagRpr(style TextStyle, fonts Fonts, size int) docx_xml.Tag {
	rpr := docx_xml.GetTag("w:rPr")

	if !fonts.isDefault() {
		rFonts := getTagFont(fonts)
		rpr.AddContentTags(rFonts)
	}

	if style.TextUnderline {
		u := docx_xml.GetTag("w:u")
		u.AddAttribute("w:val", "single")
		rpr.AddContentTags(u)
	}
	if style.TextCursive {
		i := docx_xml.GetTag("w:i")
		rpr.AddContentTags(i)
	}
	if style.TextBold {
		b := docx_xml.GetTag("w:b")
		rpr.AddContentTags(b)
	}
	if size != 22 {
		textSize := strconv.Itoa(size)

		sz := docx_xml.GetTag("w:sz")
		sz.AddAttribute("w:val", textSize)
		rpr.AddContentTags(sz)

		szCs := docx_xml.GetTag("w:szCs")
		szCs.AddAttribute("w:val", textSize)
		rpr.AddContentTags(sz)
	}
	return rpr
}
