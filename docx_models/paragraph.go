package docx_models

import "docx/docx_format"

type Paragraph struct {
	texts  []*TextModel
	format docx_format.ParagraphFormat

	Numbering int
}

func (p *Paragraph) Format(format docx_format.ParagraphFormat) *Paragraph {
	p.format = format
	return p
}

func (p *Paragraph) AddText(text string) *TextModel {
	newText := TextModel{
		Text:   text,
		format: docx_format.NewTextFormat(),
	}
	p.texts = append(p.texts, &newText)
	return &newText
}

func (p *Paragraph) SetText(pFormat docx_format.ParagraphFormat, tFormat docx_format.TextFormat, text string) *Paragraph {
	p.texts = []*TextModel{{
		Text:   text,
		format: tFormat,
	}}
	p.Format(pFormat)
	return p
}
