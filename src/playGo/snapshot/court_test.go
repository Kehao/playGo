package snapshot
import (
  "testing"
  "github.com/sdegutis/go.assert"
  "strings"
)

func TestNewSnapshot(t *testing.T) {
  args := map[string]string{
    "url":"http://www.baidu.com" ,
    "out":"test.png",
  }
  court := NewCourt(args,"1","1") 
  assert.True(t,court.storage.Root == appRoot + "/crimes/company-1")
  err := court.Run()
  t.Error(err)
  assert.True(t,strings.Contains(court.command.string(),args["url"]))
}
//Equals(t, "foo", "bar")
//NotEquals(t, "foo", "foo")
//DeepEquals(t, "foo", "bar")
//True(t, "foo" == "bar")
//False(t, "foo" == "foo")
//customEqualsForTesting(t, "1", "2")
//StringContains(t, "flub", "ber")


