  #!/bin/bash

  set -ex

  keystone-manage db_sync

  if [ -f /etc/apache2/envvars ]; then
     # Loading Apache2 ENV variables
     source /etc/apache2/envvars
  fi
  chmod 0666 /var/lib/keystone/keystone.db
  # Start Apache2
  apache2 -DFOREGROUND