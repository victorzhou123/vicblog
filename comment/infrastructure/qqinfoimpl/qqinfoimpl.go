package qqinfoimpl

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	commentent "github.com/victorzhou123/vicblog/comment/domain/comment/entity"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/entity"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/qqinfo"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type qqInfoImpl struct {
	cfg Config
}

func NewQQInfoImpl(cfg Config) qqinfo.QQInfo {
	return &qqInfoImpl{cfg}
}

func (impl *qqInfoImpl) GetQQInfo(num entity.QQNumber) (commentent.CommentUserInfo, error) {

	qqInfo, err := impl.get(num.QQNumber())
	if err != nil {
		return commentent.CommentUserInfo{}, err
	}
	if !qqInfo.isSuccess() {
		return commentent.CommentUserInfo{}, errors.New("get qq information failed")
	}

	avatar, err := cmprimitive.NewUrlx(qqInfo.ImgURL)
	if err != nil {
		return commentent.CommentUserInfo{}, err
	}

	email, err := cmprimitive.NewEmail(qqInfo.Mail)
	if err != nil {
		return commentent.CommentUserInfo{}, err
	}

	return commentent.CommentUserInfo{
		Avatar:   avatar,
		Email:    email,
		NickName: qqInfo.Name,
	}, nil

}

type qqInfo struct {
	Code   int    `json:"code"`
	ImgURL string `json:"imgurl"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
}

func (r *qqInfo) isSuccess() bool {
	return r.Code == 200
}

func (impl *qqInfoImpl) get(num string) (qqInfo, error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", impl.cfg.URL+"?qq="+num, nil)
	if err != nil {
		return qqInfo{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return qqInfo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return qqInfo{}, err
	}

	info := qqInfo{}
	if err := json.Unmarshal(body, &info); err != nil {
		return qqInfo{}, err
	}

	return info, nil
}
