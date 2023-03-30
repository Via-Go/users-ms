package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"os/exec"
)

const (
	makefileUrl  = "https://github.com/Via-Go/proto/blob/main/Makefile"
	genScriptUrl = "https://github.com/Via-Go/proto/blob/main/gen_server.sh"
	protoUrl     = "https://github.com/Via-Go/proto/blob/main/users.proto"
)

func writeToFile(filename string, data string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file: %s\nerr:%v", filename, err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		fmt.Printf("Failed to write to file: %s\nerr:%v", filename, err)
	}
}

func scrapFile(c *colly.Collector, url string) string {
	var data string

	c.OnHTML(".blob-code-inner", func(e *colly.HTMLElement) {
		data += e.Text
		data += "\n"
	})

	if err := c.Visit(url); err != nil {
		fmt.Printf("Failed to scrap url\nurl:%v\nerr:%v", url, err)
	}

	return data
}

func main() {
	c := colly.NewCollector()

	makefile := scrapFile(c, makefileUrl)
	proto := scrapFile(c, protoUrl)
	genScript := scrapFile(c, genScriptUrl)

	writeToFile("tmp/Makefile", makefile)
	writeToFile("tmp/users.proto", proto)
	writeToFile("tmp/gen_server.sh", genScript)

	os.Chmod("tmp/gen_server.sh", 0777)

	cmd := exec.Command("make", "users_server", "-C", "tmp")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run make server.yml\nerr: %v", err.Error())
	}

	fmt.Println(out.String())

}
