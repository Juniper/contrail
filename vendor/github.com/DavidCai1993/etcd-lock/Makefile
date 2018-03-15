test:
	go test -v

cover:
	@rm -rf *.coverprofile
	go test -coverprofile=etcd-lock.coverprofile -v
	gover
	go tool cover -html=etcd-lock.coverprofile
