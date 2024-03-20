## Larissa
![banner](https://raw.githubusercontent.com/D3Ext/aesthetic-wallpapers/main/images/blue-black-girl.png)

## Background Project
Actually, this is the final project from the bootcamp that I participated in, but I added about observability. This time I used Prometheus as a collector of the resulting metrics, and grafana as a visualization of the obtained metrics. The reason I use Prometheus and Grafana is because these two monitoring tools are the most popular at the moment. The way it works is also not too complicated, quite easy to understand, although I was also confused when reading the traffic :u

## Installation
First you have to do the 3 steps below
```bash
# clone my project to your local
$ git clone https://github.com/rualnugrh/larissa

# go to folder
$ cd larissa

# and you download module before running program
$ go mod download
```

Next, you can run the program with the `build.sh` file
```bash
# before running you must copy .env.example to .env and edit that
$ cp .env.example .env

# and you can running file build.sh
$ chmod +x ./build.sh
$ ./build.sh
```

The two steps above are for Linux / Darwin users, for Windows users you can do this
```bash
# you can follow the initial steps above, after that
$ cp .env.example .env

# after edit .env
docker-compose up -d db
docker-compose up -d
```

## Directory
```
C:.
|   .env.example
|   .gitignore
|   .pre-commit-config.yaml
|   build.sh
|   docker-compose.yml
|   Dockerfile
|   go.mod
|   go.sum
|   LICENSE
|   README.md
|
+---.github
|   \---workflows
|           go-action.yml
|
+---api
|   |   api.go
|   |
|   \---http
|           admin.go
|           kunjungan.go
|           obat.go
|           penyakit.go
|           user.go
|
+---docs
+---infrastructure
|       grafana.yaml
|       prometheus.yml
|
+---internal
|   +---config
|   |       app.go
|   |       mongodb.go
|   |       postgresql.go
|   |
|   +---entity
|   |   +---domain
|   |   |       kunjungan.go
|   |   |       obat.go
|   |   |       penyakit.go
|   |   |       reporter.go
|   |   |       role.go
|   |   |       user.go
|   |   |
|   |   \---web
|   |           admin.go
|   |           penyakit.go
|   |           reported.go
|   |           user.go
|   |
|   +---middleware
|   |       jwt.go
|   |       validation.go
|   |
|   +---repository
|   |       kunjungan.go
|   |       obat.go
|   |       penyakit.go
|   |       reported.go
|   |       user.go
|   |
|   +---service
|   |       admin.go
|   |       kunjungan.go
|   |       obat.go
|   |       penyakit.go
|   |       user.go
|   |
|   \---util
|           error.go
|           json.go
|           response.go
|
\---pkg
        prometheus.go
```
