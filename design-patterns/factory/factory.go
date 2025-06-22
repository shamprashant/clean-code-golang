/*
🧪 Your Task (Factory Pattern Practice):
Write a factory method GetStorage(type string) that returns:

S3Storage

AzureBlobStorage

LocalDiskStorage

All should implement:

type Storage interface {
	Save(file string)
}
Call Save() in main() with a fake file name.
*/

package main

import (
	"fmt"
	"log"
	"strings"
)

type Storage interface {
	Save(file string)
}

type S3Storage struct{}
type AzureBlobStorage struct{}
type LocalDiskStorage struct{}

func (s3 S3Storage) String() string {
	return "S3Storage"
}

func (abs AzureBlobStorage) String() string {
	return "AzureBlobStorage"
}

func (lds LocalDiskStorage) String() string {
	return "LocalDiskStorage"
}

func (s3 S3Storage) Save(file string) {
	fmt.Println("File saved on S3 with name: ", file)
}

func (abs AzureBlobStorage) Save(file string) {
	fmt.Println("File saved on Azure Blob with name: ", file)
}

func (lds LocalDiskStorage) Save(file string) {
	fmt.Println("File saved on Local Disk with name: ", file)
}

func GetStorage(StorageType string) (Storage, error) {
	StorageType = strings.ToLower(StorageType)
	switch StorageType {
	case "s3":
		return S3Storage{}, nil
	case "azure":
		return AzureBlobStorage{}, nil
	case "local":
		return LocalDiskStorage{}, nil
	default:
		return nil, fmt.Errorf("Unknown Storage type: %v", StorageType)
	}
}

func main() {
	storage, err := GetStorage("S3")
	if err != nil {
		log.Fatal(err)
	}
	storage.Save("budget.log")
}
