package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tktotal/pkg/util"
)

// GetKubun
func Test_GetKubun(t *testing.T) {
	var tests = []struct {
		name   string
		args   string
		expect string
	}{{
		name:   "Quote Code is empty",
		args:   "",
		expect: "",
	},
		{
			name:   "Quote Code is XAUD/CUR",
			args:   "XAUD/CUR",
			expect: "X",
		},
		{
			name:   "Quote Code is @@AUV/T",
			args:   "@@AUV/T",
			expect: "@@",
		},
		{
			name:   "Quote Code is E1301#0/T",
			args:   "E1301#0/T",
			expect: "E",
		},
		{
			name:   "Quote Code is E1301#0 don't have /",
			args:   "E1301#0",
			expect: "E",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.GetKubun(tt.args)
			assert.Equal(t, result, tt.expect)
		})
	}
}

// GetHassin
func Test_GetHassin(t *testing.T) {
	var tests = []struct {
		name   string
		args   string
		expect string
	}{
		{
			name:   "Quote Code is empty",
			args:   "",
			expect: "",
		},
		{
			name:   "Quote Code is XAUD/CUR",
			args:   "XAUD/CUR",
			expect: "CUR",
		},
		{
			name:   "Quote Code is @@AUV/T",
			args:   "@@AUV/T",
			expect: "T",
		},
		{
			name:   "Quote Code is E1301#0/T",
			args:   "E1301#0/T",
			expect: "T",
		},
		{
			name:   "Quote Code is E1301#0 don't have /",
			args:   "E1301#0",
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.GetHassin(tt.args)
			assert.Equal(t, result, tt.expect)
		})
	}
}

func Test_GetKubunHassin(t *testing.T) {
	var tests = []struct {
		name   string
		args   string
		expect string
	}{
		{
			name:   "Quote Code is empty",
			args:   "",
			expect: "/",
		},
		{
			name:   "Quote Code is XAUD/CUR",
			args:   "XAUD/CUR",
			expect: "X/CUR",
		},
		{
			name:   "Quote Code is E1822#0/T",
			args:   "E1822#0/T",
			expect: "E/T",
		},
		{
			name:   "Quote Code is @@AUV/T",
			args:   "@@AUV/T",
			expect: "@@/T",
		},
		{
			name:   "Quote Code is E1301#0 don't have /",
			args:   "E1301#0",
			expect: "E/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.GetKubunHassinPairFromQcd(tt.args)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func Test_SplitTexts(t *testing.T) {
	var tests = []struct {
		name      string
		fullText  string
		delimiter string
		expect    struct {
			FirstString  string
			SecondString string
		}
	}{
		{
			name:      "Input text is empty",
			fullText:  "",
			delimiter: ",",
			expect: struct {
				FirstString  string
				SecondString string
			}{FirstString: "", SecondString: ""},
		},
		{
			name:      "Delimiter is empty",
			fullText:  "this is sample text, here",
			delimiter: "",
			expect: struct {
				FirstString  string
				SecondString string
			}{FirstString: "", SecondString: "his is sample text, here"},
		},
		{
			name:      "Delimiter and input text is empty",
			fullText:  "",
			delimiter: "",
			expect: struct {
				FirstString  string
				SecondString string
			}{FirstString: "", SecondString: ""},
		},
		{
			name:      "Split string by comma character",
			fullText:  "this is sample text, here",
			delimiter: ",",
			expect: struct {
				FirstString  string
				SecondString string
			}{FirstString: "this is sample text", SecondString: " here"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstString, secondString := util.SplitTexts(tt.fullText, tt.delimiter)
			assert.Equal(t, tt.expect.FirstString, firstString)
			assert.Equal(t, tt.expect.SecondString, secondString)
		})
	}
}

func Test_GetDayName(t *testing.T) {
	var tests = []struct {
		name   string
		args   time.Weekday
		expect string
	}{
		{
			name:   "Input is monday",
			args:   time.Monday,
			expect: "MON",
		},
		{
			name:   "Input is tuesday",
			args:   time.Tuesday,
			expect: "TUE",
		},
		{
			name:   "Input is wednesday",
			args:   time.Wednesday,
			expect: "WED",
		},
		{
			name:   "Input is thursday",
			args:   time.Thursday,
			expect: "THU",
		},
		{
			name:   "Input is friday",
			args:   time.Friday,
			expect: "FRI",
		},
		{
			name:   "Input is saturday",
			args:   time.Saturday,
			expect: "SAT",
		},
		{
			name:   "Input is sunday",
			args:   time.Sunday,
			expect: "SUN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dayOfWeek := util.GetDayName(tt.args)
			assert.Equal(t, tt.expect, dayOfWeek)
		})
	}
}
