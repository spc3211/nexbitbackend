package constants

var Portfolios = []struct {
	ID        int
    Name      string
    Allocation map[string]int
    Sectors    map[string]int
}{
    {
		ID: 1,
        Name: "Conservative",
        Allocation: map[string]int{
            "Large Cap": 70,
            "Mid Cap":   20,
            "Small Cap": 10,
        },
        Sectors: map[string]int{
            "FMCG":    25,
            "Pharma":  20,
            "IT":      15,
            "Banking": 15,
            "Energy":  15,
            "Auto":    10,
        },
    },
    {
		ID: 2,
        Name: "Moderate",
        Allocation: map[string]int{
            "Large Cap": 60,
            "Mid Cap":   30,
            "Small Cap": 10,
        },
        Sectors: map[string]int{
            "Banking": 25,
            "IT":      20,
            "FMCG":    15,
            "Pharma":  15,
            "Auto":    15,
            "Energy":  10,
        },
    },
	{
		ID: 3,
		Name: "Growth",
		Allocation: map[string]int{
			"Large Cap": 50,
			"Mid Cap":   35,
			"Small Cap": 15,
		},
		Sectors: map[string]int{
			"IT": 30,
			"Banking": 25,
			"Auto": 15,
			"FMCG": 10,
			"Pharma": 10,
			"Energy": 10,
		},
	},
	{
		ID: 4,
		Name: "Aggressive",
		Allocation: map[string]int{
			"Large Cap": 40,
			"Mid Cap":   40,
			"Small Cap": 20,
		},
		Sectors: map[string]int{
			"IT":      35,
			"Banking": 30,
			"Auto":    15,
			"Energy":  10,
			"FMCG":    5,
			"Pharma":  5,
		},
	},
}
