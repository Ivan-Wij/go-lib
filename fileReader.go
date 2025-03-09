package utils

import (
	"bufio"
	"os"
)

//type FileReaderContent struct {
//	Content []byte
//	IsEOF   bool
//	Err     error
//}
//type FileReader struct {
//	readChan chan FileReaderContent
//	this     *FileReader
//}
//
//func (fileReader *FileReader) OpenFile(path string) error {
//	fileReader.readChan = make(chan FileReaderContent, 5)
//	fileReader.this = fileReader
//
//	file, err := os.Open(path)
//	if err != nil {
//		return err
//	}
//
//	go func() {
//		defer file.Close()
//		for {
//			readBuff := make([]byte, 1024)
//			numBytes, err := file.Read(readBuff)
//
//			fileContent := FileReaderContent{
//				Content: readBuff[:numBytes],
//				IsEOF:   numBytes < 1024,
//				Err:     err,
//			}
//			fileReader.readChan <- fileContent
//
//			if fileContent.IsEOF && err != nil {
//				break
//			}
//		}
//	}()
//
//	return nil
//}
//
//func (fileReader *FileReader) Read() ([]byte, bool, error) {
//	fileContent := <-fileReader.readChan
//	return fileContent.Content, fileContent.IsEOF, fileContent.Err
//}
//
//func (fileReader *FileReader) CreateNByteChan(n int) (chan []byte, error) {
//	returnChan := make(chan []byte, 5)
//	go func() {
//		for {
//			readerBuff := make([]byte, n)
//			currBuffer, isEOF, err := fileReader.this.Read()
//			if err != nil {
//				return
//			}
//
//		}
//	}()
//
//}

type AsyncFileReader struct {
	file        *os.File
	ReceiveChan chan string
}

func (f *AsyncFileReader) InitializeAsyncReader(file *os.File) *AsyncFileReader {
	f.file = file
	f.ReceiveChan = make(chan string, 2)
	go func() {
		scanner := bufio.NewScanner(f.file)
		defer func() {
			close(f.ReceiveChan)
			f.file.Close()
		}()
		for scanner.Scan() {
			currLine := scanner.Text()
			f.ReceiveChan <- currLine
		}
	}()

	return f
}
