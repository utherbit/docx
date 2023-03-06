package docx_xml

func GetTagSectPr() Tag {
	sectPr := GetTag("w:sectPr")
	pgSz := GetTag("w:pgSz")

	pgSz.AddAttribute("w:w", "12240")
	pgSz.AddAttribute("w:h", "15840")

	sectPr.AddContentTags(pgSz)
	return sectPr
}

func GetTagDocument() Tag {
	document := GetTag("w:document")

	document.AddAttribute("xmlns:wpc", "http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas")
	document.AddAttribute("xmlns:cx", "http://schemas.microsoft.com/office/drawing/2014/chartex")
	document.AddAttribute("xmlns:cx1", "http://schemas.microsoft.com/office/drawing/2015/9/8/chartex")
	document.AddAttribute("xmlns:cx2", "http://schemas.microsoft.com/office/drawing/2015/10/21/chartex")
	document.AddAttribute("xmlns:cx3", "http://schemas.microsoft.com/office/drawing/2016/5/9/chartex")
	document.AddAttribute("xmlns:cx4", "http://schemas.microsoft.com/office/drawing/2016/5/10/chartex")
	document.AddAttribute("xmlns:cx5", "http://schemas.microsoft.com/office/drawing/2016/5/11/chartex")
	document.AddAttribute("xmlns:cx6", "http://schemas.microsoft.com/office/drawing/2016/5/12/chartex")
	document.AddAttribute("xmlns:cx7", "http://schemas.microsoft.com/office/drawing/2016/5/13/chartex")
	document.AddAttribute("xmlns:cx8", "http://schemas.microsoft.com/office/drawing/2016/5/14/chartex")
	document.AddAttribute("xmlns:mc", "http://schemas.openxmlformats.org/markup-compatibility/2006")
	document.AddAttribute("xmlns:aink", "http://schemas.microsoft.com/office/drawing/2016/ink")
	document.AddAttribute("xmlns:am3d", "http://schemas.microsoft.com/office/drawing/2017/model3d")
	document.AddAttribute("xmlns:o", "urn:schemas-microsoft-com:office:office")
	document.AddAttribute("xmlns:r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships")
	document.AddAttribute("xmlns:m", "http://schemas.openxmlformats.org/officeDocument/2006/math")
	document.AddAttribute("xmlns:v", "urn:schemas-microsoft-com:vml")
	document.AddAttribute("xmlns:wp14", "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing")
	document.AddAttribute("xmlns:wp", "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing")
	document.AddAttribute("xmlns:w10", "urn:schemas-microsoft-com:office:word")
	document.AddAttribute("xmlns:w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	document.AddAttribute("xmlns:w14", "http://schemas.microsoft.com/office/word/2010/wordml")
	document.AddAttribute("xmlns:w15", "http://schemas.microsoft.com/office/word/2012/wordml")
	document.AddAttribute("xmlns:w16cid", "http://schemas.microsoft.com/office/word/2016/wordml/cid")
	document.AddAttribute("xmlns:w16se", "http://schemas.microsoft.com/office/word/2015/wordml/symex")
	document.AddAttribute("xmlns:wpg", "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup")
	document.AddAttribute("xmlns:wpi", "http://schemas.microsoft.com/office/word/2010/wordprocessingInk")
	document.AddAttribute("xmlns:wne", "http://schemas.microsoft.com/office/word/2006/wordml")
	document.AddAttribute("xmlns:wps", "http://schemas.microsoft.com/office/word/2010/wordprocessingShape")
	document.AddAttribute("mc:Ignorable", "w14 w15 w16se w16cid wp14")

	return document
}
