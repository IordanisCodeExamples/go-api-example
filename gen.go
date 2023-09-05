package gen

//go:generate mockgen -package servicemock -destination internal/mocks/service/service.go -source=internal/service/service.go Store
