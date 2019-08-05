package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}


func (p *Page) save() error {
	file, err := os.OpenFile(p.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(p.Body)
	writer.Flush()
	return nil
}


func (p *Page) load() error {
	file, openError := os.Open(p.Title)
	if openError != nil {
		return openError
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, readError := reader.Read(p.Body)
	if readError != nil {
		return readError
	}
	return nil
}

func (this *Page) save2() (err error) {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page) load2(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title)
	return err
}

func main() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()
	body := make([]byte, 1024)

	page2 := Page{"Page.md", body}
	if page2.load() == nil {
		fmt.Println(page2.Body)
	}

	fmt.Fprint(os.Stdout, "aaa")




}
