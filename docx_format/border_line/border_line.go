package border_line

import "docx/docx_format/rgb"

type CellBorders struct {
	Top    BorderFormat
	Bottom BorderFormat
	Start  BorderFormat
	End    BorderFormat

	InsideH BorderFormat
	InsideV BorderFormat
	Tl2br   BorderFormat
	Tr2bl   BorderFormat
}

func NewCellBorders() CellBorders {
	external := BorderFormat{
		Color:  rgb.ColorRGB{R: 255, G: 255, B: 255},
		Shadow: false,
		Space:  0,
		Size:   4,
		Value:  BorderLineSingle,
	}
	inside := BorderFormat{
		Color:  rgb.ColorRGB{R: 255, G: 255, B: 255},
		Shadow: false,
		Space:  0,
		Size:   4,
		Value:  BorderLineNil,
	}
	return CellBorders{
		Top:    external,
		Bottom: external,
		Start:  external,
		End:    external,

		InsideH: inside,
		InsideV: inside,
		Tl2br:   inside,
		Tr2bl:   inside,
	}
}
func (b CellBorders) SetExternalBorders(format BorderFormat) CellBorders {
	b.Top = format
	b.Bottom = format
	b.Start = format
	b.End = format
	return b
}

type BorderFormat struct {
	Color  rgb.ColorRGB // Цвет
	Shadow bool         // Тень (для правой и нижней)
	Space  int          // Смещение
	Size   int          // Толщина
	Value  BorderLine   // Вид линии
}

func NewBorderFormat() BorderFormat {
	return BorderFormat{
		Color:  rgb.ColorRGB{R: 255, G: 255, B: 255},
		Shadow: false,
		Space:  0,
		Size:   4,
		Value:  BorderLineSingle,
	}
}
func (b BorderFormat) SetColor(color rgb.ColorRGB) BorderFormat {
	b.Color = color
	return b
}
func (b BorderFormat) SetShadow(shadow bool) BorderFormat {
	b.Shadow = shadow
	return b
}
func (b BorderFormat) SetSpace(space int) BorderFormat {
	b.Space = space
	return b
}
func (b BorderFormat) SetSize(size int) BorderFormat {
	b.Size = size
	return b
}
func (b BorderFormat) SetBorderLine(val BorderLine) BorderFormat {
	b.Value = val
	return b
}

type BorderLine string

const (
	BorderLineSingle         BorderLine = "single"         //- a single line
	BorderLineDashDotStroked            = "dashDotStroked" //- a line with a series of alternating thin and thick strokes
	BorderLineDashed                    = "dashed"         //- a dashed line
	BorderLineDashSmallGap              = "dashSmallGap"   //- a dashed line with small gaps
	BorderLineDotDash                   = "dotDash"        //- a line with alternating dots and dashes
	BorderLineDotDotDash                = "dotDotDash"     //- a line with a repeating dot - dot - dash sequence
	BorderLineDotted                    = "dotted"         //- a dotted line
	BorderLineDouble                    = "double"         //- a double line
	BorderLineDoubleWave                = "doubleWave"     //- a double wavy line
	BorderLineInset                     = "inset"          //- an inset set of lines
	BorderLineNil                       = "nil"            //- no border
	//BorderLineNone                              = "none"                   //- no border
	BorderLineOutset                 = "outset"                 //- an outset set of lines
	BorderLineThick                  = "thick"                  //- a single line
	BorderLineThickThinLargeGap      = "thickThinLargeGap"      //- a thick line contained within a thin line with a large-sized intermediate gap
	BorderLineThickThinMediumGap     = "thickThinMediumGap"     //- a thick line contained within a thin line with a medium-sized intermediate gap
	BorderLineThickThinSmallGap      = "thickThinSmallGap"      // - a thick line contained within a thin line with a small intermediate gap
	BorderLineThinThickLargeGap      = "thinThickLargeGap"      //- a thin line contained within a thick line with a large-sized intermediate gap
	BorderLineThinThickMediumGap     = "thinThickMediumGap"     //- a thick line contained within a thin line with a medium-sized intermediate gap
	BorderLineThinThickSmallGap      = "thinThickSmallGap"      //- a thick line contained within a thin line with a small intermediate gap
	BorderLineThinThickThinLargeGap  = "thinThickThinLargeGap"  //- a thin-thick-thin line with a large gap
	BorderLineThinThickThinMediumGap = "thinThickThinMediumGap" //- a thin-thick-thin line with a medium gap
	BorderLineThinThickThinSmallGap  = "thinThickThinSmallGap"  //- a thin-thick-thin line with a small gap
	BorderLineThreeDEmboss           = "threeDEmboss"           //- a three-staged gradient line, getting darker towards the paragraph
	BorderLineThreeDEngrave          = "threeDEngrave"          //- a three-staged gradient like, getting darker away from the paragraph
	BorderLineTriple                 = "triple"                 //- a triple line
	BorderLineWave                   = "wave"                   //- a wavy line
)
