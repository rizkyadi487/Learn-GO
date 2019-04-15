package main

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var fil []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			if filepath.Ext(path) == "" {
				//fmt.Println("index[", index, "] : Bukan File")

			} else {
				//fmt.Println("index[", index, "] :", filepath.Ext(value))
				//do nothing
				fil = append(fil, path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	//fmt.Printf("%v", fil)
	fmt.Println(len(fil))
	for index, value := range fil {
		fmt.Println("index[", index, "] :", value)
		passingBawah(value)
	}
}

func hash_file_crc32(filePath string, polynomial uint32) (string, error) {
	var returnCRC32String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnCRC32String, err
	}
	defer file.Close()
	tablePolynomial := crc32.MakeTable(polynomial)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, file); err != nil {
		return returnCRC32String, err
	}
	hashInBytes := hash.Sum(nil)[:]
	returnCRC32String = hex.EncodeToString(hashInBytes)
	return returnCRC32String, nil

}

func passingBawah(alamat string) {
	hash, err := hash_file_crc32(alamat, 0xedb88320)
	fullpath := alamat
	path := filepath.Dir(fullpath)
	file := filepath.Base(fullpath)
	format := filepath.Ext(fullpath)
	filename := file[0 : len(file)-len(format)]
	oldName := fullpath
	newName := path + "/" + filename + " [" + strings.ToUpper(hash) + "]" + format
	newName = strings.ReplaceAll(newName, "\\", "/")
	if err == nil {
		fmt.Println("Path :", path)
		fmt.Println("File :", file)
		fmt.Println("Filename :", filename)
		fmt.Println("Oldname:", oldName)
		fmt.Println("New name:", newName)
		os.Rename(oldName, newName)
		fmt.Println("File with name:", file, "has CRC32:", hash)
	}
}
