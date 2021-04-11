package main

import (
	"fmt"
	"log"

	"github.com/f-velka/go-trac-rpc/v1_1_8"
	"github.com/rkl-/digest"
)

func main() {
	client, err := v1_1_8.NewClient(
		"http://192.168.1.102:8888/trac/SampleProject/login/rpc",
		digest.NewTransport("admin", "admin"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// res1, err := client.Wiki.GetRecentChanges(time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, r := range res1 {
	// 	fmt.Println(r)
	// }
	res1, err := client.Wiki.GetAllPages()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res1 {
		fmt.Println(r)
	}
	// res, err := client.Wiki.GetPage(&v1_1_8.GetPageOption{
	// 	PageName: common.NewString("WikiStart"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(res)
}
