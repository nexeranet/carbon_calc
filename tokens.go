package carbon_calc

import "github.com/shopspring/decimal"

// Calculate the OCCs to be minted
// To calculate the carbon credits to mint for stage T, we compute the difference
// between the total carbon calculated for stage T and the total carbon calculated
// for the previously validated stage (T-1, T-2 or T-3).
func MintedOCC(netEmmisionCurrent, netEmmissionPrevious decimal.Decimal) decimal.Decimal {
	return netEmmisionCurrent.Sub(netEmmissionPrevious)
}

// Calculate the OCCs to be sent to the buffer pool
func OCCBufferPool(mintedOcc decimal.Decimal, percent float64) decimal.Decimal {
	if percent == 0 {
		percent = 0.07
	}
	return mintedOcc.Mul(decimal.NewFromFloat(percent))
}

// Calculate the OCCs to be sent to the token holders
func OCCHolders(mintedOcc decimal.Decimal, percent float64) decimal.Decimal {
	if percent == 0 {
		percent = 0.08
	}
	return mintedOcc.Mul(decimal.NewFromFloat(percent))
}

// Calculate the OCCs minted per monitoring zone
// We need to calculate and store (in our databases) the amount of OCCs minted per
// monitoring zone per stage to be able to burn (in the future) the equivalent
// amount of OCCs from the buffer pool if something happens in the forest.
// minted - OCCs to be minted at stage T
// carbonC - Carbon stored in all monitoring zones, current stage
// zoneC - Carbon stored in monitoring zone i, current stage
// carbonP - Carbon stored in all monitoring zones, previous stage
// zoneP - Carbon stored in monitoring zone i, previous stage
func OCCMintedPerMonitoringZone(minted, carbonC, zoneC, carbonP, zoneP decimal.Decimal) decimal.Decimal {
	return minted.Mul(zoneC.Sub(zoneP).Div(carbonC.Sub(carbonP)))
}
