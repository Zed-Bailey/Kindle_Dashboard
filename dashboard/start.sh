#!/bin/sh

while true;
do
    # run executable
    cd "bin"

    # run binary, generates svg
    ./dashboard


    # Check if svg exists and if it does convert it to PNG and show on screen
    if [ -e ../svg/temp.svg ]; then
        cd "../rsvg-convert/bin"
        export LD_LIBRARY_PATH=../lib:../kindle-lib:/lib:/usr/lib
        ./rsvg-convert --background-color=white /mnt/us/dashboard/svg/temp.svg -o /mnt/us/dashboard/svg/temp.png
        
        # clear the screen
        eips -c
        cd "/mnt/us/dashboard/svg"
        eips -c

        fbink -g file=temp.png > /dev/null 2>&1
        rm -f temp.svg temp.png
        
    fi

    cd "/mnt/us/dashboard/"
    sleep 60
done


# sleep 5

# # https://github.com/4dcu-be/kual-dashboard/blob/master/dashboard/bin/start.sh
# echo "" > /sys/class/rtc/rtc1/wakealarm
# # Following line contains sleep time in seconds
# # sleep for 55 seconds
# echo "+55" > /sys/class/rtc/rtc1/wakealarm
# # Following line will put device into deep sleep until the alarm above is triggered
# echo mem > /sys/power/state

# # Kill self and spawn a new instance
# /bin/sh ./start.sh && exit

