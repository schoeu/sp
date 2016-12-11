package main

import (
    "encoding/json"
    "io/ioutil"
)

type Proxy struct{
    Server string `json:"server"`
    LocalAddress string `json:"local_address"`
    LocalPort string `json:"local_port"`
    PortPassword PortPassword `json:"port_password"`
    Timeout int `json:"timeout"`
    Method string `json:"method"`
    FastOpen bool `json:"fast_open"`
}

type PortPassword map[string]string

func loadConf(p string) (*Proxy, error){
    content, err := ioutil.ReadFile(p)
    if err != nil {
        return nil, err
    }
    p := &Proxy{}
    jsonErr := json.Unmarshal(content, p)
    if jsonErr != nil {
        return nil, err 
    }
    return p, nil
}