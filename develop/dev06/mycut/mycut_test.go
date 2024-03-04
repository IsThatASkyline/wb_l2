package mycut

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyGrep(t *testing.T) {

	test_path := []string{
		`test_data.txt`,
	}

	exept_first := []string{"1. |3 qwwrsfb|", "2. 5 wetgzxc", "3. |2 qsfv|", "3. 2 QsFV|", "4. |5e wetcz", "5. |2r qsfv", "6. 5f wetz|", "7. |25 qsfax", "8. |5a wetd"}
	testnameF := fmt.Sprintf("Тест номер %d", 1)
	t.Run(testnameF, func(t *testing.T) {
		test := CreateCut("4,3", "3,1", false, test_path)
		err := test.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, exept_first, test.result)
	})

	testnameS := fmt.Sprintf("Тест номер %d", 2)
	t.Run(testnameS, func(t *testing.T) {
		test := CreateCut("4,3", "3,1", true, test_path)
		err := test.Run()
		assert.NoError(t, err)
		assert.EqualValues(t, []string(nil), test.result)
	})
}
