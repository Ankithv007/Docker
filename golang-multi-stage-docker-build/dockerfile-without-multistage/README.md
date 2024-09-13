# without multi-stage and distroless simple calculator project on go language 
### go language is statically typed langauge it does mot require any run time env so we choose this to build 
### so it does not have run time env so it dont want any maximum dependecies so it create an image so light weight ( in the  mb  sence )


```
calculator-project/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

### Build the Docker image:
```
docker build -t calculator-app .
```
### image size
```
docker images
```

```
REPOSITORY                        TAG       IMAGE ID       CREATED          SIZE
calculator-app                    latest    f112170fa9e0   4 minutes ago    873MB
<none>                            <none>    16f66d94ac85   8 minutes ago    873MB
<none>                            <none>    7732f15380e9   31 minutes ago   22.6MB
<none>                            <none>    5a1fe18bf18a   31 minutes ago   873MB
golang                            1.20      d5beeac3653f   7 months ago     846MB
hello-world                       latest    d2c94e258dcb   16 months ago    13.3kB
gcr.io/distroless/base-debian11   latest    2a6de77407bf   N/A              20.6MB
```
  - look it so big right  873 MB before it was 22.6 MB  that is the power of multi stage and distrless image
### Run the Docker container:
```
docker run --rm calculator-app -op add -a 5 -b 3
```
