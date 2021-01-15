mysql -uroot -p <<EOF   
use beego_judge;
 CREATE TABLE IF NOT EXISTS `submit_status`(
     `runid`    INT UNSIGNED
 )
