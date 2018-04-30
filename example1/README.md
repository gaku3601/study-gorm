# 前準備
migrationツールにgooseを使用している。  
検証前にdocker-composeでdbコンテナを起動し、appデータベースを作成する必要がある。

    docker-compose exec db bash
    psql -U postgres
    create database app;
    \l

以下コマンドでmigrationファイルを生成する。

    goose create CreateTableAdmin sql

dbフォルダ内に新たに、migrationファイルが生成されるので編集する。
編集後、goose upコマンドでdbにテーブルが作成されるので、正常に作成されたか確認しておく。

# 出力結果

    &main.Admin{ID:7, Name:"gakuko", Age:15, Isman:false}
    &main.Admin{ID:6, Name:"gakuko", Age:15, Isman:false}
    &main.Admin{ID:5, Name:"gakuko", Age:15, Isman:false}
    ----------------------
    &main.Admin{ID:1, Name:"gaku", Age:26, Isman:true}
