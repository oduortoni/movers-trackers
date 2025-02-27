package datatypes

type Charges struct {
	Means string
	Load int
	Cost int
}
type UserCharges struct {
    Group string
    Means string
    From string
    To string
    Load int
    CostPerDistance int
    Distance int
    TotalCost int
}
type ChargesGroup struct{
    Success bool
    Charges UserCharges
    ExistGroup bool
}

