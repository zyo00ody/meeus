package apsis_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/apsis"
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
)

func ExampleMeanApogee() {
	// Example 50.a, p. 357.
	fmt.Printf("JDE = %.4f\n", apsis.MeanApogee(1988.75))
	// Output:
	// JDE = 2447442.8191
}

func ExampleApogee() {
	// Example 50.a, p. 357.
	j := apsis.Apogee(1988.75)
	fmt.Printf("JDE = %.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	d, f := math.Modf(d)
	fmt.Printf("%d %s %d, at %.62d TD\n", y, time.Month(m), int(d),
		base.NewFmtTime(f*24*3600))
	// Output:
	// JDE = 2447442.3543
	// 1988 October 7, at 20ʰ30ᵐ TD
}

func ExampleApogeeParallax() {
	// Example 50.a, p. 357.
	p := apsis.ApogeeParallax(1988.75)
	fmt.Printf("%.3f\n", p*180/math.Pi*3600)
	fmt.Printf("%0.3d\n", base.NewFmtAngle(p))
	// Output:
	// 3240.679
	// 54′00″.679
}
