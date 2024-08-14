CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

create-mig:
	migrate create -ext sql -dir migrations -seq create_users

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table