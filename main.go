package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"AtCoderABCReport/internal"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-yaml/yaml"
)

// スクレイピングの本体
func runScraping(config *internal.Config, outputFile *os.File) error {
	configAtCoderABC := internal.AtCoderABC{
		ContestBaseUrl:     config.ContestBaseUrl,
		ContestPrefix:      config.ContestPrefix,
		ContestCount:       config.ContestCount,
		ContestTask:        config.Condition.ContestTask,
		ContestLanguage:    config.Condition.ContestLanguage,
		ContestStasus:      config.Condition.ContestStasus,
		ContestUser:        config.Condition.ContestUser,
		ContestPage:        1,
		ScrapingWaitMinute: 1,
	}

	// URL生成
	url := configAtCoderABC.MakeUrl(configAtCoderABC.ContestPage)
	internal.OutputStderr("fetching... "+url, true)

	// スクレイピング実行
	doc, err := goquery.NewDocument(url)
	if err != nil {
		internal.OutputStderr(err.Error(), true)
		return err
	}

	//// 情報取得 ////

	// タイトル: strTitle
	docTitle := doc.Find(".contest-title")
	if docTitle == nil {
		internal.OutputStderr(err.Error(), true)
		return err
	}
	strTitle := docTitle.Text()
	internal.OutputStderr("get result of "+strTitle, true)

	// 提出ページ数: numPages
	docItems := doc.Find(".pagination")
	if docItems == nil {
		internal.OutputStderr(err.Error(), true)
		return err
	}
	numPages := 0
	docItems.Find("li a").Each(func(i int, s *goquery.Selection) {
		numPages, _ = strconv.Atoi(s.Text())
	})
	internal.OutputStderr("report num is "+strconv.Itoa(numPages), true)

	count := 0
	for i := configAtCoderABC.ContestPage; i <= numPages; i++ {
		// 最初のページでなければページを取得する
		if i > configAtCoderABC.ContestPage {
			// スクレイピング用に1分ウェイトを入れる
			time.Sleep(time.Minute * time.Duration(configAtCoderABC.ScrapingWaitMinute))

			// URL生成
			url := configAtCoderABC.MakeUrl(i)
			internal.OutputStderr("fetching... "+url, true)

			// スクレイピング実行
			doc, err = goquery.NewDocument(url)
			if err != nil {
				internal.OutputStderr(err.Error(), true)
				break
			}
			internal.OutputStderr("get page "+strconv.Itoa(i)+
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
			reportData.Output(count, outputFile, config.OutputSjis)
		})
	}
	internal.OutputStderr("get "+strconv.Itoa(count)+" lines.", true)
	return nil
}

func parseTable(s *goquery.Selection, c *internal.AtCoderABC) internal.AtCoderABCReportTable {
	reportData := internal.AtCoderABCReportTable{}
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

func main() {
	// コマンドライン引数を取得
	if len(os.Args) == 1 {
		internal.OutputStderr("require config file.", true)
		os.Exit(1)
	}
	configFileName := os.Args[1]

	// 出力ファイル名は設定ファイル名の拡張子をcsvにしたものにする
	outputFileName := strings.Replace(configFileName, "yml", "csv", 1)
	internal.OutputStderr("Output file: "+outputFileName, true)

	// 設定ファイルのチェック
	if _, err := os.Stat(configFileName); err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(2)
	}

	// 設定ファイル(YAML)を読み込み
	configFile, err := os.Open(configFileName)
	if err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(3)
	}
	defer configFile.Close()
	config := internal.Config{}

	// 設定ファイル読み取り
	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(4)
	}
	internal.OutputStderr("Contest Number:"+strconv.Itoa(config.ContestCount), true)
	internal.OutputStderr("Condition.Language:"+config.Condition.ContestLanguage, true)

	// 出力ファイルの作成
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(5)
	}
	defer outputFile.Close()

	// スクレイピング実行
	if err := runScraping(&config, outputFile); err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(6)
	}

	os.Exit(0)
}
