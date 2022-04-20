package main

import (
	"os"
	"strconv"
	"strings"

	"AtCoderABCReport/internal"
)

func main() {
	// コマンドライン引数を取得
	if len(os.Args) == 1 {
		internal.OutputStderr("require config file.", true)
		os.Exit(1)
	}
	configFileName := os.Args[1]

	// 設定ファイルのチェック
	if _, err := os.Stat(configFileName); err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(2)
	}

	// 出力ファイル名は設定ファイル名の拡張子をcsvにしたものにする
	outputFileName := strings.Replace(configFileName, "yml", "csv", 1)
	internal.OutputStderr("Output file: "+outputFileName, true)

	// 設定ファイル読み込み
	config := internal.Config{}
	if err := config.ReadConfigFile(configFileName); err != nil {
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
	if err := internal.RunScraping(&config, outputFile); err != nil {
		internal.OutputStderr(err.Error(), true)
		os.Exit(6)
	}

	os.Exit(0)
}
