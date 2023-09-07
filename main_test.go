package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// Output from Trading Paints downloader
type DriverInfo struct {
	Drivers []Paint `yaml:"Drivers"`
}

type DriverRoot struct {
	DriverInfo DriverInfo `yaml:"DriverInfo"`
}

func TestTPDownloader(t *testing.T) {
	var dr DriverRoot

	p, err := os.Open("tp.yaml")
	if err != nil {
		panic(err)
	}

	defer func() { _ = p.Close() }()

	buf, err := io.ReadAll(p)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &dr)
	assert.NoError(t, err)

	b, err := json.MarshalIndent(dr.DriverInfo.Drivers, "", " ")
	assert.NoError(t, err)

	fmt.Println(string(b))

	os.WriteFile("paints.json", b, 0644)
}
