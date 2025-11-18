# Time-logger

### dev-env
Clone it and make sure you have the lastest ```npm``` version && ```go-1.18``` installed then run
```make run```

### prod
Run it in a docker container

```yaml
  version: "3"

  timelogger:
    build: 
      dockerfile: Dockerfile
      context: ./time-logger
      args:
        DOMAIN: logger.example.com
    container_name: timelogger
    restart: unless-stopped
    ports:
      - 8090:8090
    volumes:
      - ../my-cloud-data/time-logger:/app/pb_data
      - /client/node_modules

```
