# 概要
migrationツールにgooseを使用している。  
検証前にdocker-composeでdbコンテナを起動し、appデータベースを作成する必要がある。

    docker-compose exec db bash
    psql -U postgres
    create database app;
    \l

# gooseでのmigrationファイルの作成
以下コマンドでmigrationファイルを生成する。

    goose create CreateTableAdmin sql

dbフォルダ内に新たに、migrationファイルが生成されるので編集する。

