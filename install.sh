#!/bin/bash
host="localhost"
port="3306"
user="root"
password="123456"
dbname="beego_judge"


raw_sql="DROP TABLE submit_status";

mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

raw_sql="
    CREATE TABLE submit_status (
        runid INT  PRIMARY KEY,
        username VARCHAR(20) NOT NULL,
        oj VARCHAR(20) NOT NULL,
        problemid VARCHAR(10) NOT NULL,
        result VARCHAR(20) NOT NULL,
        execute_time INT,
        memory INT,
        language VARCHAR(10) NOT NULL,
        length INT NOT NULL,
        submit_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"

raw_sql="
    INSERT INTO submit_status 
    (runid,username,oj,problemid,result,language,length) 
    values (1,'LLLLLL0420','HUD','1000','Pending','G++',100);
"
mysql -h${host} -P${port} -u${user} -p${password} -D ${dbname} -e "${raw_sql}"