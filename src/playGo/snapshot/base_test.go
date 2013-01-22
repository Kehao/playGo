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
  comm := &Command{
    cutycaptArgs:args,
  }

  assert.True(t,comm.string() != "")
}
//Equals(t, "foo", "bar")
//NotEquals(t, "foo", "foo")
//DeepEquals(t, "foo", "bar")
//True(t, "foo" == "bar")
//False(t, "foo" == "foo")
//customEqualsForTesting(t, "1", "2")
//StringContains(t, "flub", "ber")


