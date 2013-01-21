package storage
import(
  "playGo/utils"
  "os"
  "path"
)

type Storage interface{
}

type FileStorage struct{
  Root string 
}

func (s *FileStorage) CheckFile(filePath string) (bool,error){
  if exist,_ := s.Exist(filePath);!exist{
    _,err := s.CheckRootDir() 
    checkError(err)
    err1 := os.MkdirAll(path.Dir(s.Retrieve(filePath)),0700)
    checkError(err1)
    _,err2 := s.Store(filePath)
    checkError(err2)
    return true,nil
  }
  return true,nil
}

func (s *FileStorage) CheckRootDir() (bool,error){
  if Exist,_ := utils.FileExists(s.Root);!Exist{
    err := os.MkdirAll(s.Root,0700)
    if err == nil{
      return true,nil
    }
    return false,err
  } 
  return true,nil
}

func (s *FileStorage) Store(filePath string) (*os.File,error){
  return os.Create(s.Retrieve(filePath))
}

func (s *FileStorage) Retrieve(filePath string) string{
  return path.Join(s.Root,filePath)
}

func (s *FileStorage) Remove(filePath string) error {
  return os.Remove(s.Retrieve(filePath))
}

func (s *FileStorage) RemoveAll() error {
  return os.RemoveAll(s.Root)
}

func (s *FileStorage) Exist(filePath string) (bool,error) {
  return utils.FileExists(s.Retrieve(filePath))
}

func checkError(err error) { if err != nil { panic(err) } }






