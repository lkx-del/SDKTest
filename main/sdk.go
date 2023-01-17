package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	client "github.com/influxdata/influxdb1-client"
	"io"
	"log"
	"net/url"
	"os"
	"time"
)

type Log struct {
	Timestamp int64  `json:"@timestamp"`
	Request   string `json:"request"`
}

func GetLinesFromFileAndJson(filename string, index int) []Log { //
	var logdata = make([]Log, 0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buff := bufio.NewReader(file)
	i := 0
	for {
		i++
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		jsonStr := data
		var logs Log
		err = json.Unmarshal(jsonStr, &logs)
		if err != nil {
			fmt.Println(i)
			panic("解析失败")
		}
		logdata = append(logdata, Log{
			Timestamp: logs.Timestamp,
			Request:   logs.Request,
		})
		if i == index {
			break
		}
	}
	return logdata
}

func main() {
	logdata := GetLinesFromFileAndJson("/mnt/sdc/data/clv_data/one_billion_all.txt", 1000000000)
	//logdata := GetLinesFromFileAndJson("../resource/50w.txt", 500000) //  ../resource/50w.txt  /mnt/sdc/data/clv_data/one_billion_01.txt
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "127.0.0.1", 8086)) // /write?db=clv&u=admin&p=At1314comi!
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host, Username: "admin", Password: "At1314comi!"})
	if err != nil {
		log.Fatal(err)
	}
	s := time.Now().UnixMicro()

	k := 0
	i := 0
	for k = 0; k < 10000000; k++ { //for k := 0; k < 5000; k++ 10000000
		pts := make([]client.Point, 0)
		var index int64
		for i = 100 * k; i < 100*k+100; i++ { //for i := 100 * k; i < 100*k+100; i++
			pt := client.Point{
				Measurement: "clvTable",
				//Tags: map[string]string{
				// "host": "A" + string(k),
				//},
				Fields: map[string]interface{}{
					"logs": logdata[i].Request,
				},
				Time:      time.Unix(0, logdata[i].Timestamp), //  int64(i)
				Precision: "ns",
			}
			pts = append(pts, pt)
			index++
		}
		if k%10000 == 0 {
			fmt.Println("current time：", time.Now(), "write(w): ", k)
		}
		bps := client.BatchPoints{
			Points:   pts,
			Database: "clv",
		}
		for retry := 0; retry < 5; retry++ {
			_, err := con.Write(bps)
			if err == nil {
				break
			}
			fmt.Println("current time：", time.Now(), "panic error, (k,i, retry):", k, i, retry, " error: ", err)
		}
	}
	e := time.Now().UnixMicro()
	fmt.Println("sdk push logs cost time:")
	fmt.Println(float64(e-s) / 1000)
}
