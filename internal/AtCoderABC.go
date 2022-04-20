package internal

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// コンテスト情報
type AtCoderABC struct {
	ContestBaseUrl, ContestPrefix                                 string
	ContestTask, ContestLanguage, ContestStasus, ContestUser      string
	ContestCount, ContestStartPage, ContestEndPage, ReportPerPage int
	ScrapingWaitMinute                                            int
}

// 取得先ULの生成
func (a AtCoderABC) MakeUrl(page int) string {
	return a.ContestBaseUrl +
		a.ContestPrefix +
		strconv.Itoa(a.ContestCount) + "/" +
		"submissions" +
		"?f.Task=" + a.ContestTask +
		"&f.LanguageName=" + a.ContestLanguage +
		"&f.Status=" + a.ContestStasus +
		"&f.User=" + a.ContestUser +
		"&page=" + strconv.Itoa(page)
}

// スクレイピングの本体
func RunScraping(config *Config, outputFile *os.File) error {
	configAtCoderABC := AtCoderABC{
		ContestBaseUrl:     config.ContestBaseUrl,
		ContestPrefix:      config.ContestPrefix,
		ContestCount:       config.ContestCount,
		ContestTask:        config.Condition.ContestTask,
		ContestLanguage:    config.Condition.ContestLanguage,
		ContestStasus:      config.Condition.ContestStasus,
		ContestUser:        config.Condition.ContestUser,
		ContestStartPage:   config.ContestStartPage,
		ContestEndPage:     config.ContestEndPage,
		ReportPerPage:      config.ReportPerPage,
		ScrapingWaitMinute: 1,
	}

	// URL生成
	url := configAtCoderABC.MakeUrl(configAtCoderABC.ContestStartPage)
	OutputStderr("fetching... "+url, true)

	// スクレイピング実行
	doc, err := goquery.NewDocument(url)
	if err != nil {
		OutputStderr(err.Error(), true)
		return err
	}

	//// 情報取得 ////

	// タイトル: strTitle
	docTitle := doc.Find(".contest-title")
	if docTitle == nil {
		OutputStderr(err.Error(), true)
		return err
	}
	strTitle := docTitle.Text()
	OutputStderr("get result of "+strTitle, true)

	// 提出ページ数: numPages
	docItems := doc.Find(".pagination")
	if docItems == nil {
		OutputStderr(err.Error(), true)
		return err
	}
	numPages := 0
	docItems.Find("li a").Each(func(i int, s *goquery.Selection) {
		numPages, _ = strconv.Atoi(s.Text())
	})
	if numPages > config.ContestEndPage {
		numPages = config.ContestEndPage
	}
	if config.ContestEndPage < config.ContestStartPage {
		numPages = config.ContestStartPage
	}
	OutputStderr("report num is "+strconv.Itoa(numPages), true)

	countStart := (config.ContestStartPage - 1) * config.ReportPerPage
	count := 0
	for i := configAtCoderABC.ContestStartPage; i <= numPages; i++ {
		// 最初のページでなければページを取得する
		if i > configAtCoderABC.ContestStartPage {
			// スクレイピング用に1分ウェイトを入れる
			OutputStderr("Waiting...", true)
			time.Sleep(time.Minute * time.Duration(configAtCoderABC.ScrapingWaitMinute))

			// URL生成
			url := configAtCoderABC.MakeUrl(i)
			OutputStderr("fetching... "+url, true)

			// スクレイピング実行
			doc, err = goquery.NewDocument(url)
			if err != nil {
				OutputStderr(err.Error(), true)
				break
			}
			OutputStderr("get page "+strconv.Itoa(i)+
				"/"+strconv.Itoa(numPages), true)
		}

		// 提出結果リストの取得
		doc.Find("div .table-responsive TBODY > TR").Each(func(i int, s *goquery.Selection) {
			reportData := parseTable(s, &configAtCoderABC)
			// ヘッダ出力
			if count == 0 {
				reportData.OutputHeader(outputFile, config.OutputSjis)
			}
			// データ出力
			count += 1
			reportData.Output(countStart+count, outputFile, config.OutputSjis)
		})
	}
	OutputStderr("get "+strconv.Itoa(count)+" lines.", true)
	return nil
}

func parseTable(s *goquery.Selection, c *AtCoderABC) AtCoderABCReportTable {
	reportData := AtCoderABCReportTable{}
	s.Find("TD").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			reportData.RepoteOn = s.Text()
		case 1:
			reportData.Task = s.Text()
		case 2:
			reportData.User = s.Text()
		case 3:
			reportData.Language = s.Text()
		case 4:
			reportData.Score, _ = strconv.Atoi(s.Text())
		case 5:
			strCodeLength := strings.Replace(s.Text(), " Byte", "", 1)
			reportData.CodeLength, _ = strconv.Atoi(strCodeLength)
		case 6:
			reportData.Status = s.Text()
		case 7:
			strCodeTime := strings.Replace(s.Text(), " ms", "", 1)
			reportData.CodeTime, _ = strconv.Atoi(strCodeTime)
		case 8:
			strCodeMemory := strings.Replace(s.Text(), " KB", "", 1)
			reportData.CodeMemory, _ = strconv.Atoi(strCodeMemory)
		case 9:
			strUrl, _ := s.Find("A").Attr("href")
			reportData.DetailUrl = c.ContestBaseUrl + strUrl
		}
	})
	return reportData
}
