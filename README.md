# Golang Embed

## Pengenalan Golang Embed
Golang embed fitur yang dapat mempermudah membaca isi file pada saat compile time secara otomatis dimasukan isi filenya dalam variable
Package ini ada pada golang versi 1.16 keatas
### Cara Embed File
Untuk melakukan embed file ke variabel kita bisa mengimport package embed terlebih dahulu
Selanjutnya kita bisa menambahkan komentar `//go:embed` diikuti dengan nama file diatas variable yang kita tuju
Variable yang dituju tidak boleh disimpan di dalam function

# Embed File Ke String
Embed file bisa kita lakukan ke variable dengan type data string
Secara otomatis isi file akan dibaca sebagai teks dan memasukkannya ke dalam variable tersebut

# Embed File Ke []Byte
Type data ini sangat cocok untuk jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain

# Embed Multiple Files
Untuk melakukan embed beberapa file sekaligus kita bisa dengan komentar `//go:embed` lebih dari satu baris 
Selain itu variablenya yang bisa kita gunakan type data embed.FS

# Patch Matcher
Selain manual satu persatu seperti sebelumnya kita juga bisa menggunakan patch matcher untuk membaca file yang kita inginkan
Ini cocok ketika misalnya kita punya pola jenis file yang kita inginkan untuk dibaca
Caranya , Kita perlu menggunakan patch matcher seperti pada package function `path.Match`

# Hasil Embed Di Compile
Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanen dan data file yang dibaca disimpan dalam binari file golangnya
Artinya bukan dilakukan secara realtime membaca file yang ada diluar
Hal ini menjadikan jika binari file golang sudah di compile, kita tidak butuh lagi file externalnya, dan bahkan jika diubah file externalnya isi variabelnya tidak akan berubah