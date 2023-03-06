package docx_models

import (
	"docx/docx_format"
	"errors"
	"fmt"
)

type Table struct {
	rows    []*TableRow
	columns []*TableColumn

	// Format
	format docx_format.TableFormat
}
type TableColumn struct {
	cells []*TableCell

	// Format
	format docx_format.TableColumnFormat
}
type TableRow struct {
	cells []*TableCell

	// Format
	format docx_format.TableRowFormat
}
type TableCell struct {
	content []elementModel

	indexX, indexY int // X, Y индекс ячейки в строке, столице (по горизонтали, по вертикали)
	merge          *mergeCell

	// Format
	format docx_format.TableCellFormat
}

type mergeCell struct {
	general bool       // Левый верхний элемент объединения
	fixed   *TableCell // (если general = false) - закреплена, за. (если general = true) - последний элемент объединения
}

/*
	Table
*/

func (t *Table) Format(format docx_format.TableFormat) *Table {
	t.format = format
	return t
}

// newTable создаёт таблицу, нужно указать размерность таблицы (кол-во рядов, кол-во колонок)
func newTable(y, x int) *Table {
	var t Table
	rows := make([]TableRow, y)    // Массив (ссылок) рядов нужной размерности
	cols := make([]TableColumn, x) // Массив (ссылок) колонок нужной размерности

	// Инициализируется интерфейс для каждой ячейки
	for rowIndex := 0; rowIndex < y; rowIndex++ {
		for colIndex := 0; colIndex < x; colIndex++ {
			cell := TableCell{indexX: colIndex, indexY: rowIndex}
			cols[colIndex].cells = append(cols[colIndex].cells, &cell) // Ссылка на ячейку добавляется в список колонок
			rows[rowIndex].cells = append(rows[rowIndex].cells, &cell) // Ссылка на ячейку добавляется в список рядов
		}
	}
	// Ссылки добавляются к таблице
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		t.rows = append(t.rows, &rows[rowIndex])
	}

	for colIndex := 0; colIndex < len(cols); colIndex++ {
		t.columns = append(t.columns, &cols[colIndex])
	}
	
	return &t
}

func (t *Table) Merge(c1 *TableCell, c2 *TableCell) error {
	// Не допустимый диапазон объединения
	if c1.indexX > c2.indexX || c1.indexY > c2.indexY {
		return errors.New("Не допустимый диапазон объединения ")
	}
	for ix := c1.indexX; ix <= c2.indexX; ix++ {
		for iy := c1.indexY; iy <= c2.indexY; iy++ {
			cell := t.Cell(iy, ix)
			cell.merge = &mergeCell{general: false, fixed: c1}
		}
	}

	c1.merge = &mergeCell{general: true, fixed: c2}
	return nil
}
func (c TableCell) GetMerge() (bool, bool) {
	if c.merge != nil {
		return true, c.merge.general
	}
	return false, false
}

// AddRow Возвращает ссылку на новую линию
func (t *Table) AddRow() *TableRow {

	var row TableRow
	for i := 0; i < len(t.columns); i++ {
		cell := TableCell{indexY: len(t.rows), indexX: i}
		t.columns[i].cells = append(t.columns[i].cells, &cell) // Ссылка на ячейку добавляется к колонке
		row.cells = append(row.cells, &cell)                   // Ссылка на ячейку добавляется к ряду
	}

	t.rows = append(t.rows, &row)
	return &row
}

// AddColumn Возвращает ссылку на новую колонку
func (t *Table) AddColumn() *TableColumn {

	var col TableColumn
	for i := 0; i < len(t.rows); i++ {
		cell := TableCell{indexY: i, indexX: len(t.columns)}
		t.rows[i].cells = append(t.rows[i].cells, &cell) // Ссылка на ячейку добавляется к ряду
		col.cells = append(col.cells, &cell)             // Ссылка на ячейку добавляется к колонке
	}

	t.columns = append(t.columns, &col)
	return &col
}

// Rows Возвращает массив ссылок на линии
func (t Table) Rows() []*TableRow { // []*TableRow
	return t.rows
}

// Row Возвращает ссылку на линию
func (t Table) Row(y int) (row *TableRow) {
	if len(t.rows) < y+1 {
		fmt.Printf("количество рядов в таблице %v, меньше чем запрашиваемый ряд %v + 1", len(t.rows), y)
	} else if y < 0 {
		fmt.Printf("x, индекс - не может быть < 0")
	} else {
		row = t.rows[y]
	}
	return
}

