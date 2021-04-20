package searcher

import (
	"math"

	log "gitlab.com/seobutik-dsp/dsp-proxy-core/pkg/logger"
)

func floatsEqual(a, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) < float64EqualityThreshold
}

// TODO sort nodes
func containsStr(arr []string, v string) bool {
	if len(arr) > 5000 {
		log.Infof("long array : %d", len(arr))
	}
	if len(arr) == 0 {
		return false
	}

	for _, e := range arr {
		if e == v {
			return true
		}
	}

	return false
}
