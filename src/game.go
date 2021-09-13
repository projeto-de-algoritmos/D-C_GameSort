package main

type Game struct {
	Name  string  `csv:"Name"`
	Year  int     `csv:"Year"`
	Sales float64 `csv:"Global_Sales"`
}

func (g *Game) GetName() string {
	return g.Name
}

func (g *Game) GetYear() int {
	return g.Year
}

func (g *Game) GetSales() float64 {
	return g.Sales
}
