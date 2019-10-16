package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
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

type FileSystemAvatar struct {}
var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client)(string, error){
	if userid, ok := c.userData["userid"]; ok{
		if useridStr, ok := userid.(string); ok{
			if files, err := ioutil.ReadDir("avatars"); err == nil{
				fmt.Println(files)
				for _, file := range files{
					if file.IsDir(){
						continue
					}
					if match, _:= filepath.Match(useridStr+"*", file.Name()); match{
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}









