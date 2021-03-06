package receipt

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

var ReceiptsDirectory string = filepath.Join("uploads")

type Receipt struct {
	ReceiptName string    `json:"name"`
	UploadDate  time.Time `json:"uploadDate"`
}

func GetReceipts() ([]Receipt, error) {
	receipts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptsDirectory)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		receipts = append(receipts, Receipt{ReceiptName: f.Name(), UploadDate: f.ModTime()})
	}
	return receipts, nil
}
