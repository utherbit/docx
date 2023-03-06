package default_files

import (
	"docx/docx_xml"
)

// GetXmlContentTypes get content filename: [Content_Types].xml
func GetXmlContentTypes() docx_xml.Tag {

	tTypes := docx_xml.GetTag("Types")
	tTypes.AddAttribute("xmlns", "http://schemas.openxmlformats.org/package/2006/content-types")

	tDefault := docx_xml.GetTag("Default")
	tDefault.AddAttribute("Extension", "rels")
	tDefault.AddAttribute("ContentType", "application/vnd.openxmlformats-package.relationships+xml")
	tTypes.AddContentTags(tDefault)

	tDefault = docx_xml.GetTag("Default")
	tDefault.AddAttribute("Extension", "xml")
	tDefault.AddAttribute("ContentType", "application/xml")
	tTypes.AddContentTags(tDefault)

	tOverride := docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/word/document.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/word/styles.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/word/settings.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/word/webSettings.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/word/fontTable.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/word/theme/theme1.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.theme+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/docProps/core.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-package.core-properties+xml")
	tTypes.AddContentTags(tOverride)

	tOverride = docx_xml.GetTag("Override")
	tOverride.AddAttribute("PartName", "/docProps/app.xml")
	tOverride.AddAttribute("ContentType", "application/vnd.openxmlformats-officedocument.extended-properties+xml")
	tTypes.AddContentTags(tOverride)

	return tTypes
}

// GetXmlRels dir: _rels filename: .rels
func GetXmlRels() docx_xml.Tag {

	relShips := docx_xml.GetTag("Relationships")
	relShips.AddAttribute("xmlns", "http://schemas.openxmlformats.org/package/2006/relationships")

	relShip3 := docx_xml.GetTag("Relationship")
	relShip3.AddAttribute("Id", "rId3")
	relShip3.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties")
	relShip3.AddAttribute("Target", "docProps/app.xml")

	relShip2 := docx_xml.GetTag("Relationship")
	relShip2.AddAttribute("Id", "rId2")
	relShip2.AddAttribute("Type", "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties")
	relShip2.AddAttribute("Target", "docProps/core.xml")

	relShip1 := docx_xml.GetTag("Relationship")
	relShip1.AddAttribute("Id", "rId1")
	relShip1.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument")
	relShip1.AddAttribute("Target", "word/document.xml")

	relShips.AddContentTags(relShip3, relShip2, relShip1)

	return relShips
}

