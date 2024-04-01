build-media-server:
	docker build -f media-server/Dockerfile -t media-server .

media-server:
	docker run --rm --network host -it \
		-v $(PWD)/media-server/mediamtx.yml:/mediamtx.yml \
		-e MTX_PROTOCOLS=tcp \
		-e MTX_WEBRTCADDITIONALHOSTS=192.168.x.x \
		media-server

api:
	go run main.go api

lint:
	golangci-lint run -v ./...