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

	err := pdf.AddTTFFont("Noto Sans JP", "./assets/fonts/NotoSansJP-Regular.ttf")
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = pdf.SetFont("Noto Sans JP", "", 16)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// 暫定: fontの設定と行幅が色んな場所で転々と利用されているのがいまいち
	lineHeight := 16 * 1.2

	// 表紙
	pdf.AddPage()
	pdf.UseImportedTemplate(pdf.ImportPage("./assets/pdf/diary_front_cover.pdf", 1, MediaBox), 0, 0, pageReact.W, pageReact.H)

	// 日記の各ページ
	pageTemplateID := pdf.ImportPage("./assets/pdf/diary_page.pdf", 1, MediaBox)
	AddDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/camp.jpg",
		`今日は自然の中での素晴らしい時間を過ごした。朝は鳥のさえずりで目覚め、清涼な空気を吸いながらの朝食は格別だった。
	昼間は木々の間を散策し、奇跡的な景色に感動した。夜には満天の星空の下、仲間との団欒が心地よかった。焚火の炎を見つめながら、幸せな時間を共有し、マシュマロを焼いて笑い合った。自然の中での静寂に包まれ、心が穏やかになった。この経験は忘れられない思い出となった。`, // ChatGPT
		lineHeight,
		pageReact,
	)
	AddDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/camp.jpg",
		`今日は自然の中での素晴らしい時間を過ごした。朝は鳥のさえずりで目覚め、清涼な空気を吸いながらの朝食は格別だった。
		昼間は木々の間を散策し、奇跡的な景色に感動した。夜には満天の星空の下、仲間との団欒が心地よかった。焚火の炎を見つめながら、幸せな時間を共有し、マシュマロを焼いて笑い合った。自然の中での静寂に包まれ、心が穏やかになった。この経験は忘れられない思い出となった。`,
		lineHeight,
		pageReact,
	)
	AddDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/camp.jpg",
		`今日は自然の中での素晴らしい時間を過ごした。朝は鳥のさえずりで目覚め、清涼な空気を吸いながらの朝食は格別だった。
	昼間は木々の間を散策し、奇跡的な景色に感動した。夜には満天の星空の下、仲間との団欒が心地よかった。焚火の炎を見つめながら、幸せな時間を共有し、マシュマロを焼いて笑い合った。自然の中での静寂に包まれ、心が穏やかになった。この経験は忘れられない思い出となった。`,
		lineHeight,
		pageReact,
	)

	// 背表紙
	pdf.AddPage()
	pdf.UseImportedTemplate(pdf.ImportPage("./assets/pdf/diary_back_cover.pdf", 1, MediaBox), 0, 0, pageReact.W, pageReact.H)

	pdf.WritePdf("./output/output.pdf")

	return nil
}

func AddDialyPage(pdf *gopdf.GoPdf, templateID int, imagePath string, text string, lineHeight float64, pageReact gopdf.Rect) error {
	pdf.AddPage()
	pdf.UseImportedTemplate(templateID, 0, 0, pageReact.W, pageReact.H)

	// Rectにリサイズが行われて、比率が変わると引き伸ばされる
	pdf.Image(imagePath, 47.64, 75, &gopdf.Rect{
		W: 500,
		H: 375,
	})

	// 横幅に適用できるように文字を分割、改行できるように
	// TODO: 改行コードの分だけスペースが空く
	lines, err := pdf.SplitText(text, 480)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	for i, line := range lines {
		if err := drawText(pdf, 40, 480+float64(i)*lineHeight, line); err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}

func drawText(pdf *gopdf.GoPdf, x float64, y float64, text string) error {
	pdf.SetXY(x, y)
	// cellを確保してもその範囲内で改行はしてくれない
	return pdf.Cell(nil, text)
}

// 座標検証用: 縦長用紙前提
func drawGrid(pdf *gopdf.GoPdf, page *gopdf.Rect) {
	ww := 10.0
	for i := 1; i <= int(page.H/ww); i++ {
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
