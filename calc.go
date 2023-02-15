package carbon_calc

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
	// "gonum.org/v1/gonum/stat/distuv"
)

func CarbonPerTree(carbonFraction float64, radius float64, height float64, formFactor float64, density float64, biomassExpansionFactor float64, rootShootRatio float64) float64 {
	if carbonFraction == 0 {
		carbonFraction = DefaultCarbonFractionofTreeBiomass
	}
	if formFactor == 0 {
		formFactor = DefaultFormFactorOfTheTree
	}
	return RationCarbonPerTree * carbonFraction * CircleArea(radius) * height * formFactor * CoefficientCarbonPerTree * (1 + rootShootRatio)
}

func CarbonStoredInPlot(sumOfCarbon float64, area float64) float64 {
	return sumOfCarbon / area
}

func CarbonStoredInEachMonitoringZone(sumOfCarbonPlot float64, numPlots int64, area float64) float64 {
	return (sumOfCarbonPlot / float64(numPlots)) * area
}

// si^2_i
func VarianceOfTreeBiomass(carbonStoredPlots []float64) float64 {
	n := float64(len(carbonStoredPlots))
	var sumCorbonStoredPlots float64 = 0
	var sumCorbonStoredPlotsSqrt float64 = 0
	for _, value := range carbonStoredPlots {
		sumCorbonStoredPlots += value
		sumCorbonStoredPlotsSqrt += math.Pow(value, 2)
	}
	return (n*sumCorbonStoredPlotsSqrt - math.Pow(sumCorbonStoredPlots, 2)) / (n * (n - 1))
}

// t_delta
func TDistribution(valueOfFreedom float64) float64 {
	dist1 := distuv.StudentsT{0, 1, valueOfFreedom, nil}
	return dist1.Quantile(0.95)
}

func UncertaintyCarbonStored(tDelta float64, variance float64, area float64, sumCorbonStoredPlots float64, plotAreas []float64) float64 {
	var sumA, sumB, b float64
	nPlot := float64(len(plotAreas))
	b = sumCorbonStoredPlots / nPlot
	for _, plot := range plotAreas {
		sumA += math.Pow(plot/area, 2) * variance / nPlot
		sumB += plot / area * b
	}
	return (tDelta * math.Sqrt(sumA)) / sumB
}

func UncertaintyDiscount(uncertainty float64) float64 {
	if uncertainty <= 0.1 {
		return 0
	} else if 0.1 < uncertainty && uncertainty <= 0.15 {
		return 0.25
	} else if 0.15 < uncertainty && uncertainty <= 0.2 {
		return 0.5
	} else if 0.2 < uncertainty && uncertainty <= 0.3 {
		return 0.75
	} else {
		return 1
	}
}

func ConservativeTotalCarbon(totalCarbon, uncertainty float64) float64 {
	return totalCarbon - (1 - uncertainty*UncertaintyDiscount(uncertainty))
}

func AboveGroundBiomass(cTotalCarbon, rootShootRatio, cfTree, totalArea float64) float64 {
	return cTotalCarbon * (12 / 44) * (1 / cfTree) * (1/1 + rootShootRatio) * (1 / totalArea)
}

func NetEmissionsRemoval(cTotalCarbon, baseline, leakeage, emissions float64) float64 {
    return cTotalCarbon - baseline - leakeage - emissions
}

func MintedOCC(netEmmisionCurrent, netEmmissionPrevious float64) float64 {
    return netEmmisionCurrent - netEmmissionPrevious
}

func OCCBufferPool(mintedOcc, bufferPercent float64) float64 {
    if bufferPercent == 0 {
        bufferPercent = 0.07
    }
    return mintedOcc * bufferPercent
}

func OCCHolders(mintedOcc, holdersPercent float64) float64 {
    if holdersPercent == 0 {
        holdersPercent = 0.08
    }
    return  mintedOcc * holdersPercent
}

func OCCMintedPerMonitoringZone (mintedOcc, carbonTotal, carbonZone float64) float64 {
    return mintedOcc * (carbonZone/ carbonTotal)
}
