run: fe-build be-build
	rm -rf release
	mkdir -p release
	cp -r ./server/time-logger ./release/
	cp -r ./client/dist ./release/pb_public
	./release/time-logger serve --http=0.0.0.0:8090

fe-build:
	cd ./client && npm install
	cd ./client/ && VITE_APP_URL=http://localhost:8090 npm run build

be-build:
	cd ./server && go mod download
	cd ./server && CGO_ENABLED=0 go build -a -installsuffix cgo

run-dev-fe:
	cd ./client && npm install
	cd ./client/ && VITE_APP_URL=http://localhost:8090 npm run dev

run-dev-fe3:
	cd ./client3 && npm install
	cd ./client3/ && VITE_APP_URL=http://localhost:8090 npm run dev

run-dev-be: 
	cd ./server && go mod download
	cd ./server/ && go run main.go serve
