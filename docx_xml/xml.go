package docx_xml

const Prolog = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`

type ContentType int

const (
	TagContentTypeEmpty ContentType = iota // Без содержания 	<t/>
	TagContentTypeText                     // С текстом 		<t>текст внутри тега</t>
	TagContentTypeTags                     // С тегами 			<t><tag_inside/></t>
)

type Tag struct {
	TagName     string      // Название 			<document/>
	Attributes  []Attribute // Перечень атрибутов 	<t version="1.0"/>
	ContentText string      // Содержание 			<t>текст внутри тега</t>
	ContentTags []Tag       // Содержание 			<t><tag_inside/></t>
	ContentType ContentType // Тип содержания
}

type Attribute struct {
	Name  string // version=
	Value string // "1.0"
}

func GetTag(name string) Tag {
	return Tag{TagName: name, ContentType: TagContentTypeEmpty}
}
func (d *Tag) AddAttribute(name, value string) {
	d.Attributes = append(d.Attributes, Attribute{Name: name, Value: value})
}
func (d *Tag) AddContentTags(tag ...Tag) {
	d.ContentTags = append(d.ContentTags, tag...)
	d.ContentType = TagContentTypeTags
}
func (d *Tag) SetContentText(text string) {
	d.ContentText = text
	d.ContentType = TagContentTypeText
}

func (d Tag) GenerateXml() (xml string) {
	xml = "<" + d.TagName

	if len(d.Attributes) != 0 {
		for _, v := range d.Attributes {
			xml += " " + v.Name + "=" + "\"" + v.Value + "\""
		}
	}

	switch d.ContentType {
	case TagContentTypeText:
		xml += ">"
		xml += d.ContentText
		xml += "</" + d.TagName + ">"
	case TagContentTypeTags:
		xml += ">"
		for _, tag := range d.ContentTags {
			xml += tag.GenerateXml()
		}
		xml += "</" + d.TagName + ">"
	default:
		xml += "/>"
	}

	return
}
