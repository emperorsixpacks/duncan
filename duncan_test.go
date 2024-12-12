package duncan

import (
	"testing"
)

func TestLoadfromfile(t *testing.T){
  err := NewFromConfig("./duncan-config.yml")
  if err != nil{
    t.Error(err)
  }
}

