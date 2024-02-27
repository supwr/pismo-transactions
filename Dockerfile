FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cespare/reflex@latest

COPY . .

COPY reflex.conf /usr/local/etc/
COPY build.sh /usr/local/bin/

EXPOSE 8000

CMD [ "reflex", "-d", "none", "-c", "/usr/local/etc/reflex.conf" ]