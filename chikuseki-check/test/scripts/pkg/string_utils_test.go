package pkg

import (
	"chikuseki-check/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToString(t *testing.T) {

	var tests = []struct {
		name   string
		input  interface{}
		expect string
	}{
		{
			"Test string",
			"string",
			"string",
		},
		{
			"Test bytes",
			[]byte{65, 66, 67},
			"ABC",
		},
		{
			"Test number",
			12,
			"",
		},
		{
			"Test float",
			12.5,
			"",
		},
		{
			"Test struct",
			struct {
				name string
				age  int
			}{
				name: "admin",
				age:  12,
			},
			"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := utils.ToString(test.input)
			assert.Equal(t, result, test.expect)
		})
	}
}
