all:
	PORT=3000 go run *.go
container:
	docker build -t donwb/map:0.1 .
runc:
	docker run -p 3000:3000 donwb/map:0.1
