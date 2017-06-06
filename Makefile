clean:
	go clean
	rm -f trt

trt:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .

docker: trt
	docker build ./
	
sous: trt
	sous build
