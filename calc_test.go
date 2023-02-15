package carbon_calc

import (
	"testing"

	"github.com/nexeranet/carbon_calc/tdist"
)


func TestTDistribution(t *testing.T) {
    dist := tdist.TDist{1}
    t.Log("INV",dist.Inv(0.9))
    t.Log("GINV",dist.GInv(0.9))
    t.Log("TDistribution",TDistribution(1))
}
