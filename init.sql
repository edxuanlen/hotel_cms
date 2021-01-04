/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2020/12/30 11:11:24                          */
/*==============================================================*/

use hotel_cms;

drop table if exists account_book;

drop table if exists activity;

drop index person_id_date_index on attendance_record;

drop table if exists attendance_record;

drop index phone_black_list_index on black_list;

drop table if exists black_list;

drop index leader_commodify_index on commodity_inventory;

drop table if exists commodity_inventory;

drop index leader_partner_index on partner;

drop table if exists partner;

drop index phone_person_id_index on user;

drop table if exists user;

drop table if exists room;

drop table if exists room_card;

drop index phone_and_date_room_order_inedx on room_order;

drop table if exists room_order;


show tables ;
/*==============================================================*/
/* Table: account_book                                          */
/*==============================================================*/
create table account_book
(
   id                   int not null auto_increment,
   amount               int not null,
   type                 tinyint not null comment '开销收入类型
            0  房间收入
            1  水电开销
            2  日用品开销
            3  人员开销
            4  其他收入
            5  其他支出
            ',
   leader               varchar(20) not null comment '负责人',
   leader_id            int not null comment '登记负责人',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id, create_time)
)charset=UTF8
 partition by range (TO_DAYS(create_time))
   (partition p0 values less than (TO_DAYS('2020-12-01')),
   partition p1 VALUES less than (TO_DAYS('2021-01-01')),
   partition p2 VALUES less than (TO_DAYS('2021-02-01')),
   partition p3 VALUES less than (TO_DAYS('2021-03-01')),
   partition p4 VALUES less than (TO_DAYS('2021-04-01')),
   partition p5 VALUES less than (TO_DAYS('2021-05-01')),
   partition p6 VALUES less than (TO_DAYS('2021-06-01')),
   partition p7 VALUES less than (TO_DAYS('2021-07-01')),
   partition p8 VALUES less than (TO_DAYS('2021-08-01')),
   partition p9 VALUES less than (TO_DAYS('2021-09-01')),
   partition p10 VALUES less than (TO_DAYS('2021-10-01')),
   partition p11 VALUES less than (TO_DAYS('2021-11-01')),
   partition p12 VALUES less than (TO_DAYS('2021-12-01')),
   partition p13 VALUES less than (TO_DAYS('2022-01-01')),
   partition p14 VALUES less than (TO_DAYS('2022-02-01')),
   partition p15 VALUES less than (TO_DAYS('2022-03-01'))
   );

alter table account_book comment '账本，流水账，按月分表';

/*==============================================================*/
/* Table: activity                                              */
/*==============================================================*/
create table activity
(
   id                   tinyint not null auto_increment comment '活动id',
   rate                 decimal not null comment '折扣或涨价幅度',
   begin_time           datetime not null comment '活动开始时间',
   end_time             datetime not null comment '活动结束时间',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table activity comment '活动';

/*==============================================================*/
/* Table: attendance_record                                     */
/*==============================================================*/
create table attendance_record
(
   id                   int not null auto_increment comment '考勤记录id
            ',
   classes              tinyint comment '班次
            0 白天
            1 夜班',
   status               tinyint,
   clock_in             datetime comment '打卡上班时间',
   clock_out            datetime,
   date                 int comment '对应的这个月的第几天',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id, create_time)
)charset=UTF8
 partition by range (TO_DAYS(create_time))
   (partition p0 values less than (TO_DAYS('2020-12-01')),
   partition p1 VALUES less than (TO_DAYS('2021-01-01')),
   partition p2 VALUES less than (TO_DAYS('2021-02-01')),
   partition p3 VALUES less than (TO_DAYS('2021-03-01')),
   partition p4 VALUES less than (TO_DAYS('2021-04-01')),
   partition p5 VALUES less than (TO_DAYS('2021-05-01')),
   partition p6 VALUES less than (TO_DAYS('2021-06-01')),
   partition p7 VALUES less than (TO_DAYS('2021-07-01')),
   partition p8 VALUES less than (TO_DAYS('2021-08-01')),
   partition p9 VALUES less than (TO_DAYS('2021-09-01')),
   partition p10 VALUES less than (TO_DAYS('2021-10-01')),
   partition p11 VALUES less than (TO_DAYS('2021-11-01')),
   partition p12 VALUES less than (TO_DAYS('2021-12-01')),
   partition p13 VALUES less than (TO_DAYS('2022-01-01')),
   partition p14 VALUES less than (TO_DAYS('2022-02-01')),
   partition p15 VALUES less than (TO_DAYS('2022-03-01'))
   );

alter table attendance_record comment '考勤记录表，按月分表';

/*==============================================================*/
/* Index: person_id_date_index                                  */
/*==============================================================*/
create index person_id_date_index on attendance_record
(
    id,
    date
);

