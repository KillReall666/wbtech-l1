package main

import "fmt"

// интерфейс представленный методом SourceInterface
type SourceInterface interface {
	GetData() string
}

// интерфейс представленный методом ConvertData
type TargetInterface interface {
	ConvertData() string
}

// структура реализует SourceInterface через метод GetData()
type DataSource struct {
	Data string
}

func (s *DataSource) GetData() string {
	return s.Data
}

type TargetAdapter struct {
	Source SourceInterface
}

func (t *TargetAdapter) ConvertData() string {
	sourceData := t.Source.GetData()
	return "Converted: " + sourceData
}

func main() {
	dataSource := &DataSource{ // создается экземпляр DataSource со строковыми данными "Hello"
		Data: "Hello",
	}

	//данные dataSource передаются в TargetAdapter
	adapter := &TargetAdapter{
		Source: dataSource,
	}

	convertedData := adapter.ConvertData()
	fmt.Println(convertedData)
}
