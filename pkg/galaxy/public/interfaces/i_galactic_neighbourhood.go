package interfaces

type StellarNeighbourhoodInterface interface {
	GetRadius() float64
	GetSystems() []StarSystemInterface
	PopulateNeighbourhood()
}
