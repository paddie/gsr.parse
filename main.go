package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

var (
	// path and regex are used to locate files
	path  string
	dir   string
	regex string
	// ids
	businessUnitId string
	consumerId     string
	reviewId       string
	merchantUrl    string
	locale         string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.StringVar(&path, "path", "", "path to a xml.gz or .xml file")
	flag.StringVar(&dir, "dir", "", "location of xml.gz or .xml files")

	flag.StringVar(&businessUnitId, "b", "", "Check for the existence of this business unit id")
	flag.StringVar(&merchantUrl, "url", "", "Check for the existense of this merchant url")
	flag.StringVar(&consumerId, "c", "", "Check for the existence of this consumer id")
	flag.StringVar(&reviewId, "r", "", "Check for the existence of review id")
	flag.StringVar(&locale, "l", "", "only look at the specified locale")

	flag.Parse()

	if path != "" {
		ProcessFeed(path)
		return
	}

	if dir != "" {
		ProcessDir(dir)
		return
	}

	ProcessDir(".")
}

func ProcessDir(dir string) error {
	potentials := []string{}

	finfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	if len(finfos) == 0 {
		return fmt.Errorf("No files in dir '%s'", dir)
	}

	locale = strings.ToLower(locale)

	for _, finfo := range finfos {
		if finfo.IsDir() || !strings.HasPrefix(finfo.Name(), "feed_") {
			continue
		}

		if strings.HasSuffix(finfo.Name(), "xml") || strings.HasSuffix(finfo.Name(), "xml.gz") {
			if locale != "" {
				if strings.Contains(strings.ToLower(finfo.Name()), locale) {
					potentials = append(potentials, finfo.Name())
				}
			} else {
				potentials = append(potentials, finfo.Name())
			}
		}
	}

	resp := make(chan int)
	for _, path := range potentials {
		go func(path string) {
			resp <- ProcessFeed(path)
		}(path)
	}
	matches := 0
	searching := businessUnitId == "" || consumerId == "" || reviewId == "" || merchantUrl == ""
	for _, _ = range potentials {
		matches += <-resp
	}

	if searching {
		fmt.Println("\nMatches: ", matches)
	} else {
		fmt.Println("\nDone")
	}

	return nil
}

func ProcessFeed(path string) int {
	log.Println("Processing: " + path)

	feed, err := ParseFeed(path)
	if err != nil {
		log.Println(err)
		fmt.Printf("%s failed to parse\n")
		return 0
	}

	if businessUnitId == "" && consumerId == "" && reviewId == "" && merchantUrl == "" {
		fmt.Printf("%s parsed successfully\n", path)
		return 0
	}

	matches := 0
	for _, m := range feed.Merchants {
		if businessUnitId != "" && businessUnitId == m.BusinessUnitId {
			fmt.Printf("%s:\n%s", path, m.String())
			matches++
		}

		if merchantUrl != "" && strings.HasPrefix(m.Url, merchantUrl) {
			fmt.Printf("%s:\n%s", path, m.String())
			matches++
		}

		if consumerId != "" || reviewId != "" {
			for _, r := range m.Reviews {
				if r.ReviewerId == consumerId {
					fmt.Printf("%s:\n%s", path, r.String())
					matches++
				}

				if r.Id == reviewId {
					fmt.Printf("%s:\n%s", path, r.String())
					matches++
				}
			}
		}
	}

	return matches
}
