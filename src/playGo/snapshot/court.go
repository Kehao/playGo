package snapshot
import (
  "playGo/storage"
  "path"
  "fmt"
  //"os/exec"
)

type Court struct {
  command *Command
  snapshotUrl string
  storage   *storage.FileStorage
}

func NewCourt(cutycaptArgs map[string]string,companyId,crimeId string) *Court {
  cutycaptArgs["out"] = fmt.Sprintf("%s.png",crimeId)
  return &Court {
    command: &Command { cutycaptArgs:cutycaptArgs },
    storage: &storage.FileStorage { Root: path.Join(appRoot, "crimes", fmt.Sprintf("company-%s",companyId))},
  }
}

func (court *Court) Run() string {
  court.storage.CheckRootDir()
  //cmd := exec.Command(court.command.string())
  //cmd:= exec.Command("xvfb-run","--server-args='-screen 0, 1024x768x24'","/home/qiu/work/playGo/CutyCapt/CutyCapt --out=1.png --url=http://www.baidu.com --js-can-open-windows=on")
  //out, err := cmd.Output()
  //fmt.Println(out)
  //fmt.Println(err)
  return court.command.string()//cmd.Run()
}




