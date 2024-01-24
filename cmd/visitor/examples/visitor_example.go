package examples

type Target1 struct {
	Name  string
	Value int
}

type Target2 struct {
	Name  string
	Value int
}

// type ITarget1Visitor interface {
// 	visitForTarget1(t *Target1)
// }

// func (t *Target1) accept(v ITarget1Visitor) {
// 	v.visitForTarget1(t)
// }

// type DefaultTarget1Visitor struct{}

// func (v *DefaultTarget1Visitor) visitForTarget1(t *Target1) {
// 	// do nothing
// }
