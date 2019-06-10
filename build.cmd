set GOOS=linux
go-bindata-assetfs --pkg=assets --o=./internal/assets/bindata.go --prefix="F:/Frontend/wproducto/build/" "F:/Frontend/wproducto/build/..."
go build -o ./build/mproducto github.com/fpawel/mproducto/cmd/mproducto