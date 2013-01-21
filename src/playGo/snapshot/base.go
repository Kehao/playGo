package snapshot

import (
  "playGo/storage"
)

type Base struct {
  command *command
  snapshotUrl string
  storage   *storage.FileStorage
}

//new  func() interface{}
type command struct{
  xvfb string 
  cutycaptPath string 
  cutycaptArgs map[string]string
}

func (comm *command) build()(string){
  
}



//CutycaptPath, _ = Config.Get("development.cutycapt_path")

