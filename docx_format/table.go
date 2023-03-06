package docx_format

import (
	"docx/docx_format/border_line"
	"docx/docx_format/rgb"
)

/*
	Formatting
	Table
*/

type TableFormat struct {
	Alignment TableAlignment
	AutoFit   bool // Фиксированные размеры таблицы
}

func NewTableFormat() TableFormat {
	return TableFormat{}
}
func (f TableFormat) SetAlignment(align TableAlignment) TableFormat {
	f.Alignment = align
	return f
}
func (f TableFormat) SetAutoFit(autoFit bool) TableFormat {
	f.AutoFit = autoFit
	return f
}

type TableAlignment int

const (
	TableAlignmentLeft   TableAlignment = iota // Слева
	TableAlignmentCenter                       // По центру
	TableAlignmentRight                        // Справа
)

/*
	Formatting
	TableColumn
*/

type TableColumnFormat struct {
	Width int
}

func NewTableColumnFormat() TableColumnFormat {
	return TableColumnFormat{}
}
func (f TableColumnFormat) SetWidth(width int) TableColumnFormat {
	f.Width = width
	return f
}

/*
	Formatting
	TableRow
*/

type TableRowFormat struct {
	Height     int
	HeightRule HeightRule
}

func NewTableRowFormat() TableRowFormat {
	return TableRowFormat{}
}
func (f TableRowFormat) SetHeight(height int) TableRowFormat {
	f.Height = height
	return f
}
func (f TableRowFormat) SetHeightRule(rule HeightRule) TableRowFormat {
	f.HeightRule = rule
	return f
}

type HeightRule int

func (r HeightRule) String() string {

	switch r {
	default:
		return "auto"
	case HeightRuleAuto:
		return "auto"
	case HeightRuleExactly:
		return "exact"
	case HeightRuleAtLeast:
		return "atLeast"
	}
}

const (
	HeightRuleAuto HeightRule = iota
	HeightRuleAtLeast
	HeightRuleExactly
)

/*
	Formatting
	TableCell
*/

type TableCellFormat struct {
	VerticalAlignment CellVerticalAlignment // todo Вертикальное выравнивание
	Width             int                   // Ширина

	Shadow rgb.ColorRGB            // Фон ячейки
	Border border_line.CellBorders // Стили границ
}

func NewTableCellFormat() TableCellFormat {
	return TableCellFormat{}
}

func (c TableCellFormat) SetVerticalAlignment(alignment CellVerticalAlignment) TableCellFormat {
	c.VerticalAlignment = alignment
	return c
}
func (c TableCellFormat) SetWidth(width int) TableCellFormat {
	c.Width = width
	return c
}
func (c TableCellFormat) SetShadow(color rgb.ColorRGB) TableCellFormat {
	c.Shadow = color
	return c
}
func (c TableCellFormat) SetBorders(borders border_line.CellBorders) TableCellFormat {
	c.Border = borders
	return c
}

type CellVerticalAlignment int

const (
	CellVerticalAlignmentTop TableAlignment = iota // Слева
	CellVerticalAlignmentCenter
	CellVerticalAlignmentBottom
	CellVerticalAlignmentBoth
)
