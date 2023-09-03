.PHONY: run db_run db_create db_drop migrate_new migrate_up migrate_down

# MySQL 连接
DB_URL=mysql://root:123456@tcp(localhost:3306)/my_website

# 启动 server
run:
	go run main.go

# 用 Docker 启动一个 MySQL 服务
db_run:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v docker-mysql:/var/lib/mysql -d mysql:8

# 创建数据库
db_create:
	docker exec -e MYSQL_PWD=123456 mysql mysql -u root -e 'CREATE DATABASE IF NOT EXISTS `my_website` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;'

# 删除数据库
db_drop:
	docker exec -e MYSQL_PWD=123456 mysql mysql -u root -e 'DROP DATABASE IF EXISTS `my_website`;'

# 创建一个新的 migration
migrate_new:
	migrate create -ext sql -dir db/migration -seq $(name)

# 应用所有的 migration
migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# 撤销所有的 migration
migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
	
# migrate up 1次
migrate_up1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

# migrate down 1次
migrate_down1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1