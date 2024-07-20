deploy:
	echo "Compiling....."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s" -a -installsuffix cgo -o api
	echo "Deleting unused images"
	docker system prune --force
	echo "compile docker image"
	docker build -t fathdemr/sandbox.101.icibot.net .
	echo "pushing..."
	docker push fathdemr/sandbox.101.icibot.net
	rm api