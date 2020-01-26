#!/bin/bash
echo USAGE:
echo "- first argument: 'up' or 'down'"
echo "- second argument: (optional) proxy"
echo

COMMAND=${1}
PROXY_UP=${2}

INIT_NODE_WAIT_TIME=30
SECONDARY_NODE_WAIT_TIME=15

if [ "$#" -lt 1 ]; then
  echo "ERROR: Specify 'up' or 'down'."
  exit 1
fi

if [ "${COMMAND}" != "up" ] && [ "${COMMAND}" != "down" ]; then
  echo "ERROR: first argument should be either 'up' or 'down'."
  exit 1
fi


function create_script {
  # first argument : script name
  # second argument: script content

  echo -e "#!/bin/bash\n${1}\n" > ${2}
}

echo "Setting COMPOSE_PROJECT_NAME in .env file..."
echo


NAME=`whoami|cut -d '.' -f 1`
PWD_MD5=`pwd|md5sum`
NAME="${NAME}.${PWD_MD5:1:6}"

grep -v COMPOSE_PROJECT_NAME .env > .env.swp
echo COMPOSE_PROJECT_NAME=${NAME} >> .env.swp
mv .env.swp .env

echo "PROJECT NAME: ${NAME}"

if [ "${COMMAND}" == "up" ]; then
  docker-compose up -d etcd_node
  sleep 5;

  docker-compose up -d node01
  echo "Waiting ${INIT_NODE_WAIT_TIME} seconds for first node to be up..."
  sleep ${INIT_NODE_WAIT_TIME};

  docker-compose up -d node02
  echo "Waiting ${SECONDARY_NODE_WAIT_TIME} seconds for second node to be up..."
  sleep ${SECONDARY_NODE_WAIT_TIME};

  docker-compose up -d node03
  echo "Waiting ${SECONDARY_NODE_WAIT_TIME} seconds for third node to be up..."
  sleep ${SECONDARY_NODE_WAIT_TIME};

  if [ "${PROXY_UP}" == "proxy" ]; then
    echo "Creating and initializing ProxySQL node..."
    docker-compose up -d proxysql_node
    sleep 2;
    PROXYSQL_CONTAINER=`docker-compose ps|grep Up|grep proxy|awk '{print $1}'`
    docker exec -it ${PROXYSQL_CONTAINER} add_cluster_nodes.sh

    echo
    echo "Use the following commands to access MySQL on the ProxySQL node:"
    echo

    echo "run_mysql_${PROXYSQL_CONTAINER}_admin"
    create_script run_mysql_${PROXYSQL_CONTAINER}_admin "docker exec -it ${PROXYSQL_CONTAINER} mysql -h127.0.0.1 -P6032 -uadmin -padmin \"\$@\""

    PROXYSQL_USER=`cat .env|grep MYSQL_PROXY_USER=|cut -d '=' -f2`
    PROXYSQL_PWD=`cat .env|grep MYSQL_PROXY_PASSWORD=|cut -d '=' -f2`
    echo "run_mysql_${PROXYSQL_CONTAINER}"
    create_script run_mysql_${PROXYSQL_CONTAINER} "docker exec -it ${PROXYSQL_CONTAINER} mysql -h127.0.0.1 -P3306 -u${PROXYSQL_USER} -p${PROXYSQL_PWD} \"\$@\""
  fi

  echo
  echo "Use the following commands to access BASH, MySQL, docker inspect and logs -f on each node:"
  echo
  for CONTAINER in `docker-compose ps|grep Up|grep -v etcd|grep -v proxysql|awk '{print $1}'`; do
    echo "run_bash_${CONTAINER}"
    create_script run_bash_${CONTAINER} "docker exec -it ${CONTAINER} bash"
    echo "run_mysql_${CONTAINER}"
    create_script run_mysql_${CONTAINER} "docker exec -it ${CONTAINER} mysql -uroot -proot \"\$@\""
    echo "run_inspect_${CONTAINER}"
    create_script run_inspect_${CONTAINER} "docker inspect ${CONTAINER}"
    echo "run_logs_${CONTAINER}"
    create_script run_logs_${CONTAINER} "docker logs -f ${CONTAINER}"
    echo
  done;

  chmod +x run_*_*
  docker-compose ps

else
  if [ "${COMMAND}" == "down" ]; then
    echo "Stopping containers and cleaning up..."
    docker-compose down

    echo "Deleting run_* scripts..."
    rm -f run_bash_* run_mysql_* run_inspect_* run_logs_*
  fi
fi

exit 0
