package internal

import (
	"os"
	"strconv"
)

// 提出結果テーブル
type AtCoderABCReportTable struct {
	RepoteOn, Task, User, Language, Status, DetailUrl string
	Score, CodeLength, CodeTime, CodeMemory           int
}

// 提出結果テーブル1行出力
func (t AtCoderABCReportTable) Output(count int, outputFile *os.File, sjis bool) {
	OutputFile(outputFile, strconv.Itoa(count)+",", false, sjis)
	OutputFile(outputFile, t.RepoteOn+",", false, sjis)
	OutputFile(outputFile, t.Task+",", false, sjis)
	OutputFile(outputFile, t.User+",", false, sjis)
	OutputFile(outputFile, t.Language+",", false, sjis)
	OutputFile(outputFile, strconv.Itoa(t.Score)+",", false, sjis)
	OutputFile(outputFile, strconv.Itoa(t.CodeLength)+",", false, sjis)
	OutputFile(outputFile, t.Status+",", false, sjis)
	OutputFile(outputFile, strconv.Itoa(t.CodeTime)+",", false, sjis)
	OutputFile(outputFile, strconv.Itoa(t.CodeMemory)+",", false, sjis)
	OutputFile(outputFile, t.DetailUrl, true, sjis)
}

// ヘッダ出力
func (t AtCoderABCReportTable) OutputHeader(outputFile *os.File, sjis bool) {
	OutputFile(outputFile, "#,", false, sjis)
	OutputFile(outputFile, "提出日時,", false, sjis)
	OutputFile(outputFile, "問題,", false, sjis)
	OutputFile(outputFile, "ユーザ,", false, sjis)
	OutputFile(outputFile, "言語,", false, sjis)
	OutputFile(outputFile, "得点,", false, sjis)
	OutputFile(outputFile, "コード長,", false, sjis)
	OutputFile(outputFile, "結果,", false, sjis)
	OutputFile(outputFile, "実行時間,", false, sjis)
	OutputFile(outputFile, "メモリ,", false, sjis)
	OutputFile(outputFile, "詳細", true, sjis)
}
