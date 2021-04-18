package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	tracrpc "github.com/f-velka/go-trac-rpc"
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
		return
	}
	t := time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
	res1, err := client.Wiki.GetRecentChanges(&t)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, r := range res1 {
		fmt.Println(r)
	}

	res2, err := client.Wiki.GetPage(tracrpc.String("WikiStart"), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print(res2)

	res3, err := client.Wiki.GetAllPages()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, r := range res3 {
		fmt.Println(r)
	}

	res4, err := client.Wiki.GetPageInfo(tracrpc.String("WikiStart"), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res4)

	res5, err := client.Wiki.GetPageInfoVersion(tracrpc.String("WikiStart"), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res5)

	res6, err := client.Wiki.PutPage(tracrpc.String("マインページ"), tracrpc.String("中身"), v1_1_8.PutPageAttributes{})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res6)

	res7, err := client.Wiki.ListAttachments(tracrpc.String("テストです"))
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, path := range res7 {
		res, err := client.Wiki.GetAttachment(tracrpc.String(path))
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s: %d\n", path, len(res))
		// os.Mkdir(filepath.Dir(path), 0777)
		// f, err := os.Create(path)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		// f.Write(res)
		// if err := f.Close(); err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
	}

	fmt.Println(os.Getwd())
	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	buf, _ := ioutil.ReadAll(f)
	res8, err := client.Wiki.PutAttachment(tracrpc.String("テストです/あああ333.txt"), buf)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res8)

	res9, err := client.Wiki.PutAttachmentEx(tracrpc.String("テストです"), tracrpc.String("EXEXEX.txt"), tracrpc.String("説明です"), buf, tracrpc.Bool(true))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res9)

	res10, err := client.Wiki.DeletePage(tracrpc.String("マインページ"), nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res10)

	res11, err := client.Wiki.DeleteAttachment(tracrpc.String("テストです/EXEXEX.txt"))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res11)

	// res12, err := client.Wiki.ListLinks(tracrpc.String("テストです"))
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println(res12)

	res13, err := client.Wiki.WikiToHtml(tracrpc.String(res2))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res13)
}
