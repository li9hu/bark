package bark

import (
	"testing"
)

func TestGetBark(t *testing.T) {
	SetUp("", "")
	_, err := Run("test", "test", "te", Sound.滴嘟滴嘟)
	if err != nil {
		print(err.Error())
		return
	}

}
