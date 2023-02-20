package carbon_calc

// Calculate the net emissions removal
// emissions - emissions from other sources
// cTotalCarbon - Conservative total carbon
// TODO: fill this with definitions
// baseline -
// leakeage -
func NetEmissionsRemoval(cTotalCarbon, baseline, leakeage, emissions float64) float64 {
	return cTotalCarbon - baseline - leakeage - emissions
}

// Calculate the OCCs to be minted
// To calculate the carbon credits to mint for stage T, we compute the difference
// between the total carbon calculated for stage T and the total carbon calculated
// for the previously validated stage (T-1, T-2 or T-3).
func MintedOCC(netEmmisionCurrent, netEmmissionPrevious float64) float64 {
	return netEmmisionCurrent - netEmmissionPrevious
}

// Calculate the OCCs to be sent to the buffer pool
func OCCBufferPool(mintedOcc, percent float64) float64 {
	if percent == 0 {
		percent = 0.07
	}
	return mintedOcc * percent
}

// Calculate the OCCs to be sent to the token holders
func OCCHolders(mintedOcc, percent float64) float64 {
	if percent == 0 {
		percent = 0.08
	}
	return mintedOcc * percent
}

// Calculate the OCCs minted per monitoring zone
// We need to calculate and store (in our databases) the amount of OCCs minted per
// monitoring zone per stage to be able to burn (in the future) the equivalent
// amount of OCCs from the buffer pool if something happens in the forest.
// minted - OCCs to be minted at stage T
// carbon - Carbon stored in all monitoring zones
// zone - Carbon stored in monitoring zone i
func OCCMintedPerMonitoringZone(minted, carbon, zone float64) float64 {
	return minted * (zone / carbon)
}
