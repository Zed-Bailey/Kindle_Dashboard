#!/bin/sh
# adds cron job to execute every minute and call start.sh
echo "* * * * * /mnt/us/dashboard/start.sh" >> /etc/crontab/root
# restart cron so the new job will execute
/etc/init.d/cron restart

echo "done."
