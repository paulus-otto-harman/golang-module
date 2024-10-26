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
gola.Input(gola.Args(gola.P("label","Masukkan teks"))) // menerima input dan mengembalikan nilai bertipe data string
gola.Input(gola.Args(gola.P("type","number"),gola.P("label","Masukkan angka"))) // menerima input dan mengembalikan nilai bertipe data int
```
gola Input() merupakan jalan pintas untuk kegiatan mengetik berulang-ulang seperti ini :
```
var teks string
fmt.Print("Masukkan teks")
fmt.Scanln(&teks)
```
ide ```gola.Input()``` berasal dari elemen HTML ```<input>```

### gola.Wait()
Menunggu user menekan Enter

contoh penggunaan :
```go
gola.Wait("Tekan Enter untuk kembali ke Menu Utama")
```
### gola.Tf()
Memformat teks

contoh penggunaan :
```go
gola.Tf(gola.Bold,"Lorem Ipsum",gola.LightBlue) // memformat teks "Lorem Ipsum" menjadi tebal (bold) dan berwarna biru terang
```