# asf
Repo for API Server Framework related work.

# Regenerating the code
After modyfying Protobuf definitions `*.pb.go` files need to be regenerated.
To regenerate the definitions call:
```
go generate ./...
```

The resulting files should be checked into the repository to allow library users import them as go library.
