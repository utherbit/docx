package docx_models

import (
	"docx/docx_format"
	"docx/docx_xml"
)

type Document struct {
	elements []elementModel
}
type elementModel interface {
	generateTag() []docx_xml.Tag
}

func (d *Document) AddParagraph() *Paragraph {
	paragraph := &Paragraph{format: docx_format.ParagraphFormat{}}
	d.elements = append(d.elements, paragraph)
	return paragraph
}

// AddTable добавляет созданную таблицу в документ, возвращает ссылку на таблицу
func (d *Document) AddTable(y, x int) *Table {
	table := newTable(y, x)
	d.elements = append(d.elements, table)
	return table
}
