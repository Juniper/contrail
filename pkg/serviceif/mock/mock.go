package serviceifmock

//go:generate mockgen -destination=gen_serviceif_mock.go -package=serviceifmock github.com/Juniper/contrail/pkg/serviceif Service
