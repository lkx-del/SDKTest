package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	client "github.com/influxdata/influxdb1-client"
)

type Log struct {
	Timestamp int64  `json:"@timestamp"`
	Request   string `json:"request"`
}

func GetLinesFromFileAndJson(filename string, index int) []Log {
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

func copy_from_50w() {
	logdata := GetLinesFromFileAndJson("../resource/50w.txt", 500000)
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "127.0.0.1", 8086))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host, Username: "admin", Password: "At1314comi!"})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	q := client.Query{
		Command: "drop database clv; create database clv",
	}

	if r, err := con.Query(q); err != nil {
		fmt.Println("create database error:", r.Err)
		log.Fatal(err)
	}

	q = client.Query{
		Command:  "create measurement clvTable(logs string field, index idx1 logs type text)",
		Database: "clv",
	}

	if r, err := con.Query(q); err != nil {
		fmt.Println("create measurment(clv.clvTable) error:", r.Err)
		log.Fatal(err)
		panic(err)
	}

	s := time.Now().UnixMicro()
	k := 0
	i := 0
	for k = 0; k < 5000; k++ {
		pts := make([]client.Point, 0)
		var index int64
		for i = 100 * k; i < 100*k+100; i++ {
			pt := client.Point{
				Measurement: "clvTable",
				Fields: map[string]interface{}{
					"logs": logdata[i].Request,
				},
				Time:      time.Unix(0, logdata[i].Timestamp),
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

type TagLog struct {
	Timestamp int64  `json:"@timestamp"`
	Request   string `json:"request"`
	ClientIP  string `json:"clientip"`
}

func GetLinesFrom50Tag(filename string, index int) []TagLog {
	var logdata = make([]TagLog, 0)
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
		var logs TagLog
		err = json.Unmarshal(jsonStr, &logs)
		if err != nil {
			fmt.Println(i)
			panic("解析失败")
		}
		logdata = append(logdata, TagLog{
			Timestamp: logs.Timestamp,
			Request:   logs.Request,
			ClientIP:  logs.ClientIP,
		})
		if i == index {
			break
		}
	}
	return logdata
}

func copy_from_50wTag() {
	logdata := GetLinesFrom50Tag("../resource/50wTag.txt", 500000)
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "127.0.0.1", 8086))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host, Username: "admin", Password: "At1314comi!"})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	q := client.Query{
		Command: "drop database clvTag; create database clvTag",
	}

	if r, err := con.Query(q); err != nil {
		fmt.Println("create database(clvTag) error:", r.Err)
		log.Fatal(err)
		panic(err)
	}

	q = client.Query{
		Command:  "create measurement clvTable(clientip string tag, logs string field, index idx1 logs type text)",
		Database: "clvTag",
	}

	if r, err := con.Query(q); err != nil {
		fmt.Println("create measurment(clvTag.clvTable) error:", r.Err)
		log.Fatal(err)
		panic(err)
	}

	s := time.Now().UnixMicro()
	k := 0
	i := 0
	for k = 0; k < 5000; k++ {
		pts := make([]client.Point, 0)
		var index int64
		for i = 100 * k; i < 100*k+100; i++ {
			pt := client.Point{
				Measurement: "clvTable",
				Fields: map[string]interface{}{
					"logs": logdata[i].Request,
				},
				Tags: map[string]string{
					"clientip": logdata[i].ClientIP,
				},
				Time:      time.Unix(0, logdata[i].Timestamp),
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
			Database: "clvTag",
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

func main() {
	copy_from_50w()
	copy_from_50wTag()
}
