package snapshot
import (
  "testing"
  _ "github.com/sdegutis/go.assert"
)

func TestArgs(t *testing.T) {
  args := map[string]string{
    "Url":
    "http://www.zxaj.cn/search/index.php?act=detail&param=a%3A8%3A%7Bs%3A8%3A%22party_id%22%3Bs%3A8%3A%2225539373%22%3Bs%3A10%3A%22party_name%22%3Bs%3A22%3A%22%BA%A3%C4%FE%CA%D0%BA%EA%CD%FE%B7%C4%D6%AF%D3%D0%CF%DE%B9%AB%CB%BE%22%3Bs%3A11%3A%22card_number%22%3Bs%3A9%3A%22749010382%22%3Bs%3A7%3A%22case_id%22%3Bs%3A27%3A%22%282009%29%BC%CE%BA%A3%D0%D0%B7%C7%D6%B4%D7%D6%B5%DA00008%BA%C5%22%3Bs%3A8%3A%22reg_date%22%3Bs%3A10%3A%222009-02-25%22%3Bs%3A10%3A%22case_state%22%3Bs%3A4%3A%22%D2%D1%BD%E1%22%3Bs%3A11%3A%22apply_money%22%3Bs%3A8%3A%2250000.00%22%3Bs%3A10%3A%22court_name%22%3Bs%3A20%3A%22%D5%E3%BD%AD%CA%A1%BA%A3%C4%FE%CA%D0%C8%CB%C3%F1%B7%A8%D4%BA%22%3B%7D&type=public&PartyName=%BA%A3%C4%FE%CA%D0%BA%EA%CD%FE%B7%C4%D6%AF%D3%D0%CF%DE%B9%AB%CB%BE&CardNumber=",
    "out":"test.png",
  }
}
