// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:10
// @Software: GoLand

package model

import (
	"LRYGoCodeGen/utils"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type DirModel struct {
	Name  string
	Path  string
	Dirs  []DirModel
	Files []FileModel
}
type FileModel struct {
	Name    string
	Path    string
	Type    string
	Content []byte
}

func GetDirModel(dirPath string) (*DirModel, error) {
	var dirModel DirModel
	var err error
	if !utils.PathExists(dirPath) {
		return nil, errors.New("文件夹不存在！")
	}
	dirModel.Name = filepath.Base(dirPath)
	dirModel.Path = "."
	if err != nil {
		return nil, err
	}
	err = dirModel.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return &dirModel, nil
}
func (f FileModel) GetContentString() string {
	return string(f.Content)
}

func (d *DirModel) ReadDir(p string) error {
	dirs, err := ioutil.ReadDir(filepath.Join(p, d.Path))
	if err != nil {
		return err
	}
	for _, fi := range dirs {
		if fi.IsDir() {
			var dir DirModel
			dir.Name = fi.Name()
			dir.Path = filepath.Join(d.Path, fi.Name())
			err = dir.ReadDir(p)
			if err != nil {
				return err
			}
			d.Dirs = append(d.Dirs, dir)
		} else {
			var file FileModel
			file.Name = fi.Name()
			file.Path = filepath.Join(d.Path, fi.Name())
			file.Type = filepath.Ext(file.Path)
			file.Content, err = ioutil.ReadFile(filepath.Join(p, file.Path))
			if err != nil {
				return err
			}
			d.Files = append(d.Files, file)
		}
	}
	return nil
}

func (d *DirModel) MakeDir(p string, replaceDict KeyWord) error {
	var err error
	if !utils.PathExists(filepath.Join(p, d.Path)) {
		err = os.Mkdir(filepath.Join(p, d.Path), os.ModePerm)
		if err != nil {
			return err
		}
	}
	for _, fi := range d.Files {
		//todo 添加字符串替换逻辑
		s := string(fi.Content)
		for k, v := range replaceDict.Include {
			s = strings.ReplaceAll(s, k, v)
		}
		fi.Content = []byte(s)
		err = os.WriteFile(filepath.Join(p, fi.Path), fi.Content, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for _, di := range d.Dirs {
		if isIgnoreDir(di.Name, replaceDict.IgnoreDir) {
			continue
		}
		err = di.MakeDir(p, replaceDict)
		if err != nil {
			return err
		}
	}
	return nil
}

func isIgnoreDir(name string, nameList []string) bool {
	for _, d := range nameList {
		if matched, err := regexp.MatchString(d, name); matched {
			if err != nil {
				return true
			}
			return true
		}
	}
	return false
}
