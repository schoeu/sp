package main

import (
    "encoding/json"
    "io/ioutil"
)

type ConfProxy struct{
    Server string `json:"server"`
    LocalAddress string `json:"local_address"`
    LocalPort string `json:"local_port"`
    PortPassword PortPassword `json:"port_password"`
    Timeout int `json:"timeout"`
    Method string `json:"method"`
    FastOpen bool `json:"fast_open"`
}

type PortPassword map[string]string

func loadConf(p string) (*ConfProxy, error){
    content, err := ioutil.ReadFile(p)
    if err != nil {
        return nil, err
    }
    cp := &ConfProxy{}
    jsonErr := json.Unmarshal(content, cp)
    if jsonErr != nil {
        return nil, err 
    }
    return cp, nil
}