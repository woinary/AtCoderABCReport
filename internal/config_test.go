package internal

import "testing"

func TestConfig_ReadConfigFile(t *testing.T) {
	type fields struct {
		ContestBaseUrl   string
		ContestPrefix    string
		ContestCount     int
		OutputSjis       bool
		ContestStartPage int
		ContestEndPage   int
		ReportPerPage    int
		Condition        condition
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TEST_NOFILE",
			fields: fields{
				ContestBaseUrl:   "",
				ContestPrefix:    "",
				ContestCount:     0,
				ContestStartPage: 0,
				ContestEndPage:   0,
				ReportPerPage:    0,
				OutputSjis:       true,
				Condition: condition{
					ContestTask:     "",
					ContestLanguage: "",
					ContestStasus:   "",
					ContestUser:     "",
				},
			},
			args:    args{fileName: "no_file.yml"},
			wantErr: true,
		},
		{
			name: "TEST_01",
			fields: fields{
				ContestBaseUrl:   "",
				ContestPrefix:    "",
				ContestCount:     0,
				ContestStartPage: 0,
				ContestEndPage:   0,
				ReportPerPage:    0,
				OutputSjis:       false,
				Condition: condition{
					ContestTask:     "",
					ContestLanguage: "",
					ContestStasus:   "",
					ContestUser:     "",
				},
			},
			args:    args{fileName: "test_01.yml"},
			wantErr: false,
		},
		{
			name: "TEST_02",
			fields: fields{
				ContestBaseUrl:   "https://atcoder.jp",
				ContestPrefix:    "/contests/abc",
				ContestCount:     123,
				ContestStartPage: 1,
				ContestEndPage:   9999,
				ReportPerPage:    20,
				OutputSjis:       true,
				Condition: condition{
					ContestTask:     "abc247_a",
					ContestLanguage: "Go",
					ContestStasus:   "AC",
					ContestUser:     "user",
				},
			},
			args:    args{fileName: "test_02.yml"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{}
			err := c.ReadConfigFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ReadConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if c.ContestBaseUrl != tt.fields.ContestBaseUrl {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.ContestBaseUrl, tt.fields.ContestBaseUrl)
				}
				if c.ContestPrefix != tt.fields.ContestPrefix {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.ContestPrefix, tt.fields.ContestPrefix)
				}
				if c.ContestCount != tt.fields.ContestCount {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.ContestCount, tt.fields.ContestCount)
				}
				if c.ContestStartPage != tt.fields.ContestStartPage {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.ContestStartPage, tt.fields.ContestStartPage)
				}
				if c.ContestEndPage != tt.fields.ContestEndPage {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.ContestEndPage, tt.fields.ContestEndPage)
				}
				if c.ReportPerPage != tt.fields.ReportPerPage {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.ReportPerPage, tt.fields.ReportPerPage)
				}
				if c.Condition.ContestTask != tt.fields.Condition.ContestTask {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.Condition.ContestTask, tt.fields.Condition.ContestTask)
				}
				if c.Condition.ContestLanguage != tt.fields.Condition.ContestLanguage {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.Condition.ContestLanguage, tt.fields.Condition.ContestLanguage)
				}
				if c.Condition.ContestStasus != tt.fields.Condition.ContestStasus {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.Condition.ContestStasus, tt.fields.Condition.ContestStasus)
				}
				if c.Condition.ContestUser != tt.fields.Condition.ContestUser {
					t.Errorf("Config.ReadConfigFile() error = %v, want %v", c.Condition.ContestUser, tt.fields.Condition.ContestUser)
				}
			}

		})
	}
}
