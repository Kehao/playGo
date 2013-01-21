package snapshot
import (
  "testing"
  "github.com/sdegutis/go.assert"
)

func TestArgs(t *testing.T) {
  args := map[string]string{
    "url":"http://www.baidu.com" ,
    "out":"test.png",
  }
  comm := &command{
    cutycaptArgs:args,
  }

  assert.Equals(t,comm.build() != "",true)
}
