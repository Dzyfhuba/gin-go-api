#!/bin/sh

timestamp=$(date +"%Y%m%d%H%M%S")
mysqldump -h mysql -uubaid -pplmoknijb gin_go_api > "/backup/backup_$timestamp.sql"
