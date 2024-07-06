package main

import (
	"fmt"
	"sync"
)

// Problem: Create concurrent Pipeline
// Data Loader -> Data Processor -> Data Writer

// 1. How would you implement the Data Loader stage to support concurrent data loading?
// Would you use a mutex, a channel, or another synchronization mechanism, and why?
// (for e.g. You have multiple source it could be database, excel, e.t.c.)

// 2. For the Data Processor stage,
// which requires exclusive access to a shared resource during processing, would a mutex be appropriate?
// If so, how would you use it to ensure that only one record is processed at a time?

// 3. In the Data Writer stage,
// considering that multiple processed records might be ready to write to the data store at the same time,
// how would you synchronize access to the data store?
// Discuss the trade-offs between using a mutex and a channel for this purpose.

// TODO: REST API(Stage status update)

// Data Loader: Completed
// Data Processor: In-Progress
// Data Loader: Pending
var count = 20

type SharedResource struct {
	Mutex sync.Mutex
}

type Data struct {
	ID    int
	Value string
}

type DataLoader interface {
	Load() ([]Data, error)
}

type DataProcessor interface {
	ProcessData(Data) (Data, error)
}

type DataWriter interface {
	WriteData(Data) error
}

// Stage 1 - Data Loading
type DatabaseLoader struct{}

func (dl *DatabaseLoader) Load() ([]Data, error) {
	// Simulate data loading from a database
	var data = []Data{}
	for i := 0; i < count; i++ {
		data = append(data, Data{ID: i, Value: fmt.Sprintf("Database data %d", i)})
	}

	return data, nil
}

type ExcelLoader struct{}

func (el *ExcelLoader) Load() ([]Data, error) {
	// Simulate data loading fron a database
	var data = []Data{}
	for i := 0; i < count; i++ {
		data = append(data, Data{ID: i, Value: fmt.Sprintf("Excel data %d", i)})
	}

	return data, nil
}

// Stage 2 - Data Processing
type SimpleProcessor struct {
	Shared *SharedResource
}

func (sp *SimpleProcessor) ProcessData(data Data) (Data, error) {
	// Simulate processing data
	var dataRes = Data{}
	sp.Shared.Mutex.Lock()
	defer sp.Shared.Mutex.Unlock()
	dataRes.ID = data.ID
	dataRes.Value = data.Value
	fmt.Println("Processed: ", dataRes.Value)
	return dataRes, nil
}

// Stage 3 - Data Writing
type ConsoleWriter struct {
	Shared *SharedResource
}

func (cw *ConsoleWriter) WriteData(data Data) error {
	cw.Shared.Mutex.Lock()
	defer cw.Shared.Mutex.Unlock()
	// Simulate writing data
	fmt.Println("Writing Data:", data)
	return nil
}

// For the Data Loader stage, which involves loading data from multiple sources (e.g., database, Excel files),
// you want to enable concurrent loading.
// This can be efficiently managed using channels for communication and coordination between goroutines in Go.

// Why Use Channels?
// Channels provide a safe way to pass data between goroutines, eliminating the need for explicit locks and avoiding potential deadlocks.
// Concurrency is managed naturally, as channels can be used to signal when data is available for processing.
func startLoader(wg *sync.WaitGroup, loaders []DataLoader, dataChan chan<- Data) {
	defer wg.Done()
	for _, loader := range loaders {
		data, err := loader.Load()
		if err != nil {
			fmt.Println("Error Loading data:", err)
			continue
		}

		for _, d := range data {
			dataChan <- d
		}
	}
	fmt.Println("Data Loader: Completed")
	close(dataChan)
}

// For the Data Processor stage, which requires exclusive access to a shared resource during processing, a mutex would be appropriate.
// This ensures that only one goroutine can process a record at a time, avoiding race conditions.
func startProcessing(wg *sync.WaitGroup, dataChan <-chan Data, processor *SimpleProcessor, processedChan chan<- Data) {
	defer wg.Done()
	for data := range dataChan {
		processedData, err := processor.ProcessData(data)
		if err != nil {
			fmt.Println("Error while processing data: ", err)
		}
		processedChan <- processedData
	}
	fmt.Println("Data processing: Completed")
	close(processedChan)
}

// For the Data Writer stage, where multiple processed records might be ready to write to the data store simultaneously, both mutexes and channels are viable options, but they come with trade-offs.

// Using Mutex
// Pros: Simple to implement, provides exclusive access.
// Cons: Can lead to contention and reduced throughput if many goroutines are competing for the lock.
// Using Channels
// Pros: Channels can help to decouple the data processing and writing stages, providing better throughput.
// Cons: More complex to manage, as you need to handle the channel capacity and ensure proper synchronization.

func startWriting(wg *sync.WaitGroup, processedChan <-chan Data, writer *ConsoleWriter) {
	defer wg.Done()
	for data := range processedChan {
		if err := writer.WriteData(data); err != nil {
			fmt.Println("Error while writing data: ", err)
			continue
		}
	}
	fmt.Println("Data writing: Completed")
}

func main() {
	var wg sync.WaitGroup
	// Intansiate the implementation
	loaders := []DataLoader{&DatabaseLoader{}, &ExcelLoader{}}
	processor := &SimpleProcessor{Shared: &SharedResource{}}
	writer := &ConsoleWriter{Shared: &SharedResource{}}

	dataChannel := make(chan Data, count*2)
	processedChannel := make(chan Data, count*2)

	wg.Add(3)
	// Data Loader - Stage 1
	go startLoader(&wg, loaders, dataChannel)

	// Data Processor - Stage 2
	// wg.Add(1)
	go startProcessing(&wg, dataChannel, processor, processedChannel)

	// Data Writing - Stage 3
	// wg.Add(1)
	go startWriting(&wg, processedChannel, writer)

	wg.Wait()

	fmt.Println("Pipeline: Completed Successfully")
}
