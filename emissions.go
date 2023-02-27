package carbon_calc

import "math"

func BaselineInMon(manual, area, deltaTime float64) float64 {
	return manual * area * deltaTime
}

// Calculate the baseline
func Baseline(plots []float64) float64 {
	return Sum(plots)
}

// Net GHG emissions from nitrogen fertilizer in the project in year
func NetGHGEmissions(fertilizes, cO2eNdirectt, cO2eNindirectt float64) float64 {
	return fertilizes * (cO2eNdirectt + cO2eNindirectt)
}

func CO2eNdirectt(synthFertilizer, nContentOfSynthFertilizer, massOfOrganicFertilizer, nContentOfOrganicFertilizer, emissionFactorForNitrousOxideEmissions, globalWarmingPotential float64) float64 {
	if synthFertilizer == 0 {
		synthFertilizer = 0.1
	}
	if massOfOrganicFertilizer == 0 {
		massOfOrganicFertilizer = 0.1
	}
	if emissionFactorForNitrousOxideEmissions == 0 {
		emissionFactorForNitrousOxideEmissions = 0.1
	}
	if globalWarmingPotential == 0 {
		globalWarmingPotential = 265
	}
	return (synthFertilizer*nContentOfSynthFertilizer + massOfOrganicFertilizer*nContentOfOrganicFertilizer) * emissionFactorForNitrousOxideEmissions * 44 / 22 * globalWarmingPotential
}

func CO2eNindirectt(nfertVolatIT, nfertLeachIT float64) float64 {
	return nfertVolatIT + nfertLeachIT
}

func NfertVolatIT(M_SF_i_t, NC_SF_i_t, Frac_GASF, M_OF_i_t, NC_OF_i_t, EF_Nvolat, GWP_N2O float64) float64 {
	return math.Abs((M_SF_i_t*NC_SF_i_t*Frac_GASF)+(M_OF_i_t*NC_OF_i_t*Frac_GASF)) * EF_Nvolat * (44.00 / 22.00) * GWP_N2O
}

func Nfert_leachIT(M_SFit, NC_SFit, M_OFit, NC_OFit, Frac_Leach, EF_NLeach, GWP_N2O float64) float64 {
	return (M_SFit*NC_SFit + M_OFit*NC_OFit) * Frac_Leach * EF_NLeach * (44.00 / 22.00) * GWP_N2O
}

// Calculate the net emissions removal
// emissions - emissions from other sources
// cTotalCarbon - Conservative total carbon
// TODO: fill this with definitions
// baseline -
// leakeage -
func NetEmissionsRemoval(cTotalCarbon, baseline, leakeage float64, emissions ...float64) float64 {
	return cTotalCarbon*(1-leakeage) - baseline - Sum(emissions)
}
