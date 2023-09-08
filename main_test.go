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

// Recreate paints.json from the TP downloader which has the current colours
// and customer id which is used to formulate the tga filename
// TODO iRating is also in the yaml which could be used to weight the AI drivers' skill level
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
