package lib

import (
	"fmt"

	"github.com/raru-ex/pdf-sample/core/model"
	"github.com/signintech/gopdf"
)

const (
	MediaBox = "/MediaBox"
	CropBox  = "/CropBox"
	BleedBox = "/BleedBox"
	TrimdBox = "/TrimBox"
	ArtBox   = "/ArtBox"
)

// TODO: 適当にとりあえず
func ExportDiary() error {
	pdf := gopdf.GoPdf{}
	pazeSize := model.NewA4()
	pageReact := gopdf.Rect{
		W: pazeSize.Width(),
		H: pazeSize.Height(),
	}

	pdf.Start(gopdf.Config{
		PageSize: pageReact,
	})

	pdf.AddPage()

	err := pdf.AddTTFFont("Noto Sans JP", "./assets/fonts/NotoSansJP-Regular.ttf")
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = pdf.SetFont("Noto Sans JP", "", 16)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	tpl := pdf.ImportPage("./assets/pdf/diary_template.pdf", 1, MediaBox)
	pdf.UseImportedTemplate(tpl, 0, 0, pageReact.W, pageReact.H)

	drawGrid(&pdf, &pageReact)

	// Rectにリサイズが行われて、比率が変わると引き伸ばされる
	pdf.Image("./assets/photo/camp.jpg", 50, 50, &gopdf.Rect{
		W: 400,
		H: 300,
	})

	pdf.WritePdf("./output/output.pdf")

	return nil
}

// 座標検証用
func drawGrid(pdf *gopdf.GoPdf, page *gopdf.Rect) {
	ww := 10.0
	for i := 1; i < int(page.W/ww); i++ {
		if i%10 == 0 {
			pdf.SetLineWidth(0.8)
			pdf.SetStrokeColor(50, 50, 100)
		} else {
			pdf.SetLineWidth(0.3)
			pdf.SetStrokeColor(100, 100, 130)
		}
		x, y := float64(i)*ww, float64(i)*ww
		pdf.Line(x, 0, x, page.H)
		pdf.Line(0, y, page.W, y)
	}
}
