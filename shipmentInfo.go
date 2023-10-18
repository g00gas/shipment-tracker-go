package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ShipmentInfo struct {
	ShipmentUrl   string `json:"shipmentUrl"`
	EmailToNotify string `json:"emailToNotfiy"`
}

func (s ShipmentInfo) checkUrlValidity() bool {
	pattern := regexp.MustCompile(`^(https?)://([A-Za-z0-9.-]+)(:[0-9]+)?(/[^/\s]*)?(\?[^\s]*)?$`)
	matched := pattern.MatchString(s.ShipmentUrl)
	if matched == false {
		log.Fatal("Url %s is invalid.", s.ShipmentUrl)
	}
	return matched
}

func (s ShipmentInfo) checkEmailToNotify() bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	matched := pattern.MatchString(s.EmailToNotify)
	if matched == false {
		log.Fatal("Email %s is invalid.", s.EmailToNotify)
	}
	return matched
}

func readShipmentFromFile() []ShipmentInfo {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Add path to JSON with shipment info:")

	path, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	path = strings.TrimSpace(path)
	fileExtension := filepath.Ext(path)
	if fileExtension != ".json" {
		fmt.Printf("extension %s", fileExtension)
		panic("File extension is not .json")
	}
	jsonContents, err := os.Open(path)
	if err != nil {
		log.Fatal("File not found.", err)
	}
	defer jsonContents.Close()
	jsonString, err := io.ReadAll(jsonContents)
	if err != nil {
		log.Fatal("Could not read the file", err)
	}
	fmt.Printf(string(jsonString))
	var ShipmentInfo []ShipmentInfo
	error := json.Unmarshal(jsonString, &ShipmentInfo)
	if error != nil {
		log.Fatal("error", err)
	}
	for _, s := range ShipmentInfo {
		s.validateShipmentInfo()
	}

	return ShipmentInfo
}

func (s ShipmentInfo) validateShipmentInfo() bool {
	isEmailOk := s.checkEmailToNotify()
	isUrlOk := s.checkUrlValidity()

	if isEmailOk && isUrlOk {
		return true
	} else {
		return false
	}

}
