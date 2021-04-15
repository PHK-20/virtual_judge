#!/bin/bash
host="localhost"
port="3306"
user="root"
password="123456"
dbname="beego_judge"

raw_sql="
    CREATE TABLE IF NOT EXISTS submit_status (
        runid INT NOT NULL,
        matchid INT,
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
    )character set = utf8;
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"


# raw_sql="DROP TABLE user_info";

# mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

raw_sql="
    CREATE TABLE  IF NOT EXISTS user_info (
        username VARCHAR(20) NOT NULL,
        password VARCHAR(32) NOT NULL,
        nickname VARCHAR(20) NOT NULL,
        register_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (username)
    )character set utf8;
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

raw_sql="
    CREATE TABLE  IF NOT EXISTS contest (
        matchid INT NOT NULL,
        title VARCHAR(32) NOT NULL,
        onwer VARCHAR(20) NOT NULL,
        descr VARCHAR(128) NOT NULL,
        problem VARCHAR(128) NOT NULL,
        problem_title VARCHAR(128) NOT NULL,
        begin_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        end_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (matchid)
    )character set utf8;
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

