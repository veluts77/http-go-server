package swagger

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate winpty .\swagger.bat generate server --quiet --target server --name hello-api --spec swagger.yml --exclude-main