// GetXmlDocumentXmlRels dir: word\_rels filename: document.xml.rels
func GetXmlDocumentXmlRels() docx_xml.Tag {

	relShips := docx_xml.GetTag("Relationships")
	relShips.AddAttribute("xmlns", "http://schemas.openxmlformats.org/package/2006/relationships")

	relShip3 := docx_xml.GetTag("Relationship")
	relShip3.AddAttribute("Id", "rId3")
	relShip3.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings")
	relShip3.AddAttribute("Target", "webSettings.xml")

	relShip2 := docx_xml.GetTag("Relationship")
	relShip2.AddAttribute("Id", "rId2")
	relShip2.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings")
	relShip2.AddAttribute("Target", "settings.xml")

	relShip1 := docx_xml.GetTag("Relationship")
	relShip1.AddAttribute("Id", "rId1")
	relShip1.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles")
	relShip1.AddAttribute("Target", "styles.xml")

	relShip5 := docx_xml.GetTag("Relationship")
	relShip5.AddAttribute("Id", "rId5")
	relShip5.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme")
	relShip5.AddAttribute("Target", "theme/theme1.xml")

	relShip4 := docx_xml.GetTag("Relationship")
	relShip4.AddAttribute("Id", "rId4")
	relShip4.AddAttribute("Type", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable")
	relShip4.AddAttribute("Target", "fontTable.xml")

	relShips.AddContentTags(relShip3, relShip2, relShip1, relShip5, relShip4)

	return relShips
}

// GetXmlDocPropsApp dir: docProps filename: app.xml
func GetXmlDocPropsApp() docx_xml.Tag {
	properties := docx_xml.GetTag("Properties")
	properties.AddAttribute("xmlns", "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties")
	properties.AddAttribute("xmlns:vt", "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes")

	template := docx_xml.GetTag("Template")
	template.SetContentText("Normal.dotm")

	totalTime := docx_xml.GetTag("TotalTime")
	totalTime.SetContentText("0")

	pages := docx_xml.GetTag("Pages")
	pages.SetContentText("1")

	words := docx_xml.GetTag("Words")
	words.SetContentText("0")

	characters := docx_xml.GetTag("Characters")
	characters.SetContentText("0")

	application := docx_xml.GetTag("Application")
	application.SetContentText("Microsoft Office Word")

	docSecurity := docx_xml.GetTag("DocSecurity")
	docSecurity.SetContentText("0")

	lines := docx_xml.GetTag("Lines")
	lines.SetContentText("0")

	paragraphs := docx_xml.GetTag("Paragraphs")
	paragraphs.SetContentText("0")

	scaleCrop := docx_xml.GetTag("ScaleCrop")
	scaleCrop.SetContentText("false")

	company := docx_xml.GetTag("Company")
	company.SetContentText("")

	linksUpToDate := docx_xml.GetTag("LinksUpToDate")
	linksUpToDate.SetContentText("false")

	charactersWithSpaces := docx_xml.GetTag("CharactersWithSpaces")
	charactersWithSpaces.SetContentText("0")

	sharedDoc := docx_xml.GetTag("SharedDoc")
	sharedDoc.SetContentText("false")

	hyperlinksChanged := docx_xml.GetTag("HyperlinksChanged")
	hyperlinksChanged.SetContentText("false")

	appVersion := docx_xml.GetTag("AppVersion")
	appVersion.SetContentText("16.0000")

	properties.AddContentTags(template, totalTime, pages, words, characters, application, docSecurity, lines,
		paragraphs, scaleCrop, company, linksUpToDate, charactersWithSpaces, sharedDoc, hyperlinksChanged, appVersion)

	return properties
}

// GetXmlDocPropsCore dir: docProps filename: core.xml
func GetXmlDocPropsCore() docx_xml.Tag {
	coreProperties := docx_xml.GetTag("cp:coreProperties")
	coreProperties.AddAttribute("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	coreProperties.AddAttribute("xmlns:dcmitype", "http://purl.org/dc/dcmitype/")
	coreProperties.AddAttribute("xmlns:dcterms", "http://purl.org/dc/terms/")
	coreProperties.AddAttribute("xmlns:dc", "http://purl.org/dc/elements/1.1/")
	coreProperties.AddAttribute("xmlns:cp", "http://schemas.openxmlformats.org/package/2006/metadata/core-properties")

	title := docx_xml.GetTag("dc:title")
	title.SetContentText("")

	subject := docx_xml.GetTag("dc:subject")
	subject.SetContentText("")

	creator := docx_xml.GetTag("dc:creator")
	creator.SetContentText("generation")

	keywords := docx_xml.GetTag("cp:keywords")
	keywords.SetContentText("")

	description := docx_xml.GetTag("dc:description")
	description.SetContentText("")

	lastModifiedBy := docx_xml.GetTag("cp:lastModifiedBy")
	lastModifiedBy.SetContentText("generation")

	revision := docx_xml.GetTag("cp:revision")
	revision.SetContentText("1")

	created := docx_xml.GetTag("dcterms:created")
	created.AddAttribute("xsi:type", "dcterms:W3CDTF")
	created.SetContentText("2022-10-11T16:37:00Z")

	modified := docx_xml.GetTag("dcterms:modified")
	modified.AddAttribute("xsi:type", "dcterms:W3CDTF")
	modified.SetContentText("2022-10-11T16:37:00Z")

	coreProperties.AddContentTags(title, subject, creator, keywords, description,
		lastModifiedBy, revision, created, modified)

	return coreProperties
}

// GetXmlFontTable dir: word filename: fontTable.xml
func GetXmlFontTable() docx_xml.Tag {
	fonts := docx_xml.GetTag("w:fonts")
	fonts.AddAttribute("xmlns:mc", "http://schemas.openxmlformats.org/markup-compatibility/2006")
	fonts.AddAttribute("xmlns:r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships")
	fonts.AddAttribute("xmlns:w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	fonts.AddAttribute("xmlns:w14", "http://schemas.microsoft.com/office/word/2010/wordml")
	fonts.AddAttribute("xmlns:w15", "http://schemas.microsoft.com/office/word/2012/wordml")
	fonts.AddAttribute("xmlns:w16cid", "http://schemas.microsoft.com/office/word/2016/wordml/cid")
	fonts.AddAttribute("xmlns:w16se", "http://schemas.microsoft.com/office/word/2015/wordml/symex")
	fonts.AddAttribute("mc:Ignorable", "w14 w15 w16se w16cid")

	fontCalibri := docx_xml.GetTag("w:font")
	fontCalibri.AddAttribute("w:name", "Calibri")
	{
		panose1 := docx_xml.GetTag("w:panose1")
		panose1.AddAttribute("w:val", "020F0502020204030204")

		charset := docx_xml.GetTag("w:charset")
		charset.AddAttribute("w:val", "CC")

		family := docx_xml.GetTag("w:family")
		family.AddAttribute("w:val", "swiss")

		pitch := docx_xml.GetTag("w:pitch")
		pitch.AddAttribute("w:val", "variable")

		sig := docx_xml.GetTag("w:sig")
		sig.AddAttribute("w:usb0", "E4002EFF")
		sig.AddAttribute("w:usb1", "C000247B")
		sig.AddAttribute("w:usb2", "00000009")
		sig.AddAttribute("w:usb3", "00000000")
		sig.AddAttribute("w:csb0", "000001FF")
		sig.AddAttribute("w:csb1", "00000000")
		fontCalibri.AddContentTags(panose1, charset, family, pitch, sig)
	}

	fontTimesNewRoman := docx_xml.GetTag("w:font")
	fontTimesNewRoman.AddAttribute("w:name", "Times New Roman")
	{
		panose1 := docx_xml.GetTag("w:panose1")
		panose1.AddAttribute("w:val", "02020603050405020304")

		charset := docx_xml.GetTag("w:charset")
		charset.AddAttribute("w:val", "CC")

		family := docx_xml.GetTag("w:family")
		family.AddAttribute("w:val", "roman")

		pitch := docx_xml.GetTag("w:pitch")
		pitch.AddAttribute("w:val", "variable")

		sig := docx_xml.GetTag("w:sig")
		sig.AddAttribute("w:usb0", "E0002EFF")
		sig.AddAttribute("w:usb1", "C000785B")
		sig.AddAttribute("w:usb2", "00000009")
		sig.AddAttribute("w:usb3", "00000000")
		sig.AddAttribute("w:csb0", "000001FF")
		sig.AddAttribute("w:csb1", "00000000")
		fontTimesNewRoman.AddContentTags(panose1, charset, family, pitch, sig)
	}

	fontCalibriLight := docx_xml.GetTag("w:font")
	fontCalibriLight.AddAttribute("w:name", "Calibri Light")
	{
		panose1 := docx_xml.GetTag("w:panose1")
		panose1.AddAttribute("w:val", "020F0302020204030204")

		charset := docx_xml.GetTag("w:charset")
		charset.AddAttribute("w:val", "CC")

		family := docx_xml.GetTag("w:family")
		family.AddAttribute("w:val", "roman")

		pitch := docx_xml.GetTag("w:pitch")
		pitch.AddAttribute("w:val", "variable")

		sig := docx_xml.GetTag("w:sig")
		sig.AddAttribute("w:usb0", "E4002EFF")
		sig.AddAttribute("w:usb1", "C000247B")
		sig.AddAttribute("w:usb2", "00000009")
		sig.AddAttribute("w:usb3", "00000000")
		sig.AddAttribute("w:csb0", "000001FF")
		sig.AddAttribute("w:csb1", "00000000")
		fontCalibriLight.AddContentTags(panose1, charset, family, pitch, sig)
	}

	fonts.AddContentTags(fontCalibri, fontTimesNewRoman, fontCalibriLight)

	return fonts
}

// GetXmlSettings dir: word filename: settings.xml
func GetXmlSettings() docx_xml.Tag {
	settings := docx_xml.GetTag("w:settings")
	settings.AddAttribute("xmlns:mc", "http://schemas.openxmlformats.org/markup-compatibility/2006")
	settings.AddAttribute("xmlns:o", "urn:schemas-microsoft-com:office:office")
	settings.AddAttribute("xmlns:r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships")
	settings.AddAttribute("xmlns:m", "http://schemas.openxmlformats.org/officeDocument/2006/math")
	settings.AddAttribute("xmlns:v", "urn:schemas-microsoft-com:vml")
	settings.AddAttribute("xmlns:w10", "urn:schemas-microsoft-com:office:word")
	settings.AddAttribute("xmlns:w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	settings.AddAttribute("xmlns:w14", "http://schemas.microsoft.com/office/word/2010/wordml")
	settings.AddAttribute("xmlns:w15", "http://schemas.microsoft.com/office/word/2012/wordml")
	settings.AddAttribute("xmlns:w16cid", "http://schemas.microsoft.com/office/word/2016/wordml/cid")
	settings.AddAttribute("xmlns:w16se", "http://schemas.microsoft.com/office/word/2015/wordml/symex")
	settings.AddAttribute("xmlns:sl", "http://schemas.openxmlformats.org/schemaLibrary/2006/main")
	settings.AddAttribute("mc:Ignorable", "w14 w15 w16se w16cid")

	zoom := docx_xml.GetTag("w:zoom")
	zoom.AddAttribute("w:percent", "100")

	defaultTabStop := docx_xml.GetTag("w:defaultTabStop")
	defaultTabStop.AddAttribute("w:val", "708")

	characterSpacingControl := docx_xml.GetTag("w:characterSpacingControl")
	characterSpacingControl.AddAttribute("w:val", "doNotCompress")

	compat := docx_xml.GetTag("w:compat")
	{
		compatSetting := docx_xml.GetTag("w:compatSetting")
		compatSetting.AddAttribute("w:name", "compatibilityMode")
		compatSetting.AddAttribute("w:uri", "http://schemas.microsoft.com/office/word")
		compatSetting.AddAttribute("w:val", "15")
		compat.AddContentTags(compatSetting)
	}
	{
		compatSetting := docx_xml.GetTag("w:compatSetting")
		compatSetting.AddAttribute("w:name", "overrideTableStyleFontSizeAndJustification")
		compatSetting.AddAttribute("w:uri", "http://schemas.microsoft.com/office/word")
		compatSetting.AddAttribute("w:val", "1")
		compat.AddContentTags(compatSetting)
	}
	{
		compatSetting := docx_xml.GetTag("w:compatSetting")
		compatSetting.AddAttribute("w:name", "enableOpenTypeFeatures")
		compatSetting.AddAttribute("w:uri", "http://schemas.microsoft.com/office/word")
		compatSetting.AddAttribute("w:val", "1")
		compat.AddContentTags(compatSetting)
	}
	{
		compatSetting := docx_xml.GetTag("w:compatSetting")
		compatSetting.AddAttribute("w:name", "doNotFlipMirrorIndents")
		compatSetting.AddAttribute("w:uri", "http://schemas.microsoft.com/office/word")
		compatSetting.AddAttribute("w:val", "1")
		compat.AddContentTags(compatSetting)
	}
	{
		compatSetting := docx_xml.GetTag("w:compatSetting")
		compatSetting.AddAttribute("w:name", "differentiateMultirowTableHeaders")
		compatSetting.AddAttribute("w:uri", "http://schemas.microsoft.com/office/word")
		compatSetting.AddAttribute("w:val", "1")
		compat.AddContentTags(compatSetting)
	}
	{
		compatSetting := docx_xml.GetTag("w:compatSetting")
		compatSetting.AddAttribute("w:name", "useWord2013TrackBottomHyphenation")
		compatSetting.AddAttribute("w:uri", "http://schemas.microsoft.com/office/word")
		compatSetting.AddAttribute("w:val", "0")
		compat.AddContentTags(compatSetting)
	}

	rsids := docx_xml.GetTag("w:rsids")
	{
		rsidRoot := docx_xml.GetTag("w:rsidRoot")
		rsidRoot.AddAttribute("w:val", "004E6D86")
		rsids.AddContentTags(rsidRoot)
	}
	{
		rsid := docx_xml.GetTag("w:rsid")
		rsid.AddAttribute("w:val", "003174CB")
		rsids.AddContentTags(rsid)
	}
	{
		rsid := docx_xml.GetTag("w:rsid")
		rsid.AddAttribute("w:val", "004E6D86")
		rsids.AddContentTags(rsid)
	}
	{
		rsid := docx_xml.GetTag("w:rsid")
		rsid.AddAttribute("w:val", "00546596")
		rsids.AddContentTags(rsid)
	}
	{
		rsid := docx_xml.GetTag("w:rsid")
		rsid.AddAttribute("w:val", "005E61D5")
		rsids.AddContentTags(rsid)
	}
	{
		rsid := docx_xml.GetTag("w:rsid")
		rsid.AddAttribute("w:val", "00E46963")
		rsids.AddContentTags(rsid)
	}

	mathPr := docx_xml.GetTag("w:mathPr")
	{
		mathFont := docx_xml.GetTag("m:mathFont")
		mathFont.AddAttribute("m:val", "Cambria Math")

		brkBin := docx_xml.GetTag("m:brkBin")
		brkBin.AddAttribute("m:val", "before")

		brkBinSub := docx_xml.GetTag("m:brkBinSub")
		brkBinSub.AddAttribute("m:val", "--")

		smallFrac := docx_xml.GetTag("m:smallFrac")
		smallFrac.AddAttribute("m:val", "0")

		dispDef := docx_xml.GetTag("m:dispDef")

		lMargin := docx_xml.GetTag("m:lMargin")
		lMargin.AddAttribute("m:val", "0")

		rMargin := docx_xml.GetTag("m:rMargin")
		rMargin.AddAttribute("m:val", "0")

		defJc := docx_xml.GetTag("m:defJc")
		defJc.AddAttribute("m:val", "centerGroup")

		wrapIndent := docx_xml.GetTag("m:wrapIndent")
		wrapIndent.AddAttribute("m:val", "1440")

		intLim := docx_xml.GetTag("m:intLim")
		intLim.AddAttribute("m:val", "subSup")

		naryLim := docx_xml.GetTag("m:naryLim")
		naryLim.AddAttribute("m:val", "undOvr")

		mathPr.AddContentTags(mathFont, brkBin, brkBinSub, smallFrac, dispDef,
			lMargin, rMargin, defJc, wrapIndent, intLim, intLim, naryLim)
	}
	themeFontLang := docx_xml.GetTag("w:themeFontLang")
	themeFontLang.AddAttribute("w:val", "ru-RU")

	clrSchemeMapping := docx_xml.GetTag("w:clrSchemeMapping")
	clrSchemeMapping.AddAttribute("w:bg1", "light1")
	clrSchemeMapping.AddAttribute("w:t1", "dark1")
	clrSchemeMapping.AddAttribute("w:bg2", "light2")
	clrSchemeMapping.AddAttribute("w:t2", "dark2")
	clrSchemeMapping.AddAttribute("w:accent1", "accent1")
	clrSchemeMapping.AddAttribute("w:accent2", "accent2")
	clrSchemeMapping.AddAttribute("w:accent3", "accent3")
	clrSchemeMapping.AddAttribute("w:accent4", "accent4")
	clrSchemeMapping.AddAttribute("w:accent5", "accent5")
	clrSchemeMapping.AddAttribute("w:accent6", "accent6")
	clrSchemeMapping.AddAttribute("w:hyperlink", "hyperlink")
	clrSchemeMapping.AddAttribute("w:followedHyperlink", "followedHyperlink")

	shapeDefaults := docx_xml.GetTag("w:shapeDefaults")
	{
		oShapedefaults := docx_xml.GetTag("o:shapedefaults")
		oShapedefaults.AddAttribute("v:ext", "edit")
		oShapedefaults.AddAttribute("spidmax", "1026")

		shapelayout := docx_xml.GetTag("o:shapelayout")
		shapelayout.AddAttribute("v:ext", "edit")

		idmap := docx_xml.GetTag("o:idmap")
		idmap.AddAttribute("v:ext", "edit")
		idmap.AddAttribute("data", "1")

		shapelayout.AddContentTags(idmap)

		shapeDefaults.AddContentTags(oShapedefaults, shapelayout)
	}

	decimalSymbol := docx_xml.GetTag("w:decimalSymbol")
	decimalSymbol.AddAttribute("w:val", ",")

	listSeparator := docx_xml.GetTag("w:listSeparator")
	listSeparator.AddAttribute("w:val", ";")

	w14docId := docx_xml.GetTag("w14:docId")
	w14docId.AddAttribute("w14:val", "6984F6B4")

	chartTrackingRefBased := docx_xml.GetTag("w15:chartTrackingRefBased")

	w15docId := docx_xml.GetTag("w15:docId")
	w15docId.AddAttribute("w15:val", "{7AD3CFB2-A0F6-46D5-B941-44B34EBB6463}")

	settings.AddContentTags(zoom, defaultTabStop, characterSpacingControl, compat, rsids, mathPr, themeFontLang,
		clrSchemeMapping, shapeDefaults, decimalSymbol, listSeparator, w14docId, chartTrackingRefBased, w15docId)

	return settings
}

// GetXmlStyles dir: word filename: styles.xml
func GetXmlStyles() docx_xml.Tag {
	styles := docx_xml.GetTag("w:styles")
	styles.AddAttribute("xmlns:mc", "http://schemas.openxmlformats.org/markup-compatibility/2006")
	styles.AddAttribute("xmlns:r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships")
	styles.AddAttribute("xmlns:w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	styles.AddAttribute("xmlns:w14", "http://schemas.microsoft.com/office/word/2010/wordml")
	styles.AddAttribute("xmlns:w15", "http://schemas.microsoft.com/office/word/2012/wordml")
	styles.AddAttribute("xmlns:w16cid", "http://schemas.microsoft.com/office/word/2016/wordml/cid")
	styles.AddAttribute("xmlns:w16se", "http://schemas.microsoft.com/office/word/2015/wordml/symex")
	styles.AddAttribute("mc:Ignorable", "w14 w15 w16se w16cid")

	docDefaults := docx_xml.GetTag("w:docDefaults")
	{
		rPrDefault := docx_xml.GetTag("w:rPrDefault")
		rpr := docx_xml.GetTag("w:rPr")
		{
			rFonts := docx_xml.GetTag("w:rFonts")
			rFonts.AddAttribute("w:asciiTheme", "minorHAnsi")
			rFonts.AddAttribute("w:eastAsiaTheme", "minorHAnsi")
			rFonts.AddAttribute("w:hAnsiTheme", "minorHAnsi")
			rFonts.AddAttribute("w:cstheme", "minorBidi")
			sz := docx_xml.GetTag("w:sz")
			sz.AddAttribute("w:val", "22")
			szCs := docx_xml.GetTag("w:szCs")
			szCs.AddAttribute("w:val", "22")
			lang := docx_xml.GetTag("w:lang")
			lang.AddAttribute("w:val", "ru-RU")
			lang.AddAttribute("w:eastAsia", "en-US")
			lang.AddAttribute("w:bidi", "ar-SA")
			rpr.AddContentTags(rFonts, sz, szCs, lang)
		}
		rPrDefault.AddContentTags(rpr)
		docDefaults.AddContentTags(rPrDefault)
	}
	{

		pPrDefault := docx_xml.GetTag("w:pPrDefault")
		pPr := docx_xml.GetTag("w:pPr")
		spacing := docx_xml.GetTag("w:spacing")
		spacing.AddAttribute("w:after", "160")
		spacing.AddAttribute("w:line", "259")
		spacing.AddAttribute("w:lineRule", "auto")
		pPr.AddContentTags(spacing)

		pPrDefault.AddContentTags(pPr)
		docDefaults.AddContentTags(pPrDefault)
	}

	latentStyles := docx_xml.GetTag("w:latentStyles")
	latentStyles.AddAttribute("w:defLockedState", "0")
	latentStyles.AddAttribute("w:defUIPriority", "99")
	latentStyles.AddAttribute("w:defSemiHidden", "0")
	latentStyles.AddAttribute("w:defUnhideWhenUsed", "0")
	latentStyles.AddAttribute("w:defQFormat", "0")
	latentStyles.AddAttribute("w:count", "375")
	{
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Normal")
			lsdException.AddAttribute("w:uiPriority", "0")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 6")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 7")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 8")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "heading 9")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "9")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 6")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 7")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 8")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index 9")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 6")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 7")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 8")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toc 9")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Normal Indent")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "footnote text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "annotation text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "header")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "footer")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "index heading")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "caption")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "35")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "table of figures")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "envelope address")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "envelope return")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "footnote reference")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "annotation reference")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "line number")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "page number")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "endnote reference")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "endnote text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "table of authorities")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "macro")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "toa heading")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Bullet")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Number")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Bullet 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Bullet 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Bullet 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Bullet 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Number 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Number 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Number 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Number 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Title")
			lsdException.AddAttribute("w:uiPriority", "10")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Closing")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Signature")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Default Paragraph Font")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text Indent")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Continue")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Continue 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Continue 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Continue 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Continue 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Message Header")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Subtitle")
			lsdException.AddAttribute("w:uiPriority", "11")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Salutation")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Date")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text First Indent")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text First Indent 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Note Heading")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text Indent 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Body Text Indent 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Block Text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Hyperlink")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "FollowedHyperlink")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Strong")
			lsdException.AddAttribute("w:uiPriority", "22")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Emphasis")
			lsdException.AddAttribute("w:uiPriority", "20")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Document Map")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Plain Text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "E-mail Signature")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Top of Form")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Bottom of Form")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Normal (Web")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Acronym")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Address")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Cite")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Code")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Definition")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Keyboard")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Preformatted")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Sample")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Typewriter")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "HTML Variable")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Normal Table")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "annotation subject")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "No List")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Outline List 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Outline List 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Outline List 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Simple 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Simple 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Simple 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Classic 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Classic 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Classic 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Classic 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Colorful 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Colorful 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Colorful 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Columns 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Columns 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Columns 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Columns 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Columns 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 6")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 7")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid 8")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 4")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 5")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 6")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 7")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table List 8")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table 3D effects 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table 3D effects 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table 3D effects 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Contemporary")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Elegant")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Professional")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Subtle 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Subtle 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Web 1")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Web 2")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Web 3")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Balloon Text")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Grid")
			lsdException.AddAttribute("w:uiPriority", "39")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Table Theme")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Placeholder Text")
			lsdException.AddAttribute("w:semiHidden", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "No Spacing")
			lsdException.AddAttribute("w:uiPriority", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading Accent 1")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List Accent 1")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid Accent 1")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Revision")
			lsdException.AddAttribute("w:semiHidden", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Paragraph")
			lsdException.AddAttribute("w:uiPriority", "34")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Quote")
			lsdException.AddAttribute("w:uiPriority", "29")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Intense Quote")
			lsdException.AddAttribute("w:uiPriority", "30")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List Accent 1")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading Accent 1")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List Accent 1")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid Accent 1")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading Accent 2")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List Accent 2")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid Accent 2")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List Accent 2")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading Accent 2")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List Accent 2")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid Accent 2")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading Accent 3")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List Accent 3")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid Accent 3")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List Accent 3")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading Accent 3")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List Accent 3")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid Accent 3")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading Accent 4")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List Accent 4")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid Accent 4")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List Accent 4")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading Accent 4")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List Accent 4")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid Accent 4")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading Accent 5")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List Accent 5")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid Accent 5")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List Accent 5")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading Accent 5")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List Accent 5")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid Accent 5")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Shading Accent 6")
			lsdException.AddAttribute("w:uiPriority", "60")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light List Accent 6")
			lsdException.AddAttribute("w:uiPriority", "61")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Light Grid Accent 6")
			lsdException.AddAttribute("w:uiPriority", "62")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 1 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "63")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Shading 2 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "64")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 1 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "65")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium List 2 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "66")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 1 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "67")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 2 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "68")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Medium Grid 3 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "69")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Dark List Accent 6")
			lsdException.AddAttribute("w:uiPriority", "70")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Shading Accent 6")
			lsdException.AddAttribute("w:uiPriority", "71")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful List Accent 6")
			lsdException.AddAttribute("w:uiPriority", "72")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Colorful Grid Accent 6")
			lsdException.AddAttribute("w:uiPriority", "73")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Subtle Emphasis")
			lsdException.AddAttribute("w:uiPriority", "19")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Intense Emphasis")
			lsdException.AddAttribute("w:uiPriority", "21")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Subtle Reference")
			lsdException.AddAttribute("w:uiPriority", "31")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Intense Reference")
			lsdException.AddAttribute("w:uiPriority", "32")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Book Title")
			lsdException.AddAttribute("w:uiPriority", "33")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Bibliography")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "37")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "TOC Heading")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:uiPriority", "39")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			lsdException.AddAttribute("w:qFormat", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Plain Table 1")
			lsdException.AddAttribute("w:uiPriority", "41")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Plain Table 2")
			lsdException.AddAttribute("w:uiPriority", "42")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Plain Table 3")
			lsdException.AddAttribute("w:uiPriority", "43")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Plain Table 4")
			lsdException.AddAttribute("w:uiPriority", "44")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Plain Table 5")
			lsdException.AddAttribute("w:uiPriority", "45")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table Light")
			lsdException.AddAttribute("w:uiPriority", "40")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light Accent 1")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark Accent 1")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful Accent 1")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful Accent 1")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light Accent 2")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark Accent 2")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful Accent 2")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful Accent 2")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light Accent 3")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark Accent 3")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful Accent 3")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful Accent 3")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light Accent 4")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark Accent 4")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful Accent 4")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful Accent 4")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light Accent 5")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark Accent 5")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful Accent 5")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful Accent 5")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 1 Light Accent 6")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 2 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 3 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 4 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 5 Dark Accent 6")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 6 Colorful Accent 6")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Grid Table 7 Colorful Accent 6")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light Accent 1")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4 Accent 1")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark Accent 1")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful Accent 1")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful Accent 1")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light Accent 2")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4 Accent 2")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark Accent 2")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful Accent 2")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful Accent 2")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light Accent 3")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4 Accent 3")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark Accent 3")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful Accent 3")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful Accent 3")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light Accent 4")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4 Accent 4")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark Accent 4")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful Accent 4")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful Accent 4")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light Accent 5")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4 Accent 5")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark Accent 5")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful Accent 5")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful Accent 5")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 1 Light Accent 6")
			lsdException.AddAttribute("w:uiPriority", "46")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 2 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "47")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 3 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "48")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 4 Accent 6")
			lsdException.AddAttribute("w:uiPriority", "49")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 5 Dark Accent 6")
			lsdException.AddAttribute("w:uiPriority", "50")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 6 Colorful Accent 6")
			lsdException.AddAttribute("w:uiPriority", "51")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "List Table 7 Colorful Accent 6")
			lsdException.AddAttribute("w:uiPriority", "52")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Mention")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Smart Hyperlink")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Hashtag")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
		{
			lsdException := docx_xml.GetTag("w:lsdException")
			lsdException.AddAttribute("w:name", "Unresolved Mention")
			lsdException.AddAttribute("w:semiHidden", "1")
			lsdException.AddAttribute("w:unhideWhenUsed", "1")
			latentStyles.AddContentTags(lsdException)
		}
	}

	styleParagraph := docx_xml.GetTag("w:style")
	styleParagraph.AddAttribute("w:type", "paragraph")
	styleParagraph.AddAttribute("w:default", "1")
	styleParagraph.AddAttribute("w:styleId", "a")
	{
		name := docx_xml.GetTag("w:name")
		name.AddAttribute("w:val", "Normal")
		qFormat := docx_xml.GetTag("w:qFormat")
		styleParagraph.AddContentTags(name, qFormat)
	}

	styleCharacter := docx_xml.GetTag("w:style")
	styleCharacter.AddAttribute("w:type", "character")
	styleCharacter.AddAttribute("w:default", "1")
	styleCharacter.AddAttribute("w:styleId", "a0")
	{
		name := docx_xml.GetTag("w:name")
		name.AddAttribute("w:val", "Default Paragraph Font")
		uiPriority := docx_xml.GetTag("w:uiPriority")
		uiPriority.AddAttribute("w:val", "1")
		semiHidden := docx_xml.GetTag("w:semiHidden")
		unhideWhenUsed := docx_xml.GetTag("w:unhideWhenUsed")
		styleCharacter.AddContentTags(name, uiPriority, semiHidden, unhideWhenUsed)
	}

	styleTable := docx_xml.GetTag("w:style")
	styleTable.AddAttribute("w:type", "table")
	styleTable.AddAttribute("w:default", "1")
	styleTable.AddAttribute("w:styleId", "a1")
	{
		name := docx_xml.GetTag("w:name")
		name.AddAttribute("w:val", "Normal Table")
		uiPriority := docx_xml.GetTag("w:uiPriority")
		uiPriority.AddAttribute("w:val", "99")
		semiHidden := docx_xml.GetTag("w:semiHidden")
		unhideWhenUsed := docx_xml.GetTag("w:unhideWhenUsed")

		tblPr := docx_xml.GetTag("w:tblPr")
		{
			tblInd := docx_xml.GetTag("w:tblInd")
			tblInd.AddAttribute("w:w", "0")
			tblInd.AddAttribute("w:type", "dxa")
			tblCellMar := docx_xml.GetTag("w:tblCellMar")
			{
				top := docx_xml.GetTag("w:top")
				top.AddAttribute("w:w", "0")
				top.AddAttribute("w:type", "dxa")

				left := docx_xml.GetTag("w:left")
				left.AddAttribute("w:w", "108")
				left.AddAttribute("w:type", "dxa")

				bottom := docx_xml.GetTag("w:bottom")
				bottom.AddAttribute("w:w", "0")
				bottom.AddAttribute("w:type", "dxa")

				right := docx_xml.GetTag("w:right")
				right.AddAttribute("w:w", "108")
				right.AddAttribute("w:type", "dxa")

				tblCellMar.AddContentTags(top, left, bottom, right)
			}
			tblPr.AddContentTags(tblInd, tblCellMar)
		}

		styleTable.AddContentTags(name, uiPriority, semiHidden, unhideWhenUsed, tblPr)
	}

	styleNumbering := docx_xml.GetTag("w:style")
	styleNumbering.AddAttribute("w:type", "numbering")
	styleNumbering.AddAttribute("w:default", "1")
	styleNumbering.AddAttribute("w:styleId", "a2")
	{
		name := docx_xml.GetTag("w:name")
		name.AddAttribute("w:val", "No List")
		uiPriority := docx_xml.GetTag("w:uiPriority")
		uiPriority.AddAttribute("w:val", "99")
		semiHidden := docx_xml.GetTag("w:semiHidden")
		unhideWhenUsed := docx_xml.GetTag("w:unhideWhenUsed")
		styleNumbering.AddContentTags(name, uiPriority, semiHidden, unhideWhenUsed)
	}

	styles.AddContentTags(docDefaults, latentStyles, styleParagraph, styleCharacter, styleTable, styleNumbering)
	return styles
}

