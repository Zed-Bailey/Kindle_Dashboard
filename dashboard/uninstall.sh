
echo "uninstalling :("

# the -v in the grep command matches all lines that dont contain the string
cat /etc/crontab/root | grep -v "* * * * * /mnt/us/dashboard/start.sh" > /etc/crontab/root

# restart cron to stop job
/etc/init.d/cron restart

echo "done."