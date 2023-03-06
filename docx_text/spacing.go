package docx_text

type TextSpacing struct {
	Left   int // Отступ слева
	Right  int // Отступ справа
	Before int // Интервал перед
	After  int // Интервал после
}

func (s TextSpacing) isDefaultLeft() bool {
	return s.Left == 0
}
func (s TextSpacing) isDefaultRight() bool {
	return s.Right == 0
}
func (s TextSpacing) isDefaultBefore() bool {
	return s.Before == 0
}
func (s TextSpacing) isDefaultAfter() bool {
	return s.After == 8
}
