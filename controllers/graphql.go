package controllers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "reviewer-load-check/constants"
    "reviewer-load-check/models"
)

func post(request models.RequestParam, response interface{}) {
    request_json, _ := json.Marshal(request)

    // タイムアウトを30秒に指定してClient構造体を生成
    client := &http.Client{Timeout: time.Duration(30) * time.Second}

    // TODO: 100件よりも多く取得する場合、PageInfoを利用する
    // 生成したURLを元にRequest構造体を生成
    req, _ := http.NewRequest("POST", constants.REQ_URL, bytes.NewBuffer(request_json))

    // リクエストにヘッダ情報を追加
    req.Header.Add("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + constants.GITHUB_TOKEN)

    // POSTリクエスト発行
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()

    // レスポンスを取得
    byteArray, _ := ioutil.ReadAll(resp.Body)
    jsonBytes := ([]byte)(byteArray)
    if err := json.Unmarshal(jsonBytes, response); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }
}