/*==============================================================*/
/* Table: black_list                                            */
/*==============================================================*/
create table black_list
(
   id                   int not null,
   name                 varchar(20) not null comment '客户姓名',
   phone                varchar(20) not null comment '客户用于预定的手机号',
   cause                varchar(120) comment '原因',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table black_list comment '黑名单';

/*==============================================================*/
/* Index: phone_black_list_index                                */
/*==============================================================*/
create unique index phone_black_list_index on black_list
(
   phone
);

/*==============================================================*/
/* Table: commodity_inventory                                   */
/*==============================================================*/
create table commodity_inventory
(
   id                   int not null auto_increment comment '商品id',
   commodity_name       varchar(20) not null comment '商品名',
   stock                int not null comment '库存',
   leader               varchar(20) not null comment '负责人',
   leader_id            int not null comment '负责人的工号id',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table commodity_inventory comment '日用品库存';

/*==============================================================*/
/* Index: leader_commodify_index                                */
/*==============================================================*/
create index leader_commodify_index on commodity_inventory
(
   leader_id
);

/*==============================================================*/
/* Table: partner                                               */
/*==============================================================*/
create table partner
(
   id                   int not null auto_increment,
   partner_name         varchar(20) not null comment '合作伙伴名称',
   discount             decimal not null comment '折扣力度',
   status               tinyint not null comment '目前状态
            0 不合作
            1 合作',
   leader               varchar(20) not null,
   leader_id            int not null,
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table partner comment '合作伙伴、旅行社等';

/*==============================================================*/
/* Index: leader_partner_index                                  */
/*==============================================================*/
create index leader_partner_index on partner
(
   leader_id
);

/*==============================================================*/
/* Table: user                                             */
/*==============================================================*/
create table user
(
   id                   int not null auto_increment comment '人事id',
   identity_card        varchar(20) not null comment '身份证
            ',
   name                 varchar(20) not null comment '姓名',
   phone                varchar(20) not null comment '手机号',
   identity             int not null comment '身份
            00 总经理 01 副总经理
            10 财务部总监 11 收银员  12 采货员
            20 人力资源总监 21 HR
            30 市场营销总监 31 市场营销员工
            40 大堂经理  41 前台
            50 行政管家  51 分部门主管  52 客房服务员
            60 工程总监  61 技工
            70 保安主管  71 保安
            ',
   password             varchar(64) not null comment '加密密码',
   salary               int comment '月薪',
   status               tinyint not null default 1 comment '在职状态
            0 离职
            1 在职',
   access               tinyint comment '权利
            0 所有权限
            1 财务（薪资、账本、汇总）
            2 仓库管理
            3 外交（合作伙伴）
            4 房间管理（前台、清理人员）
            ',
   entry_time           datetime not null comment '入职时间',
   last_payday          datetime not null comment '上一次发薪日期',
   create_time          datetime not null comment '创建时间',
   modify_time          datetime not null comment '更新时间',
   primary key (id)
)charset=UTF8;

alter table user comment '人事管理表';

/*==============================================================*/
/* Index: phone_person_id_index                                 */
/*==============================================================*/
create unique index phone_person_id_index on user
(
   phone
);

/*==============================================================*/
/* Table: room                                                  */
/*==============================================================*/
create table room
(
   id                   int not null auto_increment comment '房间id',
   type                 tinyint not null comment '房间类型
            0 单人房
            1 标间
            2 大床
            3 亲子房
            4 套房 ',
   price                int not null comment '房间当天价格',
   number               int not null default 0 comment '房间数量',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table room comment '房间管理';

/*==============================================================*/
/* Table: room_card                                             */
/*==============================================================*/
create table room_card
(
   id                   int not null auto_increment comment '房卡id
            ',
   room_id              int comment '房间id',
   status               tinyint not null default 0 comment '房卡状态
            0 未绑定
            1 使用中
            2 已遗失',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table room_card comment '房卡管理';

/*==============================================================*/
/* Table: room_order                                            */
/*==============================================================*/
create table room_order
(
   id                   int not null comment '订单id，按月分表',
   phone                varchar(20) not null comment '用户用于预定的手机号',
   name                 varchar(20) not null comment '用户姓名',
   begin_time           datetime not null comment '预计开房时间',
   end_time             datetime not null comment '退房时间',
   date                 int not null comment '预定这个月的第几天，用于联合索引',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id, create_time)
)charset=UTF8
 partition by range (TO_DAYS(create_time))
   (partition p0 values less than (TO_DAYS('2020-12-01')),
   partition p1 VALUES less than (TO_DAYS('2021-01-01')),
   partition p2 VALUES less than (TO_DAYS('2021-02-01')),
   partition p3 VALUES less than (TO_DAYS('2021-03-01')),
   partition p4 VALUES less than (TO_DAYS('2021-04-01')),
   partition p5 VALUES less than (TO_DAYS('2021-05-01')),
   partition p6 VALUES less than (TO_DAYS('2021-06-01')),
   partition p7 VALUES less than (TO_DAYS('2021-07-01')),
   partition p8 VALUES less than (TO_DAYS('2021-08-01')),
   partition p9 VALUES less than (TO_DAYS('2021-09-01')),
   partition p10 VALUES less than (TO_DAYS('2021-10-01')),
   partition p11 VALUES less than (TO_DAYS('2021-11-01')),
   partition p12 VALUES less than (TO_DAYS('2021-12-01')),
   partition p13 VALUES less than (TO_DAYS('2022-01-01')),
   partition p14 VALUES less than (TO_DAYS('2022-02-01')),
   partition p15 VALUES less than (TO_DAYS('2022-03-01'))
   );

alter table room_order comment '房间订单，按月分表';

/*==============================================================*/
/* Index: phone_and_date_room_order_inedx                       */
/*==============================================================*/
create index phone_and_date_room_order_inedx on room_order
(
   phone,
   name
);

show tables ;