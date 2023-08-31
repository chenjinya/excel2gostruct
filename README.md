# excel2gostruct

Convert excel datasets to golang struct 


```golang
type Strategy struct {
	Number      string `json:"number"`
	Stage       string `json:"stage"`
	Description string `json:"description"`
	Strategy    string `json:"strategy"`
}

func main() {
	// 打开 Excel 文件
	targets := []Strategy{}
	err := ConvertExcelToGo("stage.xlsx", &targets)
	if err != nil {
		panic(err)
	}
	//println(utils.MustMarshalToString(targets))
}

```