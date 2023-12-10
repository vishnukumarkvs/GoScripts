# Based on Akhil Sharma Tutorial

- Used Go Fiber and postgres gorm
- https://youtu.be/1XPktts9USg?si=S3lz-2o5jm-u4x2w
- Postgres - Use containerd and nerdctl (in wsl) to run postgres container
- nerdctl run -e POSTGRES_USER=vishnukvs -e POSTGRES_PASSWORD=vishnukvs -e POSTGRES_DB=fiber -p 5432:5432 --name postgres-container postgres
