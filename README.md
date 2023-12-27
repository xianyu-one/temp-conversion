# temp-conversion

golang temp模板转换小工具

## 编译

```
go build -ldflags="-s -w" -o conversion main.go && upx -9 conversion
```


## 使用方法

go run main.go -p <模板文件路径> -o <输出文件路径> -db <CSV文件路径> -variable <随机ID长度>


- `-p`：指定模板文件的路径
- `-o`：指定输出文件的路径（可选，默认为`release.txt`）
- `-db`：指定CSV文件的路径
- `-variable`：指定随机ID的长度（可选，默认为4）

### 模板文件

模板文件使用Go的文本模板语法，其中可以包含变量和控制结构。程序会根据提供的CSV文件中的数据替换模板中的变量，并将渲染结果写入输出文件。

以下是模板文件的示例：

```temp
{{range .}}
{
    "id": "{{.USER_ID}}",
    "title": "请给{{.USER_NAME}}打分",
    "type": "Score",
    "attribute": {
        "required": true
    },
    "children": [
        {
            "attribute": {
                "scoreStyle": "star",
                "scope": "[0,10]"
            },
            "id": "{{.STR_ID}}"
        }
    ]
}
{{end}}
```

在模板文件中，可以使用和`.USER_NAME`来引用CSV文件中的对应列数据。


### CSV文件
CSV文件应包含列标题行和相应的数据行。程序会根据列标题与数据行的值进行匹配，将值与变量名对应。

以下是CSV文件的示例：

```csv
USER_ID,USER_NAME
12345678,Alice
87654321,Bob
```

在上面的示例中，第一行是列标题行，第二行是数据行。程序会将USER_ID列的值与USER_ID进行匹配，将USER_NAME列的值与.USER_NAME进行匹配。

### 随机ID

`.STR_ID`变量用于生成随机ID,随机ID由小写字母和数字组成。

### 示例
假设有一个模板文件template.json，CSV文件data.csv，随机ID长度为6，可以使用以下命令来生成输出文件：

```bash
go run main.go -p template -o output.txt -db data.csv -variable 6
```

程序将根据模板文件和CSV文件的数据生成输出文件output.txt。