# Kindle Dashboard

This application has only been test on the kindle 4 nt(no touch). The readme also assumes that you are using a kindle 4 nt as well


download kindle 4 jail break from https://www.mobileread.com/forums/showthread.php?t=225030
follow instructions in read me

**after successful jailbreak**
download usbnetworkhack from same link
copy the .bin file for kindle 4 into kindle root directory

follow instructions to install the hack and login into kindle root

once logged into kindle
you can view some hardware information about your kindle with `cat /proc/cpuinfo`

# Building application
vector template was created using online vector tool called 'vectr'
**on your computer not kindle**
make sure you have go installed, currently the project is using go 1.19
clone this repo
create config.json file in the `dashboard` folder
and enter the following keys
```json
{
	"lat" : <your latitude>,
	"lon" : <your longitude>,
	"openweatherApiKey" : "<key>",
	"name" : "<optional>"
}
```

then run the build.sh script

# Installation

to stop the kindle from going to sleep you can open the keyboard and type `~ds`.
to turn auto sleep back on, open the keyboard again and type `~ds`


**on computer**
run the bundle script in the project root directory to bundle the project into a zip
then copy zip to kindle: `scp bundle.zip root@{kindle ip}:/mnt/us/`

**on kindle**
cd into `/mnt/us`
`unzip bundle.zip`
cd into dashboard
`./start.sh &` will start the script and detach from it so that it will run in the background


that's it. You can put your kindle in a photo frame to hide the body or build a small stand for it

# resources
for rsvg-convert
https://www.mobileread.com/forums/showthread.php?t=200621&page=3