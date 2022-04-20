package internal

import (
	"os"

	"gopkg.in/yaml.v2"
)

//コンテスト情報
// https://atcoder.jp/contests/abc247/submissions?f.Task=&f.LanguageName=Dart&f.Status=AC&f.User=
// https://atcoder.jp/contests/abc247/submissions?f.LanguageName=Go&f.Status=AC&f.Task=&f.User=&page=2

type condition struct {
	// 問題
	ContestTask string `yaml:"contestTask"`
	// 言語
	ContestLanguage string `yaml:"contestLanguage"`
	// ステータス
	ContestStasus string `yaml:"contestStasus"`
	// ユーザ
	ContestUser string `yaml:"contestUser"`
}

type Config struct {
	// AtCoderのURL（原則変更しない）
	ContestBaseUrl string `yaml:"contestBaseUrl"`
	// ABCのURL（原則変更しない）
	ContestPrefix string `yaml:"contestPrefix"`
	// コンテストの回数
	ContestCount int `yaml:"contestCount"`
	// 出力エンコーデイング
	OutputSjis bool `yaml:"outputSjis"`
	// 開始ページ
	ContestStartPage int `yaml:"contestStartPage"`
	// 終了ページ
	ContestEndPage int `yaml:"contestEndPage"`
	// 1ページあたりのレポート数
	ReportPerPage int `yaml:"reportPerPage"`
	// 検索条件
	Condition condition `yaml:"condition"`
}

func (c *Config) ReadConfigFile(fileName string) error {
	// 設定ファイル(YAML)を読み込み
	configFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// 設定ファイル読み取り
	if err := yaml.NewDecoder(configFile).Decode(c); err != nil {
		return err
	}

	return nil
}
