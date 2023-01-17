package main

import (
	"fmt"
	"os"
)

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func main() {
	indexPath := "/home/guest09/media/guest09/ltc/" + "clvTable/" + "logs/" + "VTOKEN/" + "index/"
	//os.MkdirAll(indexPath, os.ModePerm)
	//indexFile, err := os.OpenFile(indexPath+"index0.txt", os.O_CREATE|os.O_WRONLY, 0644)
	//defer indexFile.Close()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	fmt.Println(FileExist(indexPath+"index1.txt"))

}
