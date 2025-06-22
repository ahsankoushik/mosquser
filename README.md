# Mosquser

***Thanks to [jpmens](https://github.com/jpmens) for the awosome [mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug).***

## Run Or Build this project 

### Prerequisites

1. JavaScript Runtime (Nodejs) | mine -> v22.14.0
2. Go | mine -> go1.24.0
3. Docker for mosquitto (You can also run mosquitto as you like, for that see [mosquitto.conf](mosquitto/mossquitto.conf)

### Build
After cloning the repo
```bash
# build frontend
cd frontend
npm install 
npm run buid 

# build backend
cd ..
go build main.go
```

### Run
Run the server
```bash 
./main 

# speicfy the host and port
./main --host="0.0.0.0:8080"
```

### You will also need to run mosquitto server (I am using docker)
```bash
docker compose up # add -d to run in detached mode
```

### Then migrate the database
```bash
./main migrate
```

### Lastly, create a superuser
```bash
./main superuser 
# you will be prompted to enter username and password
```


