package dbops

import (
	"github.com/signintech/gopdf"
	"strconv"
	"go-test/pdf/defs"
	"go-test/pdf/comm"
	"log"
)

var (
	downloadData []*defs.MemberInfo // 数据
	curX         float64
	curY         float64
	k            int
	page         int
)

func getSliceByString(str string) []string {
	var s []string
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		s = append(s, string(r[i]))
	}
	return s
}

func getStringFormSliceString(strSlice []string) string {
	var str string
	for i := 0; i < len(strSlice); i++ {
		str += strSlice[i]
	}
	return str
}

func WritePdf(filePath string, data []*defs.MemberInfo) error {
	// 创建路径
	err = comm.CreateFilePath(filePath)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	pdf := gopdf.GoPdf{}
	defer pdf.Close()
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	pdf.AddPage()
	page = page + 1 //

	err := pdf.AddTTFFont("msyh", "./pdf/ttf/MicrosoftYaqiHeiLight-2.ttf") // 添加字体文件
	if err != nil {
		return err
	}
	err = pdf.SetFont("msyh", "", 14) // 设置字体和大小
	if err != nil {
		return err
	}

	pdf.SetLineWidth(0.5)     //  设置线宽度
	pdf.SetLineType("normal") // 设置线类型

	pdf.SetX(220)
	pdf.SetY(40)
	pdf.Text("斗罗大陆史莱克七怪人员名单")

	pdf.Line(30, 60, 565.28, 60)      // 表单顶部
	pdf.Line(30, 60, 30, 815)         // 表单左1
	pdf.Line(70, 60, 70, 815)         // 表单左2
	pdf.Line(140, 60, 140, 815)       // 表单左3
	pdf.Line(210, 60, 210, 815)       // 表单左4
	pdf.Line(280, 60, 280, 815)       // 表单左5
	pdf.Line(320, 60, 320, 815)       // 表单左6
	pdf.Line(565.28, 60, 565.28, 815) // 表单右1

	pdf.Line(30, 815, 565.28, 815) // 表单底部

	for i := 0; i <= 12; i++ {
		pdf.Line(30, 35+60*float64(i+1), 565.28, 35+60*float64(i+1)) // 生成每行线
	}

	// 表单头部
	pdf.SetX(36)
	pdf.SetY(70)
	pdf.Cell(nil, "序号")

	pdf.SetX(90)
	pdf.SetY(70)
	pdf.Cell(nil, "封号")

	pdf.SetX(160)
	pdf.SetY(70)
	pdf.Cell(nil, "外号")

	pdf.SetX(230)
	pdf.SetY(70)
	pdf.Cell(nil, "姓名")

	pdf.SetX(285)
	pdf.SetY(70)
	pdf.Cell(nil, "性别")
	//
	pdf.SetX(420)
	pdf.SetY(70)
	pdf.Cell(nil, "签名")

	// 主体
	for k, val := range data {
		pdf.SetY(120 + float64((k)*60))

		pdf.SetX(45)
		pdf.Cell(nil, strconv.Itoa(k+1))

		pdf.SetX(75)
		pdf.Cell(nil, val.Username)

		pdf.SetX(145)
		pdf.Cell(nil, val.Nickname)

		pdf.SetX(225)
		pdf.Cell(nil, val.Realname)

		pdf.SetX(295)
		pdf.Cell(nil, val.Sex)

		tipSlcie := getSliceByString(val.Signature)
		tiplen := len(tipSlcie)

		pdf.SetX(330)
		pdf.SetY(110 + float64((k)*60))
		pdf.Cell(nil, getStringFormSliceString(tipSlcie[0:16]))

		if (tiplen >= 16) {
			pdf.SetX(330)
			pdf.SetY(130 + float64((k)*60))
			pdf.Cell(nil, getStringFormSliceString(tipSlcie[16:32]))
		}

	}
	err = pdf.WritePdf(filePath)
	return err
}
