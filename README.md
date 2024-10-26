# gola (Go Lumoshive Academy)
adalah Package untuk func-func yang sering digunakan selama pembelajaran bahasa pemrograman Go

## Cara download :
```
go get github.com/paulus-otto-harman/golang-module
```

## Func
### gola.ClearScreen()

digunakan untuk membersihkan layar

contoh penggunaan :
```go
gola.ClearScreen()
```

### gola.Input()

digunakan untuk input ke variabel int atau string

contoh penggunaan :
```go
gola.Input(Args(P("label","Masukkan teks")))
```
gola Input() merupakan jalan pintas untuk kegiatan mengetik berulang-ulang seperti ini :
```
var teks string
fmt.Print("Masukkan teks")
fmt.Scanln(&teks)
```

### gola.Wait()

Menunggu user menekan Enter

contoh penggunaan :
```go
gola.Wait("Tekan Enter untuk kembali ke Menu Utama")
```
