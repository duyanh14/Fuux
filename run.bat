docker build -t duyanh14/universal:file-server .
docker run -p 2210:2210 --name File-server -v D:\File-server:/app/data --restart=always duyanh14/universal:file-server