package duncan

import (
	"testing"
)

func TestLoadfromfile(t *testing.T) {
	_, err := FromConfig("./duncan-config.yml")
	if err != nil {
		t.Error(err)
	}
}
