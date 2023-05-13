FROM golang:1.18

RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

# copy all files to /app
COPY . .

# Expose port 3030 to the host on
EXPOSE 3030

# build the binary
RUN go build -o todoapp

# start the application
CMD ["./todoapp"]
