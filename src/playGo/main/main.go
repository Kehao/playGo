package main

import (
  "fmt"
  "log"
  "flag"
  "os"
  "github.com/kylelemons/go-gypsy/yaml"
  "io/ioutil"
  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/native" 
)
var (
  config *yaml.File
  config_file   = flag.String("conf", "./conf/database.yml", "the database.yml")
  log_file_name = flag.String("log", "./log/server.log", "where does the log go?")
)

func initlogAndConfig() {
  //create log
  log_file, err := os.OpenFile(*log_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
  if err != nil {
    panic(err)
  }
  log.SetOutput(log_file)
  log.SetFlags(5)

  //read the config and build config stuff
  c_file, err := ioutil.ReadFile(*config_file)
  if err != nil {
    log.Panic(err)
  }
  config = yaml.Config(string(c_file))
}

func checkError(err error) {
  if err != nil {
    panic(err)
    //os.Exit(1)
  }
}

func checkedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row,
mysql.Result) {
  checkError(err)
  return rows, res
}

func init(){
  initlogAndConfig()
}

func main(){
  username, _ := config.Get("development.username")
  password, _ := config.Get("development.password")
  database, _ := config.Get("development.database")

  db := mysql.New("tcp", "", "127.0.0.1:3306", username, password, database)
  checkError(db.Connect())


  sel, err := db.Prepare("SELECT * FROM notices n where subject_id = ? LIMIT 0,1000")
  checkError(err)

  rows, res := checkedResult(sel.Exec(1))
  id := res.Map("id")
  subjectType := res.Map("subject_type")

  for _, row := range rows {
    fmt.Printf(row.Str(id))
    fmt.Printf(row[id])
    fmt.Printf(row.Str(subjectType))
  }
}


