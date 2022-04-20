package internal

import (
	"fmt"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// 標準出力
func OutputStdout(s string, doBreak bool) {
	fmt.Fprint(os.Stdout, s)
	if doBreak == true {
		fmt.Fprintln(os.Stdout, "")
	}
}

// ファイル出力
func OutputFile(o *os.File, s string, doBreak bool, sjis bool) error {
	if sjis == true {
		// Excelで扱うためにSJISで出力する
		writer := transform.NewWriter(o, japanese.ShiftJIS.NewEncoder())
		_, err := writer.Write([]byte(s))
		if err != nil {
			return err
		}
		if doBreak == true {
			_, err := writer.Write([]byte("\n"))
			if err != nil {
				return err
			}
		}
	} else {
		_, err := o.WriteString(s)
		if err != nil {
			return err
		}
		if doBreak == true {
			_, err := o.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 標準エラー出力
func OutputStderr(s string, doBreak bool) {
	fmt.Fprint(os.Stderr, s)
	if doBreak == true {
		fmt.Fprintln(os.Stderr, "")
	}
}
