# Xflow sample config

## How to use

1. After building, startisdfdsffng testenv and resetting mysql run the contrail backend:
   ```
   contrail run -c sample/contrail.yml
   ```
1. Then sync the `xflow-setup.yaml`:
   ```
   contrailcli sync -c sample/contrail.yml sample/xflow/setup.yaml
   ```
1. Then deploy the `xflow-deploy.yaml`:
   ```
   contrailgo deploy -c sample/xflow/deploy.yaml
   ```

The resulting instances file should be in ` /var/tmp/contrail_cluster/49c19da8-6d62-45c0-9a60-20855d228628/instances.yml`
