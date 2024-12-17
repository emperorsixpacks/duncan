package duncan

import (
	"os"
	"testing"
)

func TestLoadfromfile(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	_, err := FromConfig("./duncan-config.yml")
	if err != nil {
		t.Error(err)
	}
}
