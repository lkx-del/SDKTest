package main

import (
	"bufio"
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"os"
	"time"
	//"time"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "admin",
		Password: "At1314comi!",
	})
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	defer c.Close()

	fuzzyQuery := make([]string, 0)
	file, err := os.Open("fuzzy.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() //string
		querystr := "select * from clvTable where logs fuzzy '" + line + "'"
		fuzzyQuery = append(fuzzyQuery, querystr)
	}

	for i := 0; i < len(fuzzyQuery); i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(fuzzyQuery[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			if len(response.Results) == 0 {
				continue
			}
			if len(response.Results[0].Series) == 0 {
				continue
			}
			fmt.Println(len(response.Results[0].Series[0].Values))
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}
}
