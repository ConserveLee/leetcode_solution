package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type problems []problem

type problem struct {
	ID			string
	Title		string
	Difficulty	int
	Lang		string
}

func newSolution() problems {
	const slPath = "../solution"
	///** 从文件读取 */
	ps := getSl(slPath)

	return ps
}

func (ps problems) PrintTable() string {
	return ps.table()
}

func (ps problems) table() string {
	res := "|题号|题目|难度|语言|\n"
	res += "| ---- | ---- | ---- | ---- |\n"
	for _, p := range ps {
		res += p.tableLine()
	}
	return res
}

func (p problem) tableLine() string {
	// 题号
	res := fmt.Sprintf("|[%s](%s)", p.ID, p.link())
	// 标题
	res += fmt.Sprintf("|[%s](%s)", p.Title, p.dir())
	// 难度
	res += fmt.Sprintf("|%s", p.diff())
	// 语言 TODO 以后dir写在语言下，同题目同一行tableLine
	res += fmt.Sprintf("|%s\n", p.Lang)

	return res
}

func (p problem) link() string {
	return fmt.Sprintf("https://leetcode-cn.com/problems/%s/", p.Title)
}

func (p problem) dir() string {
	switch p.Lang {
	case "go":
		return fmt.Sprintf("/%s/%s-%v/%s.go", "solution/go", p.ID, p.Difficulty, p.Title)
	case "python":
		return fmt.Sprintf("/%s/%s-%v/%s.go", "solution/python", p.ID, p.Difficulty, p.Title)
	}
	return "语言类型错误"
}

func (p problem) diff() string {
	switch p.Difficulty {
	case 1:
		return "简单"
	case 2:
		return "中等"
	case 3:
		return "困难"
	default:
		return "难度错误"
	}
}

func getSl(dir string) problems {
	ps	:= make([]problem, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return nil
		}
		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			pathArray := strings.Split(path, "/")
			IDMix 	  := strings.Split(pathArray[len(pathArray)-2], "-")
			titleMix  := strings.Split(info.Name(), ".")
			id, title, lang := IDMix[0], titleMix[0], titleMix[1]
			diff, _ := strconv.Atoi(IDMix[1])
			ps = append(ps, problem{
				ID:         id,
				Title:      title,
				Difficulty: diff,
				Lang:       lang,
			})
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return ps
}