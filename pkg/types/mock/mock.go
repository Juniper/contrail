package typesmock

//go:generate mockgen -destination=gen_db_service_mock.go -package=typesmock github.com/Juniper/contrail/pkg/types DBServiceInterface
