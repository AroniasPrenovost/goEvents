# specify the base image for the Go application
FROM golang:latest 
# create an /app directory within image that holds the application's source files
RUN mkdir /app
# copy everything in the root directory to our /app directory
ADD . /app
# specify that we now wish to execute any further commands inside our /app directory
WORKDIR /app
# automatically rebuild and restart the Go application if any of the Go source files change
RUN go get github.com/githubnemo/CompileDaemon
# run go build to compile the binary executable of our Go program
RUN go build -o main .
# the start command kicks off our newly created binary executable
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main