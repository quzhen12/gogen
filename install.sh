echo "Building..."
go build -o gogen ./main.go
mv gogen /usr/local/bin/
echo "Installed!"
