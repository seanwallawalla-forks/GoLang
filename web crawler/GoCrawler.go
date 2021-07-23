/*
Title: GO CRAWLER
Author: Brad Myrick
Date: 2021-07-23
Description: A web crawler written in Go that saves the directory to a json.
*/
/*
GOALS:
ask a user for the url in console
ask for max depth
ask for max pages
ask for max time
ask for max time
create a webcrawler that will crawl the url and save the entire directory structure to a json file
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Crawler struct {
	url       string
	maxDepth  int
	maxPages  int
	maxTime   int
	file      string
	pages     []string
	dirs      []string
	files     []string
	startTime time.Time
	endTime   time.Time
}

func (c *Crawler) crawl() {
	var wg sync.WaitGroup
	c.startTime = time.Now()
	c.pages = append(c.pages, c.url)
	c.dirs = append(c.dirs, c.url)
	c.files = append(c.files, c.url)
	wg.Add(1)
	c.crawlPages(&wg)
	wg.Wait()
	c.endTime = time.Now()
	c.CrawlComplete()
	c.save()
}

func (c *Crawler) save() {
	fmt.Println("Saving...")
	file, err := os.Create(c.file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "{\n")
	fmt.Fprintf(file, "\"pages\": [\n")
	for i, page := range c.pages {
		if i == len(c.pages)-1 {
			fmt.Fprintf(file, "%s\n", page)
		} else {
			fmt.Fprintf(file, "%s,\n", page)
		}
	}
	fmt.Fprintf(file, "],\n")
	fmt.Fprintf(file, "\"dirs\": [\n")
	for i, dir := range c.dirs {
		if i == len(c.dirs)-1 {
			fmt.Fprintf(file, "%s\n", dir)
		} else {
			fmt.Fprintf(file, "%s,\n", dir)
		}
	}
	fmt.Fprintf(file, "],\n")
	fmt.Fprintf(file, "\"files\": [\n")
	for i, file := range c.files {
		if i == len(c.files)-1 {
			fmt.Fprintf(io.MultiWriter(), file, "%s\n", file)
		} else {
			fmt.Fprintf(io.MultiWriter(), file, "%s,\n", file)
		}
	}
	fmt.Fprintf(file, "]\n")
	fmt.Fprintf(file, "}\n")
}

func (c *Crawler) crawlPages(wg *sync.WaitGroup) {
	for i, page := range c.pages {
		if i >= c.maxPages {
			break
		}
		resp, err := http.Get(page)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			continue
		}
		if c.maxTime > 0 && time.Since(c.startTime) > time.Duration(c.maxTime) {
			//infinate run, change back to break
			continue
		}
		if resp.Header.Get("Content-Type") != "text/html" {
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(body), "404") {
			continue
		}
		links, err := c.getLinks(string(body))
		if err != nil {
			log.Fatal(err)
		}
		for _, link := range links {
			if c.isValidLink(link) {
				c.pages = append(c.pages, link)
				c.dirs = append(c.dirs, link)
				c.files = append(c.files, link)
			}
		}
	}
	wg.Done()
}

func (c *Crawler) isValidLink(link string) bool {
	if strings.HasPrefix(link, "http") {
		return true
	}
	if strings.HasPrefix(link, "/") {
		return true
	}
	return false
}

func (c *Crawler) getLinks(body string) ([]string, error) {
	links := []string{}
	for _, link := range strings.Split(body, "\n") {
		if strings.Contains(link, "href") {
			link = strings.Trim(link, "href=\"")
			link = strings.Trim(link, "href='")
			link = strings.Trim(link, "href=")
			link = strings.Trim(link, "\"")
			link = strings.Trim(link, "'")
			link = strings.Trim(link, " ")
			link = strings.Trim(link, "\t")
			link = strings.Trim(link, "\r")
			link = strings.Trim(link, "\n")
			if strings.Contains(link, "http") {
				links = append(links, link)
			}
		}
	}
	return links, nil
}

func (c *Crawler) CrawlComplete() {
	fmt.Println("Crawl Complete")
}

func main() {
	var url string
	var maxDepth int
	var maxPages int
	var file string
	var maxTime int
	fmt.Println("Welcome to the web crawler")
	fmt.Println("Please enter the url you want to crawl")
	fmt.Scanln(&url)
	fmt.Println("Please enter the max depth")
	fmt.Scanln(&maxDepth)
	fmt.Println("Please enter the max pages")
	fmt.Scanln(&maxPages)
	fmt.Println("Please enter the max time")
	fmt.Scanln(&maxTime)
	fmt.Println("Please enter the file name")
	fmt.Scanln(&file)
	file = file + ".json"
	c := Crawler{url: url, maxDepth: maxDepth, maxPages: maxPages, maxTime: maxTime, file: file}
	c.crawl()
}
