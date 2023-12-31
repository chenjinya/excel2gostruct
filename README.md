# excel2gostruct

Convert excel datasets to golang struct 

将 Excel 的数据转换成 Golang 的结构体

# Usage

## function

### ConvertExcelToGo(filepath string, targets interface{}, sheetIndex int) error

Convert excel datasets to golang struct

Params:

| param                | desc                              |
|----------------------|-----------------------------------|
| filepath(string)     | the excel file path               
| targets(interface{}) | the target you want to convert to 
| sheetIndex(int)      | the sheet index in excel           


```golang
package main

import 	"github.com/chenjinya/excel2gostruct"

type Strategy struct {
	Number      string `json:"number"`
	Stage       string `json:"stage"`
	Description string `json:"description"`
	Strategy    string `json:"strategy"`
}

func main() {
    rawStrategies := []Strategy{}
    _ = excel2gostruct.ConvertExcelToGo("stage.xlsx", &rawStrategies, 1)
}

```

file `stage.xlsx` example
![stage.xlsx](https://pic4.zhimg.com/80/v2-f2d7433153d8478d94593f00e86ac253.jpg)

converted to struct (json)
```json
[
 {
  "number": "1",
  "stage": "想要退款",
  "description": "学生因为各种原因，表达不想继续学习了",
  "strategy": "先表达理解"
 },
 {
  "number": "2",
  "stage": "课程相关问题",
  "description": "涉及课程的具体服务",
  "strategy": "请在知识库里找到"
 },
 {
  "number": "3",
  "stage": "知识点提问",
  "description": "学生提了一个具",
  "strategy": "请先在知识库里找"
 }
]
```

# License

MIT