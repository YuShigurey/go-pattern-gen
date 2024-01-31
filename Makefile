test:
	go run cmd/visitor/gen.go -name Test -type Target1,Target2 ./cmd/visitor/examples


test1:
	go run cmd/strategy/gen.go ./cmd/strategy/examples
