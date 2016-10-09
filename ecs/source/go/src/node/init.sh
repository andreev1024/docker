NAME="deploy"
ID="2000"

# Add new user
groupadd -g $ID $NAME && useradd -m -g $NAME -s /bin/bash -u $ID $NAME

# Copy ssh keys
cp -R /var/docker/.ssh /home/$NAME/.ssh
chown -R $NAME:$NAME /home/$NAME/.ssh

# Build app
sudo --user=$NAME --login sh -c "cd /var/www/vvc-user-frontend && npm install && bower install --config.interactive=false && ember build --environment=production"

exec /sbin/my_init