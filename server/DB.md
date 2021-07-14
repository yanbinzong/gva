

表名
questions
questionId  int(2)  否 主键 问卷编号
questName varchar(100) 否   问卷名
author char char(20)    否  发布者
startime varchar(20)  是      开始时间
endtime varchar(20) 是 截止时间
number int(4) 是  题数
type varchar(20) 是   题类型


表 titleku
ID  INT 2  否 主键  序号
questionId int 2 是    问卷编号
titleId int 2 是 题目编号
content varchar 255 是  题目内容
select1 varchar   20  是 选项1
select2 varchar   20  是 选项2
select3 varchar   20  是 选项3
select4 varchar   20  是 选项4
select5 varchar   20  是 选项5
select6 varchar   20  是 选项6


5.
表名： active
active 字段名 数据类型 宽度 是否允许为空 约束 备注 
ID int 2 否 主键 操作号 
userId int 2 是  用户编号 
questionId int 2 是  问卷编号 
titleId int 2 是  题目编号 
result varchar 50 是  结果 
datetime nchar 10 是  时间  


6．表名：
Total 字段名 数据类型 宽度 是否允许为空 约束 备注 
totalId int 2 否 主键 统计编号 
totaltime archer 20 是  统计时间 
Conditions archer 50 是  统计条件 
Results archer 100 是  结果
questionId int 2 是  问卷编号
Managerid int 2 是  管理员编号  

7. 表名：relation 字段名 数据类型 宽度 是否允许为空 约束 备注 
   ID int 2 否 主键 序号 
   ManagerId int 2 是  管理员编号 
   userId int 2 是   用户编号   
   Person Manager Questions active Total relation titleku
----------------------------------------------------------------------
