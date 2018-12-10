# Downloaded from https://github.com/zalando/patroni/blob/master/kubernetes/entrypoint.sh

if [[ $UID -ge 10000 ]]; then
    GID=$(id -g)
    sed -e "s/^postgres:x:[^:]*:[^:]*:/postgres:x:$UID:$GID:/" /etc/passwd > /tmp/passwd
    cat /tmp/passwd > /etc/passwd
    rm /tmp/passwd
fi

cat > /home/postgres/patroni.yml <<__EOF__
scope: contrail
bootstrap:
  dcs:
    postgresql:
      use_pg_rewind: true
  initdb:
  - auth-host: md5
  - auth-local: trust
  - encoding: UTF8
  - locale: en_US.UTF-8
  - data-checksums
  pg_hba:
  - host all all 0.0.0.0/0 md5
  - host replication replicator ${PATRONI_KUBERNETES_POD_IP}/16 md5
restapi:
  listen: 0.0.0.0:8008
  connect_address: '${PATRONI_KUBERNETES_POD_IP}:8008'
postgresql:
  listen: 0.0.0.0:5432
  connect_address: '${PATRONI_KUBERNETES_POD_IP}:5432'
  data_dir: /home/postgres/pgdata/pgroot/data
  pgpass: /tmp/pgpass
  authentication:
    superuser:
      username: 'postgres'
      password: 'contrail123'
    replication:
      username: 'replicator'
      password: 'rep-pass'
__EOF__

unset PATRONI_SUPERUSER_PASSWORD PATRONI_REPLICATION_PASSWORD
export KUBERNETES_NAMESPACE=$PATRONI_KUBERNETES_NAMESPACE
export POD_NAME=$PATRONI_NAME

exec /usr/bin/python3 /usr/local/bin/patroni /home/postgres/patroni.yml

