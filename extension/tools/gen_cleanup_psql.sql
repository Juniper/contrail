-- Code generated by contrailschema tool from template sql_cleanup_psql.tmpl; DO NOT EDIT.

DROP PUBLICATION IF EXISTS "syncpub";

TRUNCATE TABLE metadata CASCADE;
TRUNCATE TABLE int_pool CASCADE;
TRUNCATE TABLE ipaddress_pool CASCADE;
TRUNCATE TABLE tenant_share_sample CASCADE;
TRUNCATE TABLE domain_share_sample CASCADE;
TRUNCATE TABLE ref_sample_tag CASCADE;
TRUNCATE TABLE sample CASCADE;
TRUNCATE TABLE tenant_share_tag CASCADE;
TRUNCATE TABLE domain_share_tag CASCADE;
TRUNCATE TABLE ref_tag_tag CASCADE;
TRUNCATE TABLE tag CASCADE;