// GetXmlStyles dir: word filename: webSettings.xml
func GetXmlWebSettings() docx_xml.Tag {
	webSettings := docx_xml.GetTag("w:webSettings")
	webSettings.AddAttribute("xmlns:mc", "http://schemas.openxmlformats.org/markup-compatibility/2006")
	webSettings.AddAttribute("xmlns:r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships")
	webSettings.AddAttribute("xmlns:w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	webSettings.AddAttribute("xmlns:w14", "http://schemas.microsoft.com/office/word/2010/wordml")
	webSettings.AddAttribute("xmlns:w15", "http://schemas.microsoft.com/office/word/2012/wordml")
	webSettings.AddAttribute("xmlns:w16cid", "http://schemas.microsoft.com/office/word/2016/wordml/cid")
	webSettings.AddAttribute("xmlns:w16se", "http://schemas.microsoft.com/office/word/2015/wordml/symex")
	webSettings.AddAttribute("mc:Ignorable", "w14 w15 w16se w16cid")

	optimizeForBrowser := docx_xml.GetTag("w:optimizeForBrowser")
	allowPNG := docx_xml.GetTag("w:allowPNG")
	webSettings.AddContentTags(optimizeForBrowser, allowPNG)
	return webSettings
}

func GetXmlWordTheme() docx_xml.Tag {
	theme := docx_xml.GetTag("a:theme")
	theme.AddAttribute("xmlns:a", "http://schemas.openxmlformats.org/drawingml/2006/main")
	theme.AddAttribute("name", "Тема Office")

	themeElements := docx_xml.GetTag("a:themeElements")
	{
		clrScheme := docx_xml.GetTag("a:clrScheme")
		clrScheme.AddAttribute("name", "Стандартная")

		dk1 := docx_xml.GetTag("a:dk1")
		{
			sysClr := docx_xml.GetTag("a:sysClr")
			sysClr.AddAttribute("val", "windowText")
			sysClr.AddAttribute("lastClr", "000000")
			dk1.AddContentTags(sysClr)
		}

		lt1 := docx_xml.GetTag("a:lt1")
		{

			sysClr := docx_xml.GetTag("a:sysClr")
			sysClr.AddAttribute("val", "window")
			sysClr.AddAttribute("lastClr", "FFFFFF")
			lt1.AddContentTags(sysClr)
		}

		dk2 := docx_xml.GetTag("a:dk2")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "44546A")
			dk2.AddContentTags(srgbClr)
		}

		lt2 := docx_xml.GetTag("a:lt2")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "E7E6E6")
			lt2.AddContentTags(srgbClr)
		}

		accent1 := docx_xml.GetTag("a:accent1")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "4472C4")
			accent1.AddContentTags(srgbClr)
		}

		accent2 := docx_xml.GetTag("a:accent2")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "ED7D31")
			accent2.AddContentTags(srgbClr)
		}

		accent3 := docx_xml.GetTag("a:accent3")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "A5A5A5")
			accent3.AddContentTags(srgbClr)
		}

		accent4 := docx_xml.GetTag("a:accent4")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "FFC000")
			accent4.AddContentTags(srgbClr)
		}

		accent5 := docx_xml.GetTag("a:accent5")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "FFC000")
			accent5.AddContentTags(srgbClr)
		}
		accent6 := docx_xml.GetTag("a:accent6")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "70AD47")
			accent6.AddContentTags(srgbClr)
		}
		hLink := docx_xml.GetTag("a:hlink")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "0563C1")
			hLink.AddContentTags(srgbClr)
		}
		folHLink := docx_xml.GetTag("a:folHlink")
		{
			srgbClr := docx_xml.GetTag("a:srgbClr")
			srgbClr.AddAttribute("val", "954F72")
			folHLink.AddContentTags(srgbClr)
		}

		clrScheme.AddContentTags(dk1, lt1, dk2, lt2, accent1, accent2, accent3, accent4, accent5, accent6, hLink, folHLink)
		themeElements.AddContentTags(clrScheme)
	}
	{
		fontScheme := docx_xml.GetTag("a:fontScheme")
		fontScheme.AddAttribute("name", "Стандартная")

		{
			majorFont := docx_xml.GetTag("a:majorFont")
			latin := docx_xml.GetTag("a:latin")
			latin.AddAttribute("typeface", "Calibri Light")
			latin.AddAttribute("panose", "020F0302020204030204")
			ea := docx_xml.GetTag("a:ea")
			ea.AddAttribute("typeface", "")
			cs := docx_xml.GetTag("a:cs")
			cs.AddAttribute("typeface", "")
			majorFont.AddContentTags(latin, ea, cs)

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Jpan")
				font.AddAttribute("typeface", "游ゴシック Light")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hang")
				font.AddAttribute("typeface", "맑은 고딕")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hans")
				font.AddAttribute("typeface", "等线 Light")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hant")
				font.AddAttribute("typeface", "新細明體")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Arab")
				font.AddAttribute("typeface", "Times New Roman")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hebr")
				font.AddAttribute("typeface", "Times New Roman")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Thai")
				font.AddAttribute("typeface", "Angsana New")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Ethi")
				font.AddAttribute("typeface", "Nyala")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Beng")
				font.AddAttribute("typeface", "Vrinda")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Gujr")
				font.AddAttribute("typeface", "Shruti")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Khmr")
				font.AddAttribute("typeface", "MoolBoran")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Knda")
				font.AddAttribute("typeface", "Tunga")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Guru")
				font.AddAttribute("typeface", "Raavi")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Cans")
				font.AddAttribute("typeface", "Euphemia")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Cher")
				font.AddAttribute("typeface", "Plantagenet Cherokee")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Yiii")
				font.AddAttribute("typeface", "Microsoft Yi Baiti")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Tibt")
				font.AddAttribute("typeface", "Microsoft Himalaya")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Thaa")
				font.AddAttribute("typeface", "MV Boli")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Deva")
				font.AddAttribute("typeface", "Mangal")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Telu")
				font.AddAttribute("typeface", "Gautami")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Taml")
				font.AddAttribute("typeface", "Latha")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syrc")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Orya")
				font.AddAttribute("typeface", "Kalinga")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Mlym")
				font.AddAttribute("typeface", "Kartika")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Laoo")
				font.AddAttribute("typeface", "DokChampa")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Sinh")
				font.AddAttribute("typeface", "Iskoola Pota")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Mong")
				font.AddAttribute("typeface", "Mongolian Baiti")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Viet")
				font.AddAttribute("typeface", "Times New Roman")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Uigh")
				font.AddAttribute("typeface", "Microsoft Uighur")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Geor")
				font.AddAttribute("typeface", "Sylfaen")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Armn")
				font.AddAttribute("typeface", "Arial")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Bugi")
				font.AddAttribute("typeface", "Leelawadee UI")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Bopo")
				font.AddAttribute("typeface", "Microsoft JhengHei")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Java")
				font.AddAttribute("typeface", "Javanese Text")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Lisu")
				font.AddAttribute("typeface", "Segoe UI")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Mymr")
				font.AddAttribute("typeface", "Myanmar Text")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Nkoo")
				font.AddAttribute("typeface", "Ebrima")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Olck")
				font.AddAttribute("typeface", "Nirmala UI")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Osma")
				font.AddAttribute("typeface", "Ebrima")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Phag")
				font.AddAttribute("typeface", "Phagspa")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syrn")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syrj")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syre")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Sora")
				font.AddAttribute("typeface", "Nirmala UI")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Tale")
				font.AddAttribute("typeface", "Microsoft Tai Le")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Talu")
				font.AddAttribute("typeface", "Microsoft New Tai Lue")
				majorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Tfng")
				font.AddAttribute("typeface", "Ebrima")
				majorFont.AddContentTags(font)
			}
			fontScheme.AddContentTags(majorFont)

		}
		{
			minorFont := docx_xml.GetTag("a:minorFont")
			latin := docx_xml.GetTag("a:latin")
			latin.AddAttribute("typeface", "Calibri")
			latin.AddAttribute("panose", "020F0502020204030204")
			ea := docx_xml.GetTag("a:ea")
			ea.AddAttribute("typeface", "")
			cs := docx_xml.GetTag("a:cs")
			cs.AddAttribute("typeface", "")
			minorFont.AddContentTags(latin, ea, cs)

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Jpan")
				font.AddAttribute("typeface", "游明朝")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hang")
				font.AddAttribute("typeface", "맑은 고딕")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hans")
				font.AddAttribute("typeface", "等线")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hant")
				font.AddAttribute("typeface", "新細明體")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Arab")
				font.AddAttribute("typeface", "Arial")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Hebr")
				font.AddAttribute("typeface", "Arial")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Thai")
				font.AddAttribute("typeface", "Cordia New")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Ethi")
				font.AddAttribute("typeface", "Nyala")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Beng")
				font.AddAttribute("typeface", "Vrinda")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Gujr")
				font.AddAttribute("typeface", "Shruti")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Khmr")
				font.AddAttribute("typeface", "DaunPenh")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Knda")
				font.AddAttribute("typeface", "Tunga")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Guru")
				font.AddAttribute("typeface", "Raavi")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Cans")
				font.AddAttribute("typeface", "Euphemia")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Cher")
				font.AddAttribute("typeface", "Plantagenet Cherokee")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Yiii")
				font.AddAttribute("typeface", "Microsoft Yi Baiti")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Tibt")
				font.AddAttribute("typeface", "Microsoft Himalaya")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Thaa")
				font.AddAttribute("typeface", "MV Boli")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Deva")
				font.AddAttribute("typeface", "Mangal")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Telu")
				font.AddAttribute("typeface", "Gautami")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Taml")
				font.AddAttribute("typeface", "Latha")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syrc")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Orya")
				font.AddAttribute("typeface", "Kalinga")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Mlym")
				font.AddAttribute("typeface", "Kartika")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Laoo")
				font.AddAttribute("typeface", "DokChampa")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Sinh")
				font.AddAttribute("typeface", "Iskoola Pota")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Mong")
				font.AddAttribute("typeface", "Mongolian Baiti")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Viet")
				font.AddAttribute("typeface", "Arial")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Uigh")
				font.AddAttribute("typeface", "Microsoft Uighur")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Geor")
				font.AddAttribute("typeface", "Sylfaen")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Armn")
				font.AddAttribute("typeface", "Arial")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Bugi")
				font.AddAttribute("typeface", "Leelawadee UI")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Bopo")
				font.AddAttribute("typeface", "Microsoft JhengHei")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Java")
				font.AddAttribute("typeface", "Javanese Text")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Lisu")
				font.AddAttribute("typeface", "Segoe UI")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Mymr")
				font.AddAttribute("typeface", "Myanmar Text")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Nkoo")
				font.AddAttribute("typeface", "Ebrima")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Olck")
				font.AddAttribute("typeface", "Nirmala UI")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Osma")
				font.AddAttribute("typeface", "Ebrima")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Phag")
				font.AddAttribute("typeface", "Phagspa")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syrn")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syrj")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Syre")
				font.AddAttribute("typeface", "Estrangelo Edessa")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Sora")
				font.AddAttribute("typeface", "Nirmala UI")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Tale")
				font.AddAttribute("typeface", "Microsoft Tai Le")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Talu")
				font.AddAttribute("typeface", "Microsoft New Tai Lue")
				minorFont.AddContentTags(font)
			}

			{
				font := docx_xml.GetTag("a:font")
				font.AddAttribute("script", "Tfng")
				font.AddAttribute("typeface", "Ebrima")
				minorFont.AddContentTags(font)
			}
			fontScheme.AddContentTags(minorFont)
		}
		themeElements.AddContentTags(fontScheme)
	}
	{
		fmtScheme := docx_xml.GetTag("a:fmtScheme")
		fmtScheme.AddAttribute("name", "Стандартная")
		{
			fillStyleLst := docx_xml.GetTag("a:fillStyleLst")
			{
				solidFill := docx_xml.GetTag("a:solidFill")
				schemeClr := docx_xml.GetTag("a:schemeClr")
				schemeClr.AddAttribute("val", "phClr")
				solidFill.AddContentTags(schemeClr)
				fillStyleLst.AddContentTags(solidFill)
			}
			{
				gradFill := docx_xml.GetTag("a:gradFill")
				gradFill.AddAttribute("rotWithShape", "1")
				{
					gsLst := docx_xml.GetTag("a:gsLst")
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "0")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "110000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "105000")
						tint := docx_xml.GetTag("a:tint")
						tint.AddAttribute("val", "67000")

						schemeClr.AddContentTags(lumMod, satMod, tint)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "50000")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "105000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "103000")
						tint := docx_xml.GetTag("a:tint")
						tint.AddAttribute("val", "73000")

						schemeClr.AddContentTags(lumMod, satMod, tint)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "100000")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "105000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "109000")
						tint := docx_xml.GetTag("a:tint")
						tint.AddAttribute("val", "81000")

						schemeClr.AddContentTags(lumMod, satMod, tint)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					lin := docx_xml.GetTag("a:lin")
					lin.AddAttribute("ang", "5400000")
					lin.AddAttribute("scaled", "0")

					gradFill.AddContentTags(gsLst, lin)
				}
				fillStyleLst.AddContentTags(gradFill)
			}
			{
				gradFill := docx_xml.GetTag("a:gradFill")
				gradFill.AddAttribute("rotWithShape", "1")
				{
					gsLst := docx_xml.GetTag("a:gsLst")
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "0")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "103000")
						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "102000")
						tint := docx_xml.GetTag("a:tint")
						tint.AddAttribute("val", "94000")

						schemeClr.AddContentTags(satMod, lumMod, tint)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "50000")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "110000")
						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "100000")
						shade := docx_xml.GetTag("a:shade")
						shade.AddAttribute("val", "100000")

						schemeClr.AddContentTags(satMod, lumMod, shade)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "100000")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "99000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "120000")
						shade := docx_xml.GetTag("a:shade")
						shade.AddAttribute("val", "78000")

						schemeClr.AddContentTags(lumMod, satMod, shade)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					lin := docx_xml.GetTag("a:lin")
					lin.AddAttribute("ang", "5400000")
					lin.AddAttribute("scaled", "0")

					gradFill.AddContentTags(gsLst, lin)
				}
				fillStyleLst.AddContentTags(gradFill)
			}

			fmtScheme.AddContentTags(fillStyleLst)
		}
		{
			lnStyleLst := docx_xml.GetTag("a:lnStyleLst")
			{
				ln := docx_xml.GetTag("a:ln")
				ln.AddAttribute("w", "6350")
				ln.AddAttribute("cap", "flat")
				ln.AddAttribute("cmpd", "sng")
				ln.AddAttribute("algn", "ctr")
				solidFill := docx_xml.GetTag("a:solidFill")
				{
					schemeClr := docx_xml.GetTag("a:schemeClr")
					schemeClr.AddAttribute("val", "phClr")
					solidFill.AddContentTags(schemeClr)
				}
				prstDash := docx_xml.GetTag("a:prstDash")
				prstDash.AddAttribute("val", "solid")
				miter := docx_xml.GetTag("a:miter")
				miter.AddAttribute("lim", "800000")

				ln.AddContentTags(solidFill, prstDash, miter)
				lnStyleLst.AddContentTags(ln)
			}
			{
				ln := docx_xml.GetTag("a:ln")
				ln.AddAttribute("w", "12700")
				ln.AddAttribute("cap", "flat")
				ln.AddAttribute("cmpd", "sng")
				ln.AddAttribute("algn", "ctr")
				solidFill := docx_xml.GetTag("a:solidFill")
				{
					schemeClr := docx_xml.GetTag("a:schemeClr")
					schemeClr.AddAttribute("val", "phClr")
					solidFill.AddContentTags(schemeClr)
				}
				prstDash := docx_xml.GetTag("a:prstDash")
				prstDash.AddAttribute("val", "solid")
				miter := docx_xml.GetTag("a:miter")
				miter.AddAttribute("lim", "800000")

				ln.AddContentTags(solidFill, prstDash, miter)
				lnStyleLst.AddContentTags(ln)
			}
			{
				ln := docx_xml.GetTag("a:ln")
				ln.AddAttribute("w", "19050")
				ln.AddAttribute("cap", "flat")
				ln.AddAttribute("cmpd", "sng")
				ln.AddAttribute("algn", "ctr")
				solidFill := docx_xml.GetTag("a:solidFill")
				{
					schemeClr := docx_xml.GetTag("a:schemeClr")
					schemeClr.AddAttribute("val", "phClr")
					solidFill.AddContentTags(schemeClr)
				}
				prstDash := docx_xml.GetTag("a:prstDash")
				prstDash.AddAttribute("val", "solid")
				miter := docx_xml.GetTag("a:miter")
				miter.AddAttribute("lim", "800000")

				ln.AddContentTags(solidFill, prstDash, miter)
				lnStyleLst.AddContentTags(ln)
			}
			fmtScheme.AddContentTags(lnStyleLst)
		}
		{
			effectStyleLst := docx_xml.GetTag("a:effectStyleLst")
			{
				effectStyle := docx_xml.GetTag("a:effectStyle")
				effectLst := docx_xml.GetTag("a:effectLst")
				effectStyle.AddContentTags(effectLst)
				effectStyleLst.AddContentTags(effectStyle)
			}
			{
				effectStyle := docx_xml.GetTag("a:effectStyle")
				effectLst := docx_xml.GetTag("a:effectLst")
				effectStyle.AddContentTags(effectLst)
				effectStyleLst.AddContentTags(effectStyle)
			}
			{
				effectStyle := docx_xml.GetTag("a:effectStyle")
				effectLst := docx_xml.GetTag("a:effectLst")
				{
					outerShdw := docx_xml.GetTag("a:outerShdw")
					outerShdw.AddAttribute("blurRad", "57150")
					outerShdw.AddAttribute("dist", "19050")
					outerShdw.AddAttribute("dir", "5400000")
					outerShdw.AddAttribute("algn", "ctr")
					outerShdw.AddAttribute("rotWithShape", "0")
					{
						srgbClr := docx_xml.GetTag("a:srgbClr")
						srgbClr.AddAttribute("val", "000000")
						{
							alpha := docx_xml.GetTag("a:alpha")
							alpha.AddAttribute("val", "63000")
							srgbClr.AddContentTags(alpha)
						}
						outerShdw.AddContentTags(srgbClr)
					}
					effectLst.AddContentTags(outerShdw)
				}
				effectStyle.AddContentTags(effectLst)
				effectStyleLst.AddContentTags(effectStyle)
			}
			fmtScheme.AddContentTags(effectStyleLst)
		}
		{
			bgFillStyleLst := docx_xml.GetTag("a:bgFillStyleLst")
			{
				solidFill := docx_xml.GetTag("a:solidFill")
				schemeClr := docx_xml.GetTag("a:schemeClr")
				schemeClr.AddAttribute("val", "phClr")
				solidFill.AddContentTags(schemeClr)
				bgFillStyleLst.AddContentTags(solidFill)
			}
			{
				solidFill := docx_xml.GetTag("a:solidFill")
				{
					schemeClr := docx_xml.GetTag("a:schemeClr")
					schemeClr.AddAttribute("val", "phClr")

					tint := docx_xml.GetTag("a:tint")
					tint.AddAttribute("val", "95000")
					satMod := docx_xml.GetTag("a:satMod")
					satMod.AddAttribute("val", "170000")
					schemeClr.AddContentTags(tint, satMod)

					solidFill.AddContentTags(schemeClr)
				}
				bgFillStyleLst.AddContentTags(solidFill)
			}
			{
				gradFill := docx_xml.GetTag("a:gradFill")
				gradFill.AddAttribute("rotWithShape", "1")
				{
					gsLst := docx_xml.GetTag("a:gsLst")
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "0")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						tint := docx_xml.GetTag("a:tint")
						tint.AddAttribute("val", "93000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "150000")
						shade := docx_xml.GetTag("a:shade")
						shade.AddAttribute("val", "98000")
						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "102000")

						schemeClr.AddContentTags(tint, satMod, shade, lumMod)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "50000")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						tint := docx_xml.GetTag("a:tint")
						tint.AddAttribute("val", "98000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "130000")
						shade := docx_xml.GetTag("a:shade")
						shade.AddAttribute("val", "90000")
						lumMod := docx_xml.GetTag("a:lumMod")
						lumMod.AddAttribute("val", "103000")

						schemeClr.AddContentTags(tint, satMod, shade, lumMod)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					{
						gs := docx_xml.GetTag("a:gs")
						gs.AddAttribute("pos", "100000")
						schemeClr := docx_xml.GetTag("a:schemeClr")
						schemeClr.AddAttribute("val", "phClr")

						shade := docx_xml.GetTag("a:shade")
						shade.AddAttribute("val", "63000")
						satMod := docx_xml.GetTag("a:satMod")
						satMod.AddAttribute("val", "120000")

						schemeClr.AddContentTags(shade, satMod)
						gs.AddContentTags(schemeClr)
						gsLst.AddContentTags(gs)
					}
					lin := docx_xml.GetTag("a:lin")
					lin.AddAttribute("ang", "5400000")
					lin.AddAttribute("scaled", "0")

					gradFill.AddContentTags(gsLst, lin)
				}
				bgFillStyleLst.AddContentTags(gradFill)
			}
			fmtScheme.AddContentTags(bgFillStyleLst)
		}

		themeElements.AddContentTags(fmtScheme)
	}
	objectDefaults := docx_xml.GetTag("a:objectDefaults")
	extraClrSchemeLst := docx_xml.GetTag("a:extraClrSchemeLst")
	extLst := docx_xml.GetTag("a:extLst")
	{
		ext := docx_xml.GetTag("a:ext")
		ext.AddAttribute("uri", "{05A4C25C-085E-4340-85A3-A5531E510DB2}")

		themeFamily := docx_xml.GetTag("thm15:themeFamily")
		themeFamily.AddAttribute("xmlns:thm15", "http://schemas.microsoft.com/office/thememl/2012/main")
		themeFamily.AddAttribute("name", "Office Theme")
		themeFamily.AddAttribute("id", "{62F939B6-93AF-4DB8-9C6B-D6C7DFDC589F}")
		themeFamily.AddAttribute("vid", "{4A3C46E8-61CC-4603-A589-7422A47A8E4A}")

		ext.AddContentTags(themeFamily)
		extLst.AddContentTags(ext)
	}

	theme.AddContentTags(themeElements, objectDefaults, extraClrSchemeLst, extLst)
	return theme
}
