package main

import (
	"fmt"
	"io/ioutil"
)

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func main() {
	var s []string
	s, _ = GetAllFile("/home/guest09/media/guest09/ltc/bigdata/yhnTest/openGemini/lib/persistence/clvTable/logs/VTOKEN/index/", s)
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
