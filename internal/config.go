package internal

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
	// 検索条件
	Condition condition `yaml:"condition"`
}
