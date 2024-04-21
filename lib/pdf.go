package lib

import (
	"fmt"
	"time"

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

// 適当にとりあえず
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
	today := time.Now()

	// 表紙
	pdf.AddPage()
	pdf.UseImportedTemplate(pdf.ImportPage("./assets/pdf/diary_front_cover.pdf", 1, MediaBox), 0, 0, pageReact.W, pageReact.H)

	// 日記の各ページ
	pageTemplateID := pdf.ImportPage("./assets/pdf/diary_page.pdf", 1, MediaBox)
	addDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/camp.jpg",
		`今日は自然の中での素晴らしい時間を過ごした。朝は鳥のさえずりで目覚め、清涼な空気を吸いながらの朝食は格別だった。
昼間は木々の間を散策し、奇跡的な景色に感動した。夜には満天の星空の下、仲間との団欒が心地よかった。焚火の炎を見つめながら、幸せな時間を共有し、マシュマロを焼いて笑い合った。自然の中での静寂に包まれ、心が穏やかになった。この経験は忘れられない思い出となった。`, // ChatGPT,
		today,
		lineHeight,
		pageReact,
	)
	addDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/park.jpg",
		`今日は天気が良かったので、息子の太郎と一緒に公園に散歩に行きました。公園に着くと、たくさんの人々がいて、子どもたちが元気に遊んでいました。

太郎はまずブランコに乗りたいと言いましたので、一緒にブランコに乗りました。風が心地よく、太郎の笑顔がとても可愛かったです。その後、滑り台やジャングルジムで遊び、ボールを投げたり追いかけっこをして楽しい時間を過ごしました。

公園を一周すると、花壇にたくさんの花が咲いていました。太郎は花に興味津々で、一緒に色とりどりの花を見て回りました。その中で彼が一番気に入ったのは、赤いバラでした。

帰り道、太郎は公園の思い出を語りながら、手を繋いで歩いていました。`,
		today.AddDate(0, 0, 1),
		lineHeight,
		pageReact,
	)
	addDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/aquarium.jpg",
		`今日は晴れていて、息子の太郎と一緒に水族館に行ってきました。太郎は水族館が大好きで、前から行きたがっていたので、とても楽しみにしていました。

水族館に着くと、太郎はワクワクして入り口で手を振りました。館内に入ると、まずはペンギンやイルカなどの展示を見て回りました。太郎は色とりどりの魚に興味津々で、ガラス越しにじっと観察していました。

特に、巨大なサメの水槽には太郎も私も圧倒されました。彼はサメの動きを追い、大きな口を開けたときには驚いた顔をしていました。

水族館の中庭では、イルカショーを見ることができました。太郎はイルカのジャンプや芸に大喜びし、拍手を送りました。

帰り道、太郎は水族館で見た魚やイルカの話を熱心にしました。`,
		today.AddDate(0, 0, 2),
		lineHeight,
		pageReact,
	)
	addDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/sky.jpg",
		`今日は晴れた夜空を見に、息子の太郎と一緒に公園に行ってきました。夜の公園は静かで、星が輝いていました。

公園に着くと、太郎は興奮気味に空を見上げ、「星がいっぱいだね！」と喜んでいました。私たちは芝生に座り、夜空の美しさを満喫しました。

夜空には数えきれないほどの星が輝いていました。太郎は星座を見つけるのに一生懸命で、指を差しながら「あれはオリオン座だ！」と言いました。

月も明るく輝いており、その光が公園を照らしていました。太郎は月に向かって手を伸ばし、「月に行きたいな」とつぶやきました。

しばらくして、遠くの方で流れ星が流れていくのを見つけました。太郎は一生懸命に願い事をしました。`,
		today.AddDate(0, 0, 3),
		lineHeight,
		pageReact,
	)
	addDialyPage(
		&pdf,
		pageTemplateID,
		"./assets/photo/carwindow.jpg",
		`今日は天気が良かったので、息子の太郎と一緒に電車に乗ってお出かけしました。電車に乗るのは太郎にとって初めての経験で、とても楽しみにしていました。

駅に着くと、太郎は興奮気味に改札を通り、ホームで電車を待ちました。電車が到着すると、太郎は大きな声で「電車だ！」と喜んで手を振りました。

電車に乗っている間、太郎は窓の外の景色を興味深そうに眺めていました。通り過ぎる田園風景や建物に興味津々で、何度も「見て！」「あれは何？」と質問してきました。

目的地に到着すると、太郎はワクワクしながら電車から降り、周りを探検しました。公園や商店街を散策し、地元のお店で美味しいおやつを食べました。`,
		today.AddDate(0, 0, 4),
		lineHeight,
		pageReact,
	)

	// 背表紙
	pdf.AddPage()
	pdf.UseImportedTemplate(pdf.ImportPage("./assets/pdf/diary_back_cover.pdf", 1, MediaBox), 0, 0, pageReact.W, pageReact.H)

	pdf.WritePdf("./output/output.pdf")

	return nil
}

func addDialyPage(pdf *gopdf.GoPdf, templateID int, imagePath string, text string, datetime time.Time, lineHeight float64, pageReact gopdf.Rect) error {
	pdf.AddPage()
	pdf.UseImportedTemplate(templateID, 0, 0, pageReact.W, pageReact.H)

	// 日付
	formattedDate := datetime.Format("2006/01/02 15:04")
	drawText(pdf, 420, 50, formattedDate)

	// Rectにリサイズが行われて、比率が変わると引き伸ばされる
	pdf.Image(imagePath, 47.64, 75, &gopdf.Rect{
		W: 500,
		H: 375,
	})

	// 本文: 横幅に適用できるように文字を分割、改行
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
