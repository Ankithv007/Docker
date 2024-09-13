# with multi-stage and distroless simple calculator project on go language 
### go language is statically typed langauge it does mot require any run time env so we choose this to build 
### so it does not have run time env so it dont want any maximum dependecies so it create an image so light weight ( in the  mb  sence )

`
calculator-project/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
`

### Build the Docker image:
`
docker build -t calculator-app .
`
### image size
`
docker images
`
`
ubuntu@ip-172-31-13-215:~/docker/golang-multi-stage-docker-build$ docker images
REPOSITORY                        TAG       IMAGE ID       CREATED          SIZE
calculator-app                    latest    7732f15380e9   31 seconds ago   22.6MB
<none>                            <none>    5a1fe18bf18a   39 seconds ago   873MB
golang                            1.20      d5beeac3653f   7 months ago     846MB
hello-world                       latest    d2c94e258dcb   16 months ago    13.3kB
gcr.io/distroless/base-debian11   latest    2a6de77407bf   N/A              20.6MB
`

  - look it is only in 22.6 MB so light weight and secure right with the help of multi stage and distroless
  - if if build the application in java or python or in django just the some more MB will like it may goes up to 300 too 400 MB that`s it
  - if we build the java or python project  without multi stage and distroless it goes up to GB so it help to reduces the size of the images stage by stage
### Run the Docker container:
`
docker run --rm calculator-app -op add -a 5 -b 3
`