// Cells Возвращает массив ссылок на ячейки, принимает индекс строки
func (t Table) Cells(x int) (row []*TableCell) { // *TableCell
	if len(t.rows) < x+1 {
		fmt.Printf("количество рядов в таблице %v, меньше чем индекс запрашиваемого ряда %v + 1", len(t.rows), x)
	} else if x < 0 {
		fmt.Printf("x, индекс - не может быть < 0")
	} else {
		row = t.rows[x].cells
	}
	return
}

// Cell Возвращает ссылку на ячейку
func (t Table) Cell(y, x int) (cell *TableCell) { // *TableCell
	if len(t.rows) < y+1 {
		fmt.Printf("количество рядов в таблице %v, меньше чем (x) индекс запрашиваемого ряда %v + 1", len(t.rows), y)
	} else if y < 0 {
		fmt.Printf("x, индекс - не может быть < 0")
	} else if len(t.columns) < x+1 {
		fmt.Printf("количество колонок в таблице %v, меньше чем (y) индекс запрашиваемой колонки %v + 1", len(t.columns), x)
	} else if x < 0 {
		fmt.Printf("y, индекс - не может быть < 0")
	} else {
		cell = t.rows[y].cells[x]
	}
	return
}

// Columns Возвращает список колонок в таблице
func (t Table) Columns() []*TableColumn {
	return t.columns
}

// Column Возвращает колонку в таблице по индексу
func (t Table) Column(y int) (col *TableColumn) {
	if len(t.columns) < y+1 {
		fmt.Printf("количество колонок в таблице %v, меньше чем (y) индекс запрашиваемой колонки %v + 1", len(t.columns), y)
	} else if y < 0 {
		fmt.Printf("y, индекс - не может быть < 0")
	} else {
		col = t.columns[y]
	}
	return
}

/*
	TableRow
*/

func (r *TableRow) Format(format docx_format.TableRowFormat) *TableRow {
	r.format = format
	return r
}

// Cells Возвращает массив ссылок на ячейки
func (r TableRow) Cells() []*TableCell { // []*TableCell
	return r.cells
}

// Cell Возвращает ссылку на ячейку
func (r TableRow) Cell(x int) (cell *TableCell) { // *TableCell
	if len(r.cells) < x+1 {
		fmt.Printf("Количество колонок в таблице %v, меньше чем индекс запрашиваемой колонки %v + 1", len(r.cells), x)
	} else if x < 0 {
		fmt.Printf("y, индекс - не может быть < 0")
	} else {
		cell = r.cells[x]
	}
	return
}

/*
	TableColumn
*/

func (r *TableColumn) Format(format docx_format.TableColumnFormat) *TableColumn {
	r.format = format
	return r
}

// Cells Возвращает массив ссылок на ячейки
func (r TableColumn) Cells() []*TableCell { // []*TableCell
	return r.cells
}

// Cell Возвращает ссылку на ячейку
func (r TableColumn) Cell(x int) (cell *TableCell) { // *TableCell
	if len(r.cells) < x+1 {
		fmt.Printf("количество рядов в таблице %v, меньше чем индекс запрашиваемого ряда %v + 1", len(r.cells), x)
	} else if x < 0 {
		fmt.Printf("x, индекс - не может быть < 0")
	} else {
		cell = r.cells[x]
	}
	return
}

/*
	TableCell
*/

func (c *TableCell) Format(format docx_format.TableCellFormat) *TableCell {
	c.format = format
	return c
}

// SetText создаёт простой параграф и простой текст
func (c *TableCell) SetText(pFormat docx_format.ParagraphFormat, tFormat docx_format.TextFormat, text string) *TableCell {
	// Очищает содержимое ячейки
	c.content = []elementModel{}
	// Добавляет параграф, задаёт текст
	p := c.AddParagraph().Format(pFormat)
	p.AddText(text).Format(tFormat)
	return c
}

// AddParagraph Добавляет в ячейку ссылку на параграф, возвращает ссылку на параграф
func (c *TableCell) AddParagraph() *Paragraph {
	paragraph := &Paragraph{}
	c.content = append(c.content, elementModel(paragraph))
	return paragraph
}
func (c *TableCell) AddTable(x, y int) *Table {
	table := newTable(x, y) // создаёт таблицу
	c.content = append(c.content, elementModel(table))
	return table
}

func (c *TableCell) GetText() string {
	cellText := ""
	for _, model := range c.content {
		switch model.(type) {
		case *Paragraph:
			item := model.(*Paragraph)
			for _, text := range item.texts {
				cellText += text.Text + " "
			}
		}
	}
	return cellText
}
