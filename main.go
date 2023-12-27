package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"
)

type User struct {
	STR_ID    string
	USER_NAME string
}

func main() {
	// 解析命令行参数
	templatePath := flag.String("p", "", "Template file path")
	outputPath := flag.String("o", "release.txt", "Output file path")
	csvPath := flag.String("db", "", "CSV file path")
	randomIDLength := flag.Int("variable", 4, "Length of random ID")
	flag.Parse()

	// 检查模板文件路径是否为空
	if *templatePath == "" {
		log.Fatal("Template file path is required")
	}

	// 读取CSV文件
	csvFile, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true

	// 读取CSV文件的首行作为变量名
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// 创建变量映射
	vars := make([]User, 0)

	// 读取CSV文件的每一行，并将值与变量名对应
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		user := User{}
		for i, value := range record {
			if header[i] == "STR_ID" {
				user.STR_ID = value
			} else if header[i] == "USER_NAME" {
				user.USER_NAME = value
			}
		}

		vars = append(vars, user)
	}

	// 为每个迭代项生成随机ID并添加到变量映射中
	for i := range vars {
		if vars[i].STR_ID == "" {
			randomID := generateRandomID(*randomIDLength)
			vars[i].STR_ID = randomID
		}
	}

	// 解析模板文件
	tmpl, err := template.ParseFiles(*templatePath)
	if err != nil {
		log.Fatal(err)
	}

	// 创建输出文件
	outputFile, err := os.Create(*outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// 将模板渲染到输出文件
	err = tmpl.Execute(outputFile, vars)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File generated successfully.")
}

func generateRandomID(length int) string {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	charsetLength := len(charset)

	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charsetLength)
		sb.WriteByte(charset[randomIndex])
	}

	return sb.String()
}
