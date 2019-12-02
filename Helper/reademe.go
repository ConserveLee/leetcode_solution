package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func buildReadme() {
	log.Println("开始，重建 README 文档")

	sl := newSolution()
	makeReadmeFile(sl)

	log.Println("完成，重建 README 文档")
}

func makeReadmeFile(sl problems) {
	file := "../README.md"
	os.Remove(file)

	var b bytes.Buffer

	tmpl := template.Must(template.New("readme").Parse(readTMPL("template.markdown")))


	err := tmpl.Execute(&b, sl)
	if err != nil {
		log.Fatal(err)
	}

	// 保存 README.md 文件
	write(file, string(b.Bytes()))
}

func readTMPL(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}