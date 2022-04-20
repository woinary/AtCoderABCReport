package internal

import "strconv"

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
