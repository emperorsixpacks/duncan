package duncan

import (
	"meetUpGuru/m/duncan"
	"testing"
)

func TestLoadfromfile(t *testing.T){
  _ := duncan.NewFromConfig("../duncan_config.yml")
}

