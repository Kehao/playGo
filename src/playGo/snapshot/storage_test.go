package snapshot
import (
  "testing"
  "github.com/sdegutis/go.assert"
)

func TestCheckRootDir(t *testing.T){
  root := "/home/qiu/gotest/go"
  fileName := "test.png"
  storage := &FileStorage{
    Root:root,
  }
  exist,err := storage.checkRootDir()
  if !exist{
    t.Error(err)
  }
  assert.Equals(t,storage.Retrieve(fileName),"/home/qiu/gotest/go/test.png")
  storage.Store(fileName)

  fileExist,_ := storage.Exist(fileName)
  assert.Equals(t,fileExist,true)

  storage.checkFile("log/app1.go")
  logfileExist,_ := storage.Exist("log/app1.go")
  assert.Equals(t,logfileExist,true)
  storage.RemoveAll()
}
