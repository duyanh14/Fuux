docker build -t file-server .
docker run -p 2210:2210 --name file-server -v E:\file-server:/app/data --restart=always file-server