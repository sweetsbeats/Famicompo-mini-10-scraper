package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeAndRename(link string, folder string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println("Could not get access to famitracker.org, err: ", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	EntryCount := 0
	var names []string

	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		EntryCount++

		test := s.Children().First()

		fmt.Println(test)
		test2 := test.Children().First()
		fmt.Println(test2)

		test3 := test2.Children().First()
		fmt.Println(test3)
		if test3.Has("b") != nil {
		}
		b := test3.Find("b").Text()
		fmt.Println(b)
		// Taking names
		names = append(names, b)
	})

	fmt.Println(EntryCount)
	//fmt.Println(names)

	fileNamePrefix := folder + "/Entry"
	fileNameSuffix := ".nsf"

	var oldName string
	// Sanitize inputs
	for i, name := range names {
		name = strings.Replace(name, ":", "", -1)
		name = strings.Replace(name, "[", "", -1)
		name = strings.Replace(name, "]", "", -1)
		name = strings.Replace(name, "/", "", -1)
		name = strings.Replace(name, "\"", "", -1)
		name = strings.Replace(name, "?", "", -1)
		name = strings.Replace(name, "*", "", -1)
		name = strings.Replace(name, "<", "", -1)
		name = strings.Replace(name, ">", "", -1)
		name = strings.Replace(name, "|", "", -1)
		names[i] = name + fileNameSuffix
	}

	iter := 1
	for iter <= EntryCount {
		if iter < 10 {
			oldName = fileNamePrefix + "00" + strconv.Itoa(iter) + fileNameSuffix
		} else if iter < 100 {
			oldName = fileNamePrefix + "0" + strconv.Itoa(iter) + fileNameSuffix
		} else {
			oldName = fileNamePrefix + strconv.Itoa(iter) + fileNameSuffix
		}

		_, err := os.Stat(oldName)
		if err == nil {
			fmt.Println("file exists!")
			err := os.Rename(oldName, folder+"/"+names[iter-1])
			if err != nil {
				fmt.Println(err)
			}
		}
		iter++
	}
}

func main() {
	ScrapeAndRename("https://famitracker.org/famicompo/fcm/vol10/list_c.html", "Cover")
	ScrapeAndRename("https://famitracker.org/famicompo/fcm/vol10/list_o.html", "Original")
}
