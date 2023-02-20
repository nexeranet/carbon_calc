package carbon_calc

func NetGHGEmissions(fertilizes, cO2e_Ndirectt, cO2e_Nindirectt float64) float64 {
	return fertilizes * (cO2e_Ndirectt + cO2e_Nindirectt)
}

func CO2e_Ndirectt(synthFertilizer, nContentOfSynthFertilizer, massOfOrganicFertilizer, nContentOfOrganicFertilizer, emissionFactorForNitrousOxideEmissions, globalWarmingPotential float64) float64 {
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

// func CO2e_Nindirectt(a, b float64) float64 {
// 	return Nfert_volatIT + Nfert_leachIT
// }

// func Nfert_volatIT() float64{
//     return math.Abs((M_SF_i_t *NC_SF_i_t* Frac_GASF)+(M_OF_i_t * NC_OF_i_t * Frac_GASF)) * EF_Nvolat * (44/ 22)  * GWP_N2O
// }

// func Nfert_leachIT() float64 {
// 	return (M_SFit*NC_SFit + M_OFit*NC_OFit) * Frac_Leach * EF_NLeach * (44 / 22) * GWP_N2O
// }
