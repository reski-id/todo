FROM golang:1.18

RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

# copy all files to /app
COPY . .

# add environment variables
ENV MYSQL_HOST=127.0.0.1
ENV MYSQL_PORT=3306
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=
ENV MYSQL_DBNAME=todoappdb

# Expose port 3030 to the host on
EXPOSE 3030

# build the binary
RUN go build -o todoapp

# start the application
CMD ["./todoapp"]
