package snapshot
import (
  "log"
  "flag"
  "os"
  "io/ioutil"
  "github.com/kylelemons/go-gypsy/yaml"
)

const appRoot = "/home/qiu/work/playGo"

var (
  config *yaml.File
  confPath = flag.String("conf", "database.yml", "the database.yml")
  logPath = flag.String("log", "app.log", "the app log")
  appStorage = &FileStorage{ Root : appRoot,} 
)

func checkError(err error) {
  if err != nil {
    panic(err)
    //os.Exit(1)
  }
}

func initLogAndConfig() {
  logFile, err := os.OpenFile(*logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    panic(err)
  }
  log.SetOutput(logFile)
  log.SetFlags(5)

  //read the config and build config stuff
  confFile, err := ioutil.ReadFile(*confPath)
  if err != nil {
    log.Panic(err)
  }
  config = yaml.Config(string(confFile))
}

func init(){
  //appStorage.checkRootDir()
  //initLogAndConfig()
}
