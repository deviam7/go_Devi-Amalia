# Menggunakan base image dari Golang versi 1.16
FROM golang:1.16

# Setelah membuat Dockerfile, buat direktori kerja di dalam container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy seluruh file yang dibutuhkan aplikasi ke dalam container
COPY . .

# Build aplikasi Golang di dalam container
RUN go build -o main .

# Expose port yang akan digunakan untuk mengakses aplikasi
EXPOSE 8080

# Menjalankan aplikasi di dalam container
CMD ["./main"]

