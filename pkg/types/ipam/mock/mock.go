package ipammock

//go:generate mockgen -destination=gen_address_manager_mock.go -package=ipammock github.com/Juniper/contrail/pkg/types/ipam AddressManager
