package main

import (
	"fmt"
	"log"
	"time"

	"github.com/f-velka/go-trac-rpc/v1_1_8"
	"github.com/rkl-/digest"
)

func main() {
	client, err := v1_1_8.NewClient(
		"http://192.168.1.102:8888/trac/SampleProject/login/rpc",
		// "http://localhost:8888",
		digest.NewTransport("admin", "admin"),
	)
	if err != nil {
		log.Fatal(err)
	}
	res1, err := client.Wiki.GetRecentChanges(time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC))
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res1 {
		fmt.Println(r)
	}

	pn := "WikiStart"
	res2, err := client.Wiki.GetPage(&v1_1_8.PageOptions{PageName: &pn})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res2)

	res3, err := client.Wiki.GetAllPages()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res3 {
		fmt.Println(r)
	}

	res4, err := client.Wiki.GetPageInfo(&v1_1_8.PageOptions{PageName: &pn})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res4)

	res5, err := client.Wiki.GetPageInfoVersion(&v1_1_8.PageOptions{PageName: &pn})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res5)

	res6, err := client.Wiki.PutPage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res6)
}
