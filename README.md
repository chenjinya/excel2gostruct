# excel2gostruct

Convert excel datasets to golang struct 


# Usage
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
![stage.xlsx](https://github.com/chenjinya/excel2gostruct/blob/main/stage_xlsx.png)

# License

MIT