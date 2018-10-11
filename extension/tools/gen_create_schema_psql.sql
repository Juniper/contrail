-- Code generated by contrailschema tool from template sql_psql.tmpl; DO NOT EDIT.

create publication "syncpub" for all tables; -- publication that is watching all tables for sync

create table metadata (
    "uuid" varchar(255),
    "type" varchar(255),
    "fq_name" varchar(255),
    unique("type", "fq_name"),
    primary key ("uuid"));

create index fq_name_index on metadata ("fq_name");

create table int_pool (
    "key" varchar(255),
    "start" int,
    "end" int
);

alter table int_pool replica identity full;

create table ipaddress_pool (
    "key" varchar(255),
    "start" inet,
    "end" inet
);

alter table ipaddress_pool replica identity full;

create table kv_store (
	"key" varchar(255),
	"value" varchar(255),
	primary key ("key")
);




create table "sample" (
    "uuid" varchar(255),
    "share" json,
    "owner_access" bigint,
    "owner" varchar(255),
    "global_access" bigint,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "name" varchar(255),
    "layout_config" varchar(255),
    "uuid_mslong" bigint,
    "uuid_lslong" bigint,
    "user_visible" bool,
    "permissions_owner_access" bigint,
    "permissions_owner" varchar(255),
    "other_access" bigint,
    "group_access" bigint,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "content_config" varchar(255),
    "container_config" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index sample_parent_uuid_index on "sample" ("parent_uuid");


create table "tag" (
    "uuid" varchar(255),
    "tag_value" varchar(255),
    "share" json,
    "owner_access" bigint,
    "owner" varchar(255),
    "global_access" bigint,
    "parent_uuid" varchar(255),
    "parent_type" varchar(255),
    "name" varchar(255),
    "uuid_mslong" bigint,
    "uuid_lslong" bigint,
    "user_visible" bool,
    "permissions_owner_access" bigint,
    "permissions_owner" varchar(255),
    "other_access" bigint,
    "group_access" bigint,
    "group" varchar(255),
    "last_modified" varchar(255),
    "enable" bool,
    "description" varchar(255),
    "creator" varchar(255),
    "created" varchar(255),
    "fq_name" json,
    "display_name" varchar(255),
    "configuration_version" bigint,
    "key_value_pair" json,
     primary key("uuid"));

create index tag_parent_uuid_index on "tag" ("parent_uuid");







create table ref_sample_tag (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    foreign key ("from") references "sample"(uuid) on delete cascade,
    foreign key ("to") references "tag"(uuid));

create index index_ref_sample_tag on ref_sample_tag ("from");




create table tenant_share_sample (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "sample"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_sample_id on tenant_share_sample("uuid");
create index index_t_sample_to on tenant_share_sample("to");

create table domain_share_sample (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "sample"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_sample_id on domain_share_sample("uuid");
create index index_d_sample_to on domain_share_sample("to");




create table ref_tag_tag (
    "from" varchar(255),
    "to" varchar(255),
    primary key ("from","to"),
    foreign key ("from") references "tag"(uuid) on delete cascade,
    foreign key ("to") references "tag"(uuid));

create index index_ref_tag_tag on ref_tag_tag ("from");




create table tenant_share_tag (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "tag"(uuid) on delete cascade,
    foreign key ("to") references project(uuid) on delete cascade);

create index index_t_tag_id on tenant_share_tag("uuid");
create index index_t_tag_to on tenant_share_tag("to");

create table domain_share_tag (
    "uuid" varchar(255),
    "to" varchar(255),
    primary key ("uuid","to"),
    "access" integer,
    foreign key ("uuid") references "tag"(uuid) on delete cascade,
    foreign key ("to") references domain(uuid) on delete cascade);

create index index_d_tag_id on domain_share_tag("uuid");
create index index_d_tag_to on domain_share_tag("to");


