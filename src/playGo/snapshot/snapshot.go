package snapshot
import (
  "fmt"
  "log"
  "flag"
  "os"
  "io/ioutil"
  "github.com/kylelemons/go-gypsy/yaml"
  "playGo/storage"
)

const appRoot = "/home/qiu/work/playGo"

var appStorage = &storage.FileStorage{ Root : appRoot,} 

var (
  confPath = flag.String("conf", "conf/database.yml", "the database.yml")

  logPath = flag.String("log", "log/app.log", "the app log")

  Config *yaml.File
)

func Tp(msg interface{}){
  fmt.Println(msg) 
}

func initLogAndConfig(logPath,confPath string) {
  logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)
  if err != nil {
    panic(err)
  }
  log.SetOutput(logFile)
  log.SetFlags(5)

  //read the config and build config stuff
  confFile, err := ioutil.ReadFile(confPath)
  if err != nil {
    log.Panic(err)
  }
  Config = yaml.Config(string(confFile))
}

func checkError(err error) { if err != nil { panic(err) } }

func init(){
  appStorage.CheckFile(*logPath)
  appStorage.CheckFile(*confPath)
  initLogAndConfig(appStorage.Retrieve(*logPath),appStorage.Retrieve(*confPath))
  
}
