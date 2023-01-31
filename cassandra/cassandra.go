package main

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	var cassandraIP string
	fmt.Println("请输入cassandra host:")
	fmt.Scanln(&cassandraIP)
	cluster := gocql.NewCluster(cassandraIP + ":9042")
	cluster.Keyspace = "device_data"
	cluster.ProtoVersion = 3
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connect cassandra ok!")
	}
	defer session.Close()
	printState := func() {
		ctx := context.Background()
		scanner := session.Query("select device_code from device_data.device_data where device_code = '11412341503074390018_4AAa9-2';", "device_code").
			WithContext(ctx).Iter().Scanner()
		for scanner.Next() {
			var device_dode string
			err = scanner.Scan(&device_dode)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(device_dode)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	printState()
}
