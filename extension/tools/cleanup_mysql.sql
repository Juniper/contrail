-- Code generated by contrailschema tool from template sql_cleanup_mysql.tmpl; DO NOT EDIT.

SET FOREIGN_KEY_CHECKS=0;

TRUNCATE TABLE metadata;
TRUNCATE TABLE int_pool;
TRUNCATE TABLE ipaddress_pool;


TRUNCATE TABLE sample;
TRUNCATE TABLE tenant_share_sample;
TRUNCATE TABLE domain_share_sample;


TRUNCATE TABLE ref_sample_tag;





TRUNCATE TABLE tag;
TRUNCATE TABLE tenant_share_tag;
TRUNCATE TABLE domain_share_tag;


TRUNCATE TABLE ref_tag_tag;






SET FOREIGN_KEY_CHECKS=1;
