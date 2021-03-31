#!/bin/bash
host="localhost"
port="3306"
user="root"
password="123456"
dbname="beego_judge"

raw_sql="
    CREATE TABLE IF NOT EXISTS submit_status (
        runid INT NOT NULL,
        remote_runid INT NOT NULL,
        username VARCHAR(20) NOT NULL,
        oj VARCHAR(20) NOT NULL,
        problemid VARCHAR(10) NOT NULL,
        result VARCHAR(20) NOT NULL,
        result_code INT NOT NULL,
        execute_time VARCHAR(20),
        memory VARCHAR(20),
        language VARCHAR(10) NOT NULL,
        length INT NOT NULL,
        submit_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (runid)
    );
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"


raw_sql="DROP TABLE user_info";

mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

raw_sql="
    CREATE TABLE  IF NOT EXISTS user_info (
        username VARCHAR(20) NOT NULL,
        password VARCHAR(32) NOT NULL,
        register_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (username)
    );
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

