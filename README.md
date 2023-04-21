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