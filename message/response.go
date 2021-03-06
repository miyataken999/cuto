// Copyright 2015 unirita Inc.
// Created 2015/04/10 honda

package message

import (
	"encoding/json"
	"fmt"
)

// レスポンスメッセージ送信元servantのバージョン
var ServantVersion string

type Response struct {
	Type       string `json:"type"`
	Version    string `json:"version"`
	NID        int    `json:"nid"`
	JID        string `json:"jid"`
	RC         int    `json:"rc"`
	Stat       int    `json:"stat"`
	Detail     string `json:"detail"`
	Var        string `json:"var"`
	St         string `json:"st"`
	Et         string `json:"et"`
	JoblogFile string `json:"joblogfile"`
}

const responseMessageType = "response"

// ジョブ実行結果JSONメッセージをパースし、Responseオブジェクトのメンバをセットする。
//
// param : メッセージ内容の文字列。
func (r *Response) ParseJSON(message string) error {
	byteMessage := []byte(message)

	err := json.Unmarshal(byteMessage, r)
	if err != nil {
		return err
	}

	if r.Type != responseMessageType {
		return fmt.Errorf("Invalid message type.")
	}

	return nil
}

// Responseオブジェクトの値を元に、ジョブ実行結果JSONメッセージを生成する
//
// return : JSONフォーマット整形後の文字列。
func (r Response) GenerateJSON() (string, error) {
	r.Type = responseMessageType
	r.Version = ServantVersion
	byteMessage, err := json.Marshal(r)
	if err != nil {
		return ``, err
	}

	return string(byteMessage), nil
}
