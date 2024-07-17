package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	type Result struct {
		Columns  []string        `json:"columns"`
		Metadata []string        `json:"metadata"`
		Rows     [][]interface{} `json:"rows"`
	}

	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://influx-proxy.com", "")
	// Get query client
	queryAPI := client.QueryAPI("org")
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket: "ems_cloud_dev")
 |> range(start: 1721107999, stop: 1721108317)
 |> filter(fn: (r) => r._measurement == "ems100@28689")
 |> filter(fn: (r) => r.projectId=="4f537620d37d40e19dd25be5ca6ad941")
 |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
 |> drop(columns: ["_start","_stop","result","_measurement","ta/*/-ble"])`)
	var res = Result{}
	if err == nil {
		var count = 0
		for result.Next() {
			if count == 0 {
				m := result.TableMetadata()
				columns := m.Columns()
				res.Columns = make([]string, 0)
				res.Metadata = make([]string, 0)
				res.Rows = make([][]interface{}, 0)
				for k := range columns {
					res.Columns = append(res.Columns, columns[k].Name())
					res.Metadata = append(res.Metadata, columns[k].DataType())
				}
				values := make([]interface{}, len(columns))

				valueMap := result.Record().Values()
				for idx := 0; idx < len(columns); idx++ {
					values[idx] = valueMap[res.Columns[idx]]
				}
			}
			//values := make([]interface{}, len(result.TableMetadata().Columns()))
			//valueMap := result.Record().Values()
			//for idx := 0; idx < len(result.TableMetadata().Columns()); idx++ {
			//	values[idx] = valueMap[res.Columns[idx]]
			//}
			//if count >= len(res.Columns) || len(values) != len(res.Columns) {
			//	break
			//}
			//res.Rows = append(res.Rows, values)

			if count == 400 {
				fmt.Println("+++")
			}
			count++
		}

		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	client.Close()
}

func toBytes(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func fromBytes(data []byte, v interface{}) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.LittleEndian, v)
}
