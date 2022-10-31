#!/bin/bash
# -------------------------------------------------------------------------- #
# Version 1.1                                                                #
# Backup arango and redis Nocloud databases                                  #
# Don't forget to chmod +x backup.sh                                         #
# cron example: 00 */1 * * * /bin/bash /root/deployment/backup.sh            #
# -------------------------------------------------------------------------- #

nocloud_deployment_path="/root/deployment/"
host_backup_path="/backups_nocloud/"                                            # Where to store backups on host; better to use remote or nfs storage

nocloud_pass_file="$nocloud_deployment_path/.env"                               # Nocloud file with db root pass
host_date=$(date "+%d-%b-%y-%Hh-%Mm-%Ss")                                       # Pretty-print date
arangodump_tmp_dir="/arango_dumps_$host_date"                                   # tmp dir to store backup in container
redis_backup_dir="redis_data_$host_date"                                        # Final redis backup location on host
arango_root_pass=$(cat $nocloud_pass_file | grep DB_PASS | cut -d\= -f2)        # arangodb root pass
arango_dump_command="/usr/bin/arangodump \
                     --all-databases \
                     --overwrite \
                     --output-directory $arangodump_tmp_dir \
                     --server.password $arango_root_pass"                       # dump execution command
arango_container_name=$(docker ps --format "{{.Names}}"| grep db)               # Find container name with arangodb
redis_container_name=$(docker ps --format "{{.Names}}"| grep redis)

[ -d $host_backup_path ] || mkdir -p $host_backup_path                          # Create backup dir on host, if not exists

docker exec -d $arango_container_name $arango_dump_command                      # Create arangodump
docker cp $arango_container_name:$arangodump_tmp_dir $host_backup_path          # Copy dump from container to host
docker exec -d $arango_container_name rm -rf $arangodump_tmp_dir                # Remove tmp dir from container

docker exec -d $redis_container_name /usr/local/bin/redis-cli SAVE && \
docker cp $redis_container_name:/data $host_backup_path/$redis_backup_dir       # Commit changes to dump.rdb and copy /data to host

find $host_backup_path/* -mtime +30 -exec rm {} \;                              # Remove backups older then 30 days
