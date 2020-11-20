package structs

type Students struct {
	First, Last string
	Grade       float64
	Majors      []string
}

func Student1() Students {
	return Students{"Turd", "Ferguson", 3.3, []string{"Poetry"}}
}

func Student2() Students {
	return Students{"Chloe", "Costanza", 3.8, []string{"Womens Studies", "Hostile Business Takeover and Franchise Dissolution"}}
}
