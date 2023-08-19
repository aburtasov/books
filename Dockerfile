FROM golang:1.17 as builder
WORKDIR /go/src/books
COPY . .
RUN go build -o /build/bin/book-app ./cmd/main.go


FROM centos
COPY --from=builder /build/bin/book-app /build/bin/book-app
COPY . .
ENV DB_PASSWORD=root
ENV DB_HOST=db:3306

CMD [ "/build/bin/book-app" ]