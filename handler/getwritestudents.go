package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func getStudentsList() (*StudentList, error) {
	f, err := os.Open("students.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var sl StudentList

	jsDocument, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsDocument, &sl); err != nil {
		return nil, err
	}
	return &sl, nil
}

func writeStudentToFile(sl *StudentList) error {
	jsonContent, err := json.MarshalIndent(sl, " ", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("students.json", jsonContent, 0644)
	if err != nil {
		return err
	}
	return nil
}
