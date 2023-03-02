package carbon_calc

import (
	"github.com/shopspring/decimal"
)

// Calculate the baseline in monitoring zone V1
// manual - Mean change in carbon stock in trees per ha and per year → Value entered
// per monitoring zone in the admin dashboard. If no value entered when the
// carbon calculation is done in the backend (after stage is validated, take 0
// as default value) ; tCO2e ha-1 yr-1
// area - Area of baseline monitoring zone i, delineated on the basis of tree crown
// cover at the start of the A/R CDM project activity; ha
// deltaTime - time elapsed between current collateralized and validated stage and previous
// validated stage (years) - take into account the end months and years of each
// stage
func BaselineInMonitoringZone(manual, area, deltaTime decimal.Decimal) decimal.Decimal {
	return manual.Mul(area).Mul(deltaTime)
}

// Calculate the baseline
func Baseline(baselines []decimal.Decimal) decimal.Decimal {
	return SumDecimal(baselines)
}

// Net GHG emissions from nitrogen fertilizer in the project in year
func NetGHGEmissions(fertilizes, cO2eNdirectt, cO2eNindirectt decimal.Decimal) decimal.Decimal {
	return fertilizes.Mul(cO2eNdirectt.Add(cO2eNindirectt))
}

// TODO:
func CO2eNdirectt(massSynthFertz, nContSynthFertz, massOrgFertz, nContOrgFertz, nitrOxdEmissSOC, gWarmingPotentl decimal.Decimal) decimal.Decimal {
	if nContSynthFertz.Equal(decimal.Zero) {
		nContSynthFertz = massSynthFertz.Mul(decimal.NewFromFloat(0.1))
	}
	if nContOrgFertz.Equal(decimal.Zero) {
		nContOrgFertz = massOrgFertz.Mul(decimal.NewFromFloat(0.1))
	}
	if nitrOxdEmissSOC.Equal(decimal.Zero) {
		nitrOxdEmissSOC = decimal.NewFromFloat(0.01)
	}
	if gWarmingPotentl.Equal(decimal.Zero) {
		gWarmingPotentl = decimal.New(265, 0)
	}
	b := decimal.NewFromFloat(44.0 / 28.0)
	return massSynthFertz.Mul(nContSynthFertz).
		Add(massOrgFertz.
			Mul(nContOrgFertz)).
		Mul(nitrOxdEmissSOC).
		Mul(b).
		Mul(gWarmingPotentl)
}

func CO2eNdirecttDefault(massSynthFertz, massOrgFertz decimal.Decimal) decimal.Decimal {
	b := decimal.NewFromFloat(1.1 * 0.001 * (44.0 / 28.0) * 265)
	return massSynthFertz.Add(massOrgFertz).Mul(b)
}

func CO2eNindirectt(nfertVolatIT, nfertLeachIT decimal.Decimal) decimal.Decimal {
	return nfertVolatIT.Add(nfertLeachIT)
}

func NfertVolatIT(massSynthFertz, nContSynthFertz, massOrgFertz, nContOrgFertz, allFractSynth, allFractOrg, nitrOxdEmissWS, gWarmingPotentl decimal.Decimal) decimal.Decimal {
	if nContSynthFertz.Equal(decimal.Zero) {
		nContSynthFertz = massSynthFertz.Mul(decimal.NewFromFloat(0.1))
	}
	if nContOrgFertz.Equal(decimal.Zero) {
		nContOrgFertz = massOrgFertz.Mul(decimal.NewFromFloat(0.1))
	}
	if nitrOxdEmissWS.Equal(decimal.Zero) {
		nitrOxdEmissWS = decimal.NewFromFloat(0.01)
	}
	if gWarmingPotentl.Equal(decimal.Zero) {
		gWarmingPotentl = decimal.New(265, 0)
	}
	if allFractSynth.Equal(decimal.Zero) {
		allFractSynth = decimal.NewFromFloat(0.1)
	}
	if allFractOrg.Equal(decimal.Zero) {
		allFractOrg = decimal.NewFromFloat(0.3)
	}
	b := decimal.NewFromFloat(44.0 / 28.0)
	return massSynthFertz.Mul(nContSynthFertz).
		Mul(allFractSynth).
		Add(massOrgFertz.Mul(nContOrgFertz).Mul(allFractOrg)).
		Abs().
		Mul(nitrOxdEmissWS).
		Mul(b).
		Mul(gWarmingPotentl)
}

func NfertVolatITDefault(massSynthFertz, massOrgFertz decimal.Decimal) decimal.Decimal {
	return massSynthFertz.Mul(decimal.NewFromFloat(1.1 * 0.1)).
		Add(massOrgFertz.Mul(decimal.NewFromFloat(1.1 * 0.3))).
		Mul(decimal.NewFromFloat(0.01 * (44.0 / 28.0) * 265))
}

func NfertLeachIT(massSynthFertz, nContSynthFertz, massOrgFertz, nContOrgFertz, nFractSoil, nitrOxdEmissLR, gWarmingPotentl decimal.Decimal) decimal.Decimal {
	if nContSynthFertz.Equal(decimal.Zero) {
		nContSynthFertz = massSynthFertz.Mul(decimal.NewFromFloat(0.1))
	}
	if nContOrgFertz.Equal(decimal.Zero) {
		nContOrgFertz = massOrgFertz.Mul(decimal.NewFromFloat(0.1))
	}
	if gWarmingPotentl.Equal(decimal.Zero) {
		gWarmingPotentl = decimal.New(265, 0)
	}
	if nFractSoil.Equal(decimal.Zero) {
		nFractSoil = decimal.NewFromFloat(0.3)
	}
	if nitrOxdEmissLR.Equal(decimal.Zero) {
		nitrOxdEmissLR = decimal.NewFromFloat(0.0075)
	}
	b := decimal.NewFromFloat(44.0 / 28.0)
	return massSynthFertz.
		Mul(nContSynthFertz).
		Add(massOrgFertz.Mul(nContOrgFertz)).
		Mul(nFractSoil).
		Mul(nitrOxdEmissLR).
		Mul(b).
		Mul(gWarmingPotentl)
}

func NfertLeachITDefault(massSynthFertz, massOrgFertz decimal.Decimal) decimal.Decimal {
	b := decimal.NewFromFloat(1.1 * 0.3 * 0.0075 * (44.0 / 28.0) * 265)
	return massSynthFertz.Add(massOrgFertz).Mul(b)
}

// Calculate the net emissions removal
// emissions - emissions from other sources
// cTotalCarbon - Conservative total carbon
// TODO: fill this with definitions
// baseline -
// leakeage -
func NetEmissionsRemoval(cTotalCarbon, baseline, leakeage decimal.Decimal, emissions ...decimal.Decimal) decimal.Decimal {
	return cTotalCarbon.Mul(decimal.New(1, 0).Sub(leakeage)).
		Sub(baseline).
		Sub(SumDecimal(emissions))
}
