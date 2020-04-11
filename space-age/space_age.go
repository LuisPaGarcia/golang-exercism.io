package space

// Planet represent the planet name
type Planet string

// This constants define the year in seconds based in of the planet vs the earth
// Define once and use every iteration
const earthYearInSec float64 = 31557600.0
const mercuryYearInSec, venusYearInSec float64 = (earthYearInSec * 0.2408467), (earthYearInSec * 0.61519726)
const marsYearInSec, jupiterYearInSec, saturnYearInSec float64 = (earthYearInSec * 1.8808158), (earthYearInSec * 11.862615), (earthYearInSec * 29.447498)
const uranusYearInSec, neptuneYearInSec float64 = (earthYearInSec * 84.016846), (earthYearInSec * 164.79132)

// Age returns the age of a person in some planet based in his age in seconds
func Age(ageInSeconds float64, planet Planet) float64 {
	// We use an switch to verify the planet input, and return the value in years from the planet
	switch planet {
	case "Mercury":
		return ageInSeconds / mercuryYearInSec
	case "Venus":
		return ageInSeconds / venusYearInSec
	case "Mars":
		return ageInSeconds / marsYearInSec
	case "Jupiter":
		return ageInSeconds / jupiterYearInSec
	case "Saturn":
		return ageInSeconds / saturnYearInSec
	case "Uranus":
		return ageInSeconds / uranusYearInSec
	case "Neptune":
		return ageInSeconds / neptuneYearInSec
	case "Earth":
		return ageInSeconds / earthYearInSec
	default:
		// If the planet doesn't match, a 0 will be returned
		return 0.0
	}
}
