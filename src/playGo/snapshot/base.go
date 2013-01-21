package snapshot

import (
  "playGo/storage"
  "strings"
  "fmt"
)

type Base struct {
  command *command
  snapshotUrl string
  storage   *storage.FileStorage
}

type command struct{
  xvfb string 
  cutycaptPath string 
  cutycaptArgs map[string]string
}

func (comm *command) build()(string){
  comm.xvfb = "xvfb-run --server-args='-screen 0, 1024x768x24'"
  comm.cutycaptPath, _ = Config.Get("development.cutycapt_path")
  cutycapt := make([]string,10)
  for prefix,value := range comm.cutycaptArgs{
    cutycapt = append(cutycapt,fmt.Sprintf("--%s=%s",prefix,value))
  }
  cutycapt = append(cutycapt,"--js-can-open-windows=on")
 return strings.Join([]string{
   comm.xvfb,
   comm.cutycaptPath,
   strings.Join(cutycapt," "),
 }," ") 
}
