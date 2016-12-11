// 客户端

package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "regexp"
    "flag"
	"fmt"
)
type ConfSite []string


func main() {

    confPath := flag.String("c","","path of config file for site map")

    flag.Parse()

    if *confPath == "" {
        log.Fatal("please input config file path.")
    }

    _, err := LoadConf(*confPath)

    s := "www.baidu.com"

    fmt.Println(regexp.MatchString("baidu.com", s))

    if err != err {
        log.Fatal("conf error")
    }
}

func LoadConf(p string) (*ConfSite, error){
    content, err := ioutil.ReadFile(p)
    if err != nil {
        return nil, err
    }
    cp := &ConfSite{}
    jsonErr := json.Unmarshal(content, cp)
    if jsonErr != nil {
        return nil, err 
    }
    return cp, nil
}
