 #!/bin/bash

#   Variables
#       SHARE_DIR (/var/www)
#       SHARE_DIR_USERS (userA userB userC)

TARGET_GID=$(stat -c "%g" $SHARE_DIR)
SHARE_DIR_USERS=($SHARE_DIR_USERS)

EXISTS=$(cat /etc/group | grep $TARGET_GID | wc -l)


if [ $EXISTS == "0" ]; then

    # Create new group using target GID and add nobody user

    groupadd -g $TARGET_GID tempgroup
    for i in "${SHARE_DIR_USERS[@]}"
    do
      usermod -a -G tempgroup $i
    done
else

    # GID exists, find group name and add

    GROUP=$(getent group $TARGET_GID | cut -d: -f1)
    for i in "${SHARE_DIR_USERS[@]}"
    do
      usermod -a -G $GROUP $i
    done
fi