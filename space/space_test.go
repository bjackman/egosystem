package space_test

import (
	"egosystem/space"
	"fmt"
	"testing"
)

func TestDistanceString(t *testing.T) {
	for distance, want := range map[space.Distance]string{
		0:                                      "0cm",
		space.Centimeter:                       "1cm",
		10 * space.Centimeter:                  "10cm",
		101 * space.Centimeter:                 "1.01m",
		space.Meter:                            "1m",
		2*space.Meter + 15*space.Centimeter:    "2.15m",
		4445*space.Meter + 99*space.Centimeter: "4445.99m",
		4445 * space.Meter:                     "4.445km",
		9001 * space.Meter:                     "9.001km",
		5 * space.Kilometer:                    "5km",
	} {
		got := fmt.Sprintf("%v", distance)
		if got != want {
			t.Errorf("Formatting Distance(%d) as string returned %q, want %q",
				int(distance), got, want)
		}
	}
}
