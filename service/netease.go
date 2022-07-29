package service

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"strings"

	"stock.api/utils"
)

func GetStockDataUrl(code string, startDate string, endDate string) string {
	baseUrl := GetStockBaseUrl()
	var bt bytes.Buffer
	bt.WriteString(baseUrl)
	bt.WriteString("code=")
	bt.WriteString(code)
	bt.WriteString("&start=")
	bt.WriteString(startDate)

	bt.WriteString("&end=")
	bt.WriteString(endDate)

	bt.WriteString("&fields=TOPEN%3BHIGH%3BLOW%3BTCLOSE")
	return bt.String()
}

func GetStockBaseUrl() string {
	return "http://quotes.money.163.com/service/chddata.html?"
}

func GetOnlineStockDataAndSaveToMongo(code string, startDate string, endDate string) string {
	
	url := GetStockDataUrl(code, startDate, endDate)
	content := utils.HttpGet(url)
	content = strings.ReplaceAll(content, "None", "0")

	return content
}

type NeteaseDayData struct {
	Date    string `json:"date"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Open    string `json:"open"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Close   string `json:"close"`
	Amount  string `json:"amount"`
	Money   string `json:"money"`
	Raise   string `json:"raise"`
	Percent string `json:"percent"`
}

func ParseNeteaseData(content string) (*[]NeteaseDayData, error) {
	buf := bufio.NewReader(strings.NewReader(content))
	reader := csv.NewReader(bufio.NewReader(buf))
	lines := []NeteaseDayData{}
	skip := 1
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if skip == 1 {
			//skip first line
			skip = 0
			continue
		}
		name, _ := utils.GbkToUtf8([]byte(line[2]))
		code := line[1]
		//code := Code.GetCode(strings.ReplaceAll(line[1], "'", ""))
		lines = append(lines, NeteaseDayData{
			Date:    line[0],
			Code:    code,
			Name:    string(name),
			Open:    line[3],
			High:    line[4],
			Low:     line[5],
			Close:   line[6],
			Amount:  line[7],
			Money:   line[8],
			Raise:   line[9],
			Percent: line[10],
		})
	}
	return &lines, nil
}

//http://quotes.money.163.com/service/chddata.html?code=0600367&start=20210101&end=20210408&fields=TOPEN%3BHIGH%3BLOW%3BTCLOSE%3BVOTURNOVER%3BVATURNOVER%3BCHG%3BPCHG
