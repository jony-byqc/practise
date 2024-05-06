package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = `{
    "code": 200,
    "requestId": "0a20f1aa-c300-4b26-b116-ef57df4e734a",
    "data": {
        "today": {
            "date_range": "today",
            "total_rows": 4,
            "success_rows": 4,
            "success_rate": "100.00%"
        },
        "month": {
            "date_range": "month",
            "total_rows": 63,
            "success_rows": 35,
            "success_rate": "58.21%"
        },
        "history": {
            "date_range": "history",
            "total_rows": 67,
            "success_rows": 37,
            "success_rate": "57.75%"
        },
        "chart": [
            {
                "date_range": "04-16",
                "total_rows": 4
            },
            {
                "date_range": "04-17",
                "total_rows": 4
            },
            {
                "date_range": "04-18",
                "total_rows": 2
            },
            {
                "date_range": "04-19",
                "total_rows": 22
            },
            {
                "date_range": "04-22",
                "total_rows": 2
            },
            {
                "date_range": "04-24",
                "total_rows": 9
            },
            {
                "date_range": "04-25",
                "total_rows": 20
            }
        ]
    }
}`
	fmt.Println(str)
	str = strings.Replace(str, "\n", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\t", "", -1)

	str = strings.Replace(str, " ", "", -1)
	fmt.Println(str)

}
