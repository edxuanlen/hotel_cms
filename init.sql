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
   type                 tinyint not null comment '������������
            0  ��������
            1  ˮ�翪��
            2  ����Ʒ����
            3  ��Ա����
            4  ��������
            5  ����֧��
            ',
   leader               varchar(20) not null comment '������',
   leader_id            int not null comment '�ǼǸ�����',
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

alter table account_book comment '�˱�����ˮ�ˣ����·ֱ�';

/*==============================================================*/
/* Table: activity                                              */
/*==============================================================*/
create table activity
(
   id                   tinyint not null auto_increment comment '�id',
   rate                 decimal not null comment '�ۿۻ��Ǽ۷���',
   begin_time           datetime not null comment '���ʼʱ��',
   end_time             datetime not null comment '�����ʱ��',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table activity comment '�';

/*==============================================================*/
/* Table: attendance_record                                     */
/*==============================================================*/
create table attendance_record
(
   id                   int not null auto_increment comment '���ڼ�¼id
            ',
   classes              tinyint comment '���
            0 ����
            1 ҹ��',
   status               tinyint,
   clock_in             datetime comment '���ϰ�ʱ��',
   clock_out            datetime,
   date                 int comment '��Ӧ������µĵڼ���',
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

alter table attendance_record comment '���ڼ�¼�����·ֱ�';

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
   name                 varchar(20) not null comment '�ͻ�����',
   phone                varchar(20) not null comment '�ͻ�����Ԥ�����ֻ���',
   cause                varchar(120) comment 'ԭ��',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table black_list comment '������';

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
   id                   int not null auto_increment comment '��Ʒid',
   commodity_name       varchar(20) not null comment '��Ʒ��',
   stock                int not null comment '���',
   leader               varchar(20) not null comment '������',
   leader_id            int not null comment '�����˵Ĺ���id',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table commodity_inventory comment '����Ʒ���';

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
   partner_name         varchar(20) not null comment '�����������',
   discount             decimal not null comment '�ۿ�����',
   status               tinyint not null comment 'Ŀǰ״̬
            0 ������
            1 ����',
   leader               varchar(20) not null,
   leader_id            int not null,
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table partner comment '������顢�������';

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
   id                   int not null auto_increment comment '����id',
   identity_card        varchar(20) not null comment '���֤
            ',
   name                 varchar(20) not null comment '����',
   phone                varchar(20) not null comment '�ֻ���',
   identity             int not null comment '���
            00 �ܾ��� 01 ���ܾ���
            10 �����ܼ� 11 ����Ա  12 �ɻ�Ա
            20 ������Դ�ܼ� 21 HR
            30 �г�Ӫ���ܼ� 31 �г�Ӫ��Ա��
            40 ���þ���  41 ǰ̨
            50 �����ܼ�  51 �ֲ�������  52 �ͷ�����Ա
            60 �����ܼ�  61 ����
            70 ��������  71 ����
            ',
   password             varchar(64) not null comment '��������',
   salary               int comment '��н',
   status               tinyint not null default 1 comment '��ְ״̬
            0 ��ְ
            1 ��ְ',
   access               tinyint comment 'Ȩ��
            0 ����Ȩ��
            1 ����н�ʡ��˱������ܣ�
            2 �ֿ����
            3 �⽻��������飩
            4 �������ǰ̨��������Ա��
            ',
   entry_time           datetime not null comment '��ְʱ��',
   last_payday          datetime not null comment '��һ�η�н����',
   create_time          datetime not null comment '����ʱ��',
   modify_time          datetime not null comment '����ʱ��',
   primary key (id)
)charset=UTF8;

alter table user comment '���¹����';

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
   id                   int not null auto_increment comment '����id',
   type                 tinyint not null comment '��������
            0 ���˷�
            1 ���
            2 ��
            3 ���ӷ�
            4 �׷� ',
   price                int not null comment '���䵱��۸�',
   number               int not null default 0 comment '��������',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table room comment '�������';

/*==============================================================*/
/* Table: room_card                                             */
/*==============================================================*/
create table room_card
(
   id                   int not null auto_increment comment '����id
            ',
   room_id              int comment '����id',
   status               tinyint not null default 0 comment '����״̬
            0 δ��
            1 ʹ����
            2 ����ʧ',
   create_time          datetime not null,
   modify_time          datetime not null,
   primary key (id)
)charset=UTF8;

alter table room_card comment '��������';

/*==============================================================*/
/* Table: room_order                                            */
/*==============================================================*/
create table room_order
(
   id                   int not null comment '����id�����·ֱ�',
   phone                varchar(20) not null comment '�û�����Ԥ�����ֻ���',
   name                 varchar(20) not null comment '�û�����',
   begin_time           datetime not null comment 'Ԥ�ƿ���ʱ��',
   end_time             datetime not null comment '�˷�ʱ��',
   date                 int not null comment 'Ԥ������µĵڼ��죬������������',
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

alter table room_order comment '���䶩�������·ֱ�';

/*==============================================================*/
/* Index: phone_and_date_room_order_inedx                       */
/*==============================================================*/
create index phone_and_date_room_order_inedx on room_order
(
   phone,
   name
);

show tables ;