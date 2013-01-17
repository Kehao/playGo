package main

import (
  "fmt"
  "log"
  "flag"
  "os"
  "github.com/kylelemons/go-gypsy/yaml"
  "io/ioutil"
  "github.com/ziutek/mymysql/mysql"
  //_ "github.com/ziutek/mymysql/native" 
  _ "github.com/ziutek/mymysql/godrv"
  "playGo/funDb"
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
type Notices struct{
  Id int
  SubjectId int
  SubjectType string
}

func main(){
  //username, _ := config.Get("development.username")
  //password, _ := config.Get("development.password")
  //database, _ := config.Get("development.database")

  //db := mysql.New("tcp", "", "127.0.0.1:3306", username, password, database)
  //checkError(db.Connect())


  //sel, err := db.Prepare("SELECT * FROM notices n where subject_id = ? LIMIT 0,1000")
  //checkError(err)

  //rows, res := checkedResult(sel.Exec(1))
  //id := res.Map("id")
  //subjectType := res.Map("subject_type")

  //for _, row := range rows {
  //  fmt.Printf(row.Str(id))
  //  fmt.Printf(row[id])
  //  fmt.Printf(row.Str(subjectType))
  //}

  //_, _ = db.Query("select 1") 
  db, err := funDb.OpenMysql("mymysql", "tcp:localhost:3306*skyeye_development/root/admin123")
  checkError(err)

  var notices []Notices
  qi := funDb.SqlQueryInfo{
    Fields: "id, subject_id, subject_type",
    Where:  "id=2",
    //Params: []interface{}{0},
    Limit:  10,
    Offset: 0,
    Group:  "",
    Order:  "id desc",
  }
  err1 := db.GetStructs(&notices, qi)
  checkError(err1)
 
  fmt.Println(notices)

}


