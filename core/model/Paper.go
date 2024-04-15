package model

type PaperSize struct {
	width  float64
	height float64
}

func NewA4() PaperSize {
	return PaperSize{
		width:  595.28,
		height: 841.89,
	}
}

func (p PaperSize) Width() float64  { return p.width }
func (p PaperSize) Height() float64 { return p.height }
