package create_docx

import (
	"archive/zip"
	df "docx/create_docx/default_files"
	"docx/docx_xml"
	"io"
	"os"
)

// WriteToFile запись в word документ
func WriteToFile(path string, userFiles map[string][]byte) (err error) {
	var target *os.File
	target, err = os.Create(path)
	if err != nil {
		return
	}
	defer target.Close()

	files := addDefaultFiles(userFiles)

	// Запись в файл
	err = write(target, files)
	return
}
func write(ioWriter io.Writer, files map[string][]byte) (err error) {
	w := zip.NewWriter(ioWriter)
	for filename, content := range files {
		var writer io.Writer

		writer, err = w.Create(filename)
		if err != nil {
			return err
		}

		writer.Write(content)
	}
	w.Close()
	return
}

// addDefaultFiles К переданной карте, добавляет default значения, если их нет в переданной карте
func addDefaultFiles(files map[string][]byte) map[string][]byte {
	xmlProlog := docx_xml.Prolog

	fileNotExist := func(key string) bool {
		_, b := files[key]
		return !b
	}
	move := func(key string, val string) {
		files[key] = []byte(val)
	}

	// root
	if fileNotExist("[Content_Types].xml") {
		move("[Content_Types].xml", xmlProlog+df.GetXmlContentTypes().GenerateXml())
	}
	{ // dir: _rels
		{
			if fileNotExist("_rels/.rels") {
				move("_rels/.rels", xmlProlog+df.GetXmlRels().GenerateXml())
			}
		}
		{ // dir: docProps
			if fileNotExist("docProps/app.xml") {
				move("docProps/app.xml", xmlProlog+df.GetXmlDocPropsApp().GenerateXml())
			}
			if fileNotExist("docProps/core.xml") {
				move("docProps/core.xml", xmlProlog+df.GetXmlDocPropsCore().GenerateXml())
			}
		}
		{ // dir: word
			{ // dir: word/_rels
				if fileNotExist("word/_rels/document.xml.rels") {
					move("word/_rels/document.xml.rels", xmlProlog+df.GetXmlDocumentXmlRels().GenerateXml())
				}
			}
			// dir: word/theme
			{
				if fileNotExist("word/theme/theme1.xml") {
					move("word/theme/theme1.xml", xmlProlog+df.GetXmlWordTheme().GenerateXml())
				}
			}
			if fileNotExist("word/fontTable.xml") {
				move("word/fontTable.xml", xmlProlog+df.GetXmlFontTable().GenerateXml())
			}
			if fileNotExist("word/settings.xml") {
				move("word/settings.xml", xmlProlog+df.GetXmlSettings().GenerateXml())
			}
			if fileNotExist("word/styles.xml") {
				move("word/styles.xml", xmlProlog+df.GetXmlStyles().GenerateXml())
			}
			if fileNotExist("word/webSettings.xml") {
				move("word/webSettings.xml", xmlProlog+df.GetXmlWebSettings().GenerateXml())
			}
		}
	}
	return files
}
