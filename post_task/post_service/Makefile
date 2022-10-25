CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
	
up-version:
	migrate -source file:./migrations/ -database 'postgres://coder:123@localhost:5432/postgres?sslmode=disable' up 

down-version:
	migrate -source file:./migrations/ -database 'postgres://coder:123@localhost:5432/postgres?sslmode=disable' down

create:
	migrate create -ext sql -dir migrations -seq create_test_table