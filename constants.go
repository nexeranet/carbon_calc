package carbon_calc

import (
	"errors"
	"github.com/shopspring/decimal"
)

var NotEnoughHeight = errors.New("Trees should be more than 1.3 m tall to be considered in the carbon calculation.")

type ForestType string

type TreeSpecies string

type RainfallType uint8

const (
	RainfallTypeDry RainfallType = iota
	RainfallTypeMoist
	RainfallTypeWet
)

// TODO: ask about this to front-end team
// TODO: change to uint constants
const (
	ForestTypeTropicalSubtropical ForestType = "Tropical and sub-tropical"
	ForestTypeTemperate           ForestType = "Temperate"
	ForestTypeBoreal              ForestType = "Boreal"
)

const (
	TreeSpeciesConiferous                  TreeSpecies = "Coniferous"
	TreeSpeciesBroadleaf                   TreeSpecies = "Broadleaf"
	TreeSpeciesForestTundra                TreeSpecies = "Forest-tundra"
	TreeSpeciesMixedConiferousAndBroadleaf TreeSpecies = "Mixed coniferous & broadleaf"
	TreeSpeciesPines                       TreeSpecies = "Pines"
)

var DensityOverBarkOfTreesRainfall map[ForestType]map[RainfallType]float64 = map[ForestType]map[RainfallType]float64{
	ForestTypeTropicalSubtropical: {
		RainfallTypeMoist: 0.55,
		RainfallTypeDry:   0.55,
		RainfallTypeWet:   0.55,
	},
}

var DensityOverBarkOfTreesDict map[ForestType]map[TreeSpecies]float64 = map[ForestType]map[TreeSpecies]float64{
	ForestTypeTemperate: {
		TreeSpeciesConiferous: 0.45,
		TreeSpeciesBroadleaf:  0.45,
	},
	ForestTypeBoreal: {
		TreeSpeciesConiferous:                  0.45,
		TreeSpeciesBroadleaf:                   0.45,
		TreeSpeciesForestTundra:                0.45,
		TreeSpeciesMixedConiferousAndBroadleaf: 0.45,
	},
}

func DensityOverBarkOfTrees(forestType ForestType, specie TreeSpecies, rainfall RainfallType) decimal.Decimal {
	baseValue := decimal.NewFromFloat(0.55)
	if forestType == ForestTypeTropicalSubtropical {
		value, ok := DensityOverBarkOfTreesRainfall[ForestTypeTropicalSubtropical][rainfall]
		if !ok {
			return baseValue
		}
		return decimal.NewFromFloat(value)
	}
	value, ok := DensityOverBarkOfTreesDict[forestType][specie]
	if !ok {
		return baseValue
	}
	return decimal.NewFromFloat(value)
}

var BiomassExpansionFactorDict map[ForestType]map[TreeSpecies]float64 = map[ForestType]map[TreeSpecies]float64{
	ForestTypeTropicalSubtropical: {
		TreeSpeciesPines:     1.3,
		TreeSpeciesBroadleaf: 3.4,
	},
	ForestTypeTemperate: {
		TreeSpeciesConiferous: 1.3,
		TreeSpeciesBroadleaf:  1.3,
	},
	ForestTypeBoreal: {
		TreeSpeciesForestTundra:                1.3,
		TreeSpeciesMixedConiferousAndBroadleaf: 1.3,
		TreeSpeciesConiferous:                  1.35,
		TreeSpeciesBroadleaf:                   1.3,
	},
}

func BiomassExpansionFactor(forestType ForestType, specie TreeSpecies) decimal.Decimal {
	value, ok := BiomassExpansionFactorDict[forestType][specie]
	if !ok {
		return decimal.NewFromFloat(1.15)
	}
	return decimal.NewFromFloat(value)
}

var RootShootRatioForTreeRainfall map[ForestType]map[RainfallType]func(v float64) float64 = map[ForestType]map[RainfallType]func(v float64) float64{
	ForestTypeTropicalSubtropical: {
		RainfallTypeDry: func(v float64) float64 {
			if v <= 20 {
				// default value
				return 0.56
			} else {
				return 0.28
			}
		},
		RainfallTypeMoist: func(v float64) float64 {
			if v <= 125 {
				// default value
				return 0.2
			} else {
				return 0.24
			}
		},
		RainfallTypeWet: func(v float64) float64 {
			return 0.37
		},
	},
}

var RootShootRatioForTreeDict map[ForestType]map[TreeSpecies]func(v float64) float64 = map[ForestType]map[TreeSpecies]func(v float64) float64{
	ForestTypeTemperate: {
		TreeSpeciesConiferous: func(v float64) float64 {
			if v <= 50 {
				// default value
				return 0.4
			} else if v > 50 && v <= 150 {
				return 0.29
			} else {
				return 0.2
			}
		},
		TreeSpeciesBroadleaf: func(v float64) float64 {
			if v <= 75 {
				// default value
				return 0.46
			} else if v > 75 && v <= 150 {
				return 0.23
			} else {
				// TODO: ask if this value correct @TM, because else previes one is bigger then this one
				return 0.24
			}
		},
	},
}

// abovegroundBiomass - 0 if you want to get default value
func RootShootRatioForTree(forestType ForestType, species TreeSpecies, rainfall RainfallType, abovegroundBiomass float64) decimal.Decimal {
	baseValue := decimal.NewFromFloat(0.25)
	if forestType == ForestTypeTropicalSubtropical {
		calc, ok := RootShootRatioForTreeRainfall[ForestTypeTropicalSubtropical][rainfall]
		if !ok {
			return baseValue
		}
		return decimal.NewFromFloat(calc(abovegroundBiomass))
	}
	if forestType == ForestTypeBoreal {
		if abovegroundBiomass <= 75 {
			// default value
			return decimal.NewFromFloat(0.39)
		} else {
			return decimal.NewFromFloat(0.24)
		}
	}
	calc, ok := RootShootRatioForTreeDict[forestType][species]
	if !ok {
		return baseValue
	}
	return decimal.NewFromFloat(calc(abovegroundBiomass))
}

func GetRainfallType(rainfallAmount int64) RainfallType {
	if rainfallAmount >= 2000 {
		return RainfallTypeWet
	} else if rainfallAmount < 2000 && rainfallAmount > 1000 {
		return RainfallTypeMoist
	} else {
		return RainfallTypeDry
	}
}
