package examples

type Strategy interface {
	execute(parameter1 int32, parameter2 int32) int32
	GetStrategyType() StrategyType
}

type StrategyType int

const (
	StrategyTypeAdd StrategyType = iota
	StrategyTypeSubtract
)

type StrategyAdd struct{}

func (s StrategyAdd) execute(parameter1 int32, parameter2 int32) int32 {
	return parameter1 + parameter2
}

type StrategySubtract struct{}

func (s StrategySubtract) execute(parameter1 int32, parameter2 int32) int32 {
	return parameter1 - parameter2
}

// Strategy with LoadStrategyParameter method will generate NewStrategy function
type TravelStrategy interface {
	execute() int32
	GetStrategyType() TravelStrategyType
	LoadStrategyParameter(distance Mile)
}

type Mile float32

type TravelStrategyType int

func (t TravelStrategyType) GetDefaultStrategy() TravelStrategy {
	return NewTravelStrategy(t, 0)
}

const (
	TravelStrategyTypeUber TravelStrategyType = iota
	TravelStrategyTypeWalk
)

type TravelStrategyUber struct {
	distance Mile
}

func (s *TravelStrategyUber) LoadStrategyParameter(distance Mile) {
	s.distance = distance
}

func (s *TravelStrategyUber) execute() int32 {
	return int32(s.distance * 1.5)
}

type TravelStrategyWalk struct {
	distance Meter
}

type Meter float32

func (s *TravelStrategyWalk) LoadStrategyParameter(distance Mile) {
	s.distance = Meter(float32(distance) * 1609.34)
}

func (s *TravelStrategyWalk) execute() int32 {
	return int32(s.distance * 1.5)
}
