package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

type Avatar interface{
	GetAvatarURL(c *client)(string, error)
}

type AuthAvatar struct{}
var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client)(string, error){
	if url, ok := c.userData["avatar_url"]; ok{
		if urlStr, ok := url.(string); ok{
			return urlStr, nil
		}

	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {}
var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client)(string, error){
	if userid, ok := c.userData["userid"]; ok{
		if useridStr, ok := userid.(string); ok{
			m := md5.New()
			io.WriteString(m, strings.ToLower(useridStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}