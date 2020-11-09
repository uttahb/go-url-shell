add the jobs.service file to /etc/systemd/system/jobs.service
then do 

sudo systemctl start jobs  

sudo systemctl enable jobs


then check status of service with 

sudo service jobs status

restart with

sudo service jobs restart

tail logs with

sudo journalctl -u jobs.service -f
