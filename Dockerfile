FROM node:current-alpine

RUN apk add make

WORKDIR /just-fe-build
COPY ./client3 .

ARG DOMAIN placeholder

RUN npm install
RUN VITE_APP_URL=https://${DOMAIN} npm run build


FROM golang:1.18

WORKDIR /app

COPY /server .

RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo

COPY --from=0 /just-fe-build/dist ./pb_public

EXPOSE 8090

CMD ["./time-logger", "serve", "--http=0.0.0.0:8090"]
