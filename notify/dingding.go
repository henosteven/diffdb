package notify

import (
    "bytes"
    "net/http"
    "io/ioutil"
    "conf"
    "fmt"
)

type Message struct {
    msgtype string
    text MessageContent
}

type MessageContent struct {
    content string
}

func SendDingDing(msgContent string) {
    body := bytes.NewBuffer([]byte("{'msgtype': 'text', 'text': {'content': '" + msgContent + "'}}"))
    res,err := http.Post(conf.DingDingUrl, "application/json;charset=utf-8", body)
    if err != nil {
        return
    }

    result, err := ioutil.ReadAll(res.Body)
    fmt.Println(result)
    res.Body.Close()
    if err != nil {
        return
    }
}
