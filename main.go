package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
)

const (
	Audi             = "Audi 90 GTO"
	MaxDrivers       = 100
	MinDriverAge     = 20
	MaxDriverAge     = 65
	PitCrewSkill     = 90
	CarIDAudi        = 76
	CarIDNissan      = 77
	CarClassIDAudi   = 83
	CarClassIDNissan = 84
)

type Driver struct {
	Class   string `csv:"Class"`
	Pos     string `csv:"Pos"`
	Name    string `csv:"Driver"`
	Car     string `csv:"Car"`
	Ignore  string `csv:"-"`
	Points  int    `csv:"Points"`
	Counted string `csv:"Counted"`
}

type AIDriver struct {
	DriverName        string `json:"driverName"`        // "Ben Regan"
	CarNumber         string `json:"carNumber"`         // "00"
	CarDesign         string `json:"carDesign"`         // "11,ffffff,111111,fc0706"
	CarTgaName        string `json:"carTgaName"`        // car_name.tga
	SuitDesign        string `json:"suitDesign"`        // "23,ffffff,111111,fc0706"
	HelmetDesign      string `json:"helmetDesign"`      // "68,ffffff,111111,fc0706"
	CarPath           string `json:"carPath"`           // "audi90gto" | "nissangtpzxt"
	CarID             int    `json:"carId"`             // 76 | 77
	CarClassID        int    `json:"carClassId"`        // 83 | 84
	Sponsor1          int    `json:"sponsor1"`          // 51
	Sponsor2          int    `json:"sponsor2"`          // 2
	NumberDesign      string `json:"numberDesign"`      // "0,0,,,"
	DriverSkill       int    `json:"driverSkill"`       // 67
	DriverAggression  int    `json:"driverAggression"`  // 39
	DriverOptimism    int    `json:"driverOptimism"`    // 87
	DriverSmoothness  int    `json:"driverSmoothness"`  // 67
	PitCrewSkill      int    `json:"pitCrewSkill"`      // 77
	StrategyRiskiness int    `json:"strategyRiskiness"` // 100
	DriverAge         int    `json:"driverAge"`         // 37
	ID                string `json:"id"`                // "27db8f7a-9a6d-cd97-3e91-03d53b343029"
	RowIndex          int    `json:"rowIndex"`          // 0
}

type AIRoster struct {
	AIDrivers []AIDriver `json:"drivers"`
}

var (
	Rnd = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
)

func main() {
	f, err := os.Open("drivers.csv")
	if err != nil {
		panic(fmt.Sprintf("error: opening file %s with %v\n", os.Args[1], err))
	}

	defer func() { _ = f.Close() }()

	drivers := []*Driver{}

	if err := gocsv.UnmarshalFile(f, &drivers); err != nil {
		panic(err)
	}

	sort.Slice(drivers, func(i, j int) bool { return drivers[i].Points > drivers[j].Points })

	topDrivers := []Driver{}
	driven := make(map[string]bool)

	for _, driver := range drivers {
		if ok := driven[driver.Name]; ok {
			continue
		}

		topDrivers = append(topDrivers, *driver)
		driven[driver.Name] = true

		if len(driven) == MaxDrivers {
			break
		}
	}

	airoster := AIRoster{}

	for i, driver := range topDrivers {
		aiDriver := AIDriver{
			ID:                uuid.New().String(),
			RowIndex:          i,
			CarNumber:         fmt.Sprintf("%02d", i+1),
			CarPath:           carPath(driver.Car),
			CarID:             carID(driver.Car),
			CarClassID:        carClassID(driver.Car),
			DriverName:        driver.Name,
			DriverAge:         rnd(MinDriverAge, MaxDriverAge),
			DriverSkill:       weightedRnd(i, driver.Points),
			DriverAggression:  weightedRnd(i, driver.Points),
			DriverOptimism:    weightedRnd(i, driver.Points),
			DriverSmoothness:  weightedRnd(i, driver.Points),
			PitCrewSkill:      PitCrewSkill,
			StrategyRiskiness: weightedRnd(i, driver.Points),

			CarDesign:    "",
			SuitDesign:   "",
			HelmetDesign: "",
			NumberDesign: "0,0,,,",
		}

		airoster.AIDrivers = append(airoster.AIDrivers, aiDriver)
	}

	b, err := json.MarshalIndent(airoster, "", "  ")
	if err != nil {
		panic(err)
	}

	//	for i := range airoster.AIDrivers {
	//		fmt.Println(airoster.AIDrivers[i].DriverName, airoster.AIDrivers[i].DriverSkill, airoster.AIDrivers[i].CarPath)
	//	}

	fmt.Println(string(b))
}

func carPath(carName string) string {
	if carName == Audi {
		return "audi90gto"
	}

	return "nissangtpzxt"
}

func carID(carName string) int {
	if carName == Audi {
		return CarIDAudi
	}

	return CarIDNissan
}

func carClassID(carName string) int {
	if carName == Audi {
		return CarClassIDAudi
	}

	return CarClassIDNissan
}

func rnd(min, max int) int {
	return Rnd.Intn(max-min) + min
}

func weightedRnd(i, points int) int {
	const (
		MaxPoints = 300
		Divisor   = 8
		Window    = 3
	)

	spread := (MaxPoints - points) / Divisor
	max := MaxDrivers - (i / Window)
	min := max - spread

	return rnd(min, max)
}
