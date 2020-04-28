build:
	protoc -I. --go_out==plugins=micro:. proto/vessel/vessel.proto
	docker build -t shippy-service-vessel .

run:
	docker run -p 5101:5100 -e MICRO_SERVER_ADDRESS=:5100 shippy-service-vessel