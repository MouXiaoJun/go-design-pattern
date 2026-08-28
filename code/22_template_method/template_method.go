package _22_template_method

import "strings"

// Processor 可变处理步骤接口，由具体管道实现
type Processor interface {
	Process(data string) string
}

// DataPipeline 数据处理管道骨架，固定「读取 → 处理 → 写出」流程
type DataPipeline struct {
	processor Processor // 模板持有步骤接口
}

// Run 模板方法：流程固定，只有处理步骤可变
func (p *DataPipeline) Run(input string) string {
	data := read(input)
	data = p.processor.Process(data)
	return write(data)
}

func read(input string) string {
	return strings.TrimRight(input, "\n") // 读取并去掉行尾换行
}

func write(data string) string {
	return data + "\n" // 写出并补回换行
}

// UpperPipeline 转大写管道：嵌入骨架，覆写处理步骤
type UpperPipeline struct {
	DataPipeline
}

func NewUpperPipeline() *UpperPipeline {
	p := &UpperPipeline{}
	p.processor = p // 把具体对象注入为可变步骤
	return p
}

func (p *UpperPipeline) Process(data string) string {
	return strings.ToUpper(data)
}

// TrimPipeline 去空白管道：嵌入骨架，覆写处理步骤
type TrimPipeline struct {
	DataPipeline
}

func NewTrimPipeline() *TrimPipeline {
	p := &TrimPipeline{}
	p.processor = p
	return p
}

func (p *TrimPipeline) Process(data string) string {
	return strings.TrimSpace(data)
}
