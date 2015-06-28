go build -a -v -o goji-skeleton main.go && docker rm -f goji && docker build -t gojiskelly . && docker run -d -p 8003:80 --name goji gojiskelly 
