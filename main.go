package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"wechat-dat2picture/asset"
	"wechat-dat2picture/imgtype"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	var datPath string
	var picPath string
	flag.StringVar(&datPath, "datPath", "./dat", "待转换的dat文件目录")
	flag.StringVar(&picPath, "picPath", "./pic", "转换后的存放目录")
	flag.Parse()
	logrus.Info("datPath = ", datPath)
	logrus.Info("picPath = ", picPath)
	err := os.MkdirAll(picPath, 0775)
	if err != nil {
		logrus.Error("mkdir for picPath error: ", err)
		return
	}
	dat2Picture(datPath, picPath)
}

func dat2Picture(datPath string, picPath string) {
	bytesSampleDat, err := asset.Asset("sample/color.dat")
	if err != nil {
		logrus.Error("readFile for sampleDat error: ", err)
		return
	}
	bytesSampleJpg, err := asset.Asset("sample/color.jpg")
	if err != nil {
		logrus.Error("readFile for sampleJpg error: ", err)
		return
	}
	var rule [256]byte
	var mapped [256]bool
	var mappedCount int32
	for i, v := range bytesSampleDat {
		if mapped[int(v)] == false {
			rule[int(v)] = bytesSampleJpg[i]
			mapped[int(v)] = true
			mappedCount++
		}
		if mappedCount == 256 {
			break
		}
	}
	files, err := ioutil.ReadDir(datPath)
	if err != nil {
		logrus.Error("ReadDir for datPath error: ", err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			logrus.Info("ignore directory: ", f.Name())
			continue
		}
		if filepath.Ext(strings.TrimSpace(f.Name())) != ".dat" {
			logrus.Info("ignore not .dat file: ", f.Name())
			continue
		}

		datFilename := filepath.Join(datPath, f.Name())
		bytesDat, err := ioutil.ReadFile(datFilename)
		if err != nil {
			logrus.Warnf("readFile: %s error: %v", datFilename, err)
			continue
		}

		datHeader := [3]byte{bytesDat[0], bytesDat[1], bytesDat[2]}
		imgType, err := imgtype.Get(datHeader)
		if err != nil {
			logrus.Error("getImgType error: ", err)
			continue
		}
		var target []byte
		for _, bDat := range bytesDat {
			target = append(target, rule[bDat])
		}
		picFilename := filepath.Join(picPath, f.Name()+"."+imgType)
		err = ioutil.WriteFile(picFilename, target, 0644)
		if err != nil {
			logrus.Error("WriteFile error: ", err)
			continue
		}
	}
}
