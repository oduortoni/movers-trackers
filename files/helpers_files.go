package files

import (
	"os"
	"log"
	"encoding/gob"
	"bytes"
	
	"farmers/datatypes"
)
func CreateFile(file string) error {
	_, err := os.Create(file)
	if err != nil {
		return err
	}
	return nil
}
func Exists(filename string) bool {
     _, err := os.Stat(filename)
     if err == nil { //file exists
         return true
     }
     return false
}
func SetConfig(data *datatypes.Application) {
	config_file := "storage/config.data"
	if ok := Exists(config_file); ok == false {
	    err := CreateFile(config_file)
	    CheckError(err)
	}
	file, err := os.OpenFile("storage/config.data", os.O_WRONLY, 0777)
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()
	
	bb := new(bytes.Buffer)
	encoder := gob.NewEncoder(bb)
	err = encoder.Encode(data)
	CheckError(err)
	_, err = file.Write(bb.Bytes())
}
func GetConfig() *datatypes.Application {
	if ok := Exists("storage/config.data"); ok == false {
		log.Fatal("No config file present")
	}
	data :=  new(datatypes.Application)
	file, err := os.OpenFile("storage/config.data", os.O_RDONLY, 0777)
	if err != nil {
	    log.Fatal(err)
	}
	
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
	    log.Fatal(err)
	}
	return data
}
func CheckError(err error) {
	if err != nil {
	    if exists := Exists("storage/logfile"); exists == false {
	        CreateFile("storage/logfile")
	    }
		file, _ := os.OpenFile("storage/logfile", os.O_APPEND | os.O_WRONLY, 0777)
		defer file.Close()
		file.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

