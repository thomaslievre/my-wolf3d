package utils

const (
	MAXSAMPLES int = 100
)

var (
	tickindex int = 0
	ticksum   int = 0
	ticklist  [MAXSAMPLES]int
)

/* need to zero out the ticklist array before starting */
/* average will ramp up until the buffer is full */
/* returns average ticks per frame over the MAXSAMPLES last frames */

func GetFrameRate(newtick int) float64 {
	ticksum -= ticklist[tickindex]
	ticksum += newtick
	ticklist[tickindex] = newtick
	tickindex = (tickindex + 1) % MAXSAMPLES
	/* return average */
	return float64(ticksum / MAXSAMPLES)
}
