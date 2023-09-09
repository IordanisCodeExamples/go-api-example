package gen

//go:generate mockgen -package servicemock -destination internal/mocks/service/service.go -source=internal/service/service.go Store
//go:generate mockgen -package transporthttpmock -destination internal/mocks/transport/http/http.go -source=internal/transport/http/http.go Service
