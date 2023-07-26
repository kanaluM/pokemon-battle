FROM golang:1.19
WORKDIR /app
RUN ["go", "mod", "init", "main"]
COPY go.mod ./
RUN ["go", "mod", "download"]
CMD ["go", "run", "."]
COPY . .
RUN ["go", "build", "-o", "/main"]
CMD ["./main"]