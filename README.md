# Project Objective:
********************
- Expose a REST API to send SMS to a Kannel (V. 1.4.4) server and Kannel server to relay to SMSC simulator.

### Author:
***********
- Chandramouli P <ReachCMP@gmail.com>

### System details and Assumptions:
***********************************
- Operating system: Ubuntu 14.04.5 LTS (X64)
- Kannel and API server private IP address: 172.31.19.70
- Kannel and API server public IP address: 35.154.95.122
- SMSC simulator server private IP address: 172.31.18.204
- SMSC simulator server public IP address: 35.154.65.222
- Programming language is used to develop the REST API: GoLang (V. 1.8)
- Software that is used as SMSC simulator: SMPPSim (V. 2.6.11) (A product of Selenium Software company)
- Home folder: /home/ubuntu
- Ports to be opened in firewall (Kannel server): 80, 3000, 13000, 13013 (All are TCP ports)
- Ports to be opened in firewall (SMSC server): 88 (TCP)
- You enjoy in working with Linux and Telecom, and also have some knowledge about internal components of Kannel like BearerBox, SMSBox, protocols, SMSC, SMPP, VoIP etc.
- Execute the commands at the command prompt. Here, I am not adding my home directory or # or $ symbol before starting of the commands. 

### Note: 
********
- "Kannel" and "REST API" (GoLang) are deployed in one server and SMSC simulator is deployed in another server. Both servers are in private network. 
- You must add your public IP address (white listing) in Kannel configuration file (/usr/local/sbin/smskannel.conf) to simulate and test the setup of this SMS service.

### Setting up the Kannel server:
*********************************
cd /home/ubuntu  
sudo apt-get update  
sudo apt-get upgrade  
sudo apt-get install build-essential wget libssl-dev libncurses5-dev libnewt-dev  libxml2-dev linux-headers-$(uname -r) libsqlite3-dev uuid-dev lynx-cur  
sudo apt-get build-dep linux-image-$(uname -r)  
sudo apt-get install kernel-package  

bison -V  
sudo apt-get remove bison  
sudo apt-get install python-dev  
sudo apt-get install python3-dev  

wget ftp://xmlsoft.org/libxml2/libxml2-2.9.4.tar.gz  
tar zxvf libxml2-2.9.4.tar.gz  
cd libxml2-2.9.4/  
./configure  
make  
sudo make install  
cd ..  

wget ftp://ftp.gnu.org/gnu/m4/m4-1.4.10.tar.gz  
tar -xvzf m4-1.4.10.tar.gz  
cd m4-1.4.10  
./configure --prefix=/usr/local/m4  
make  
sudo make install  
cd ..  

wget http://ftp.gnu.org/gnu/bison/bison-2.7.tar.gz  
tar zxvf bison-2.7.tar.gz  
cd bison-2.7/  
ls /usr/local/m4/bin/  
PATH=$PATH:/usr/local/m4/bin/  
./configure --prefix=/usr/local/bison --with-libiconv-prefix=/usr/local/libiconv/  
make  
sudo make install  
cd ..  

sudo apt-get install apache2  
	(Testing: open this page in a web browser: http://35.154.95.122)  
sudo /etc/init.d/apache2 restart  

wget http://www.kannel.org/download/1.4.4/gateway-1.4.4.tar.gz  
tar zxvf gateway-1.4.4.tar.gz  
cd gateway-1.4.4/  
./configure  
touch .depend  
make depend  
make  
make check  
sudo make install  

cd /var/log  
sudo mkdir kannel  

cd /usr/local/sbin/  
sudo cp /home/ubuntu/gateway-1.4.4/gw/smskannel.conf .  

### Configure the Kannel server (You can delete the content in existing configuration file and add the below configuration using a text editor):
*******************************
cd /usr/local/sbin  
sudo nano smskannel.conf  

group = core  
admin-port = 13000  
admin-password = Test12  
status-password = Test12  
admin-deny-ip = "*.*.*.*"  
admin-allow-ip = "127.0.0.1;49.206.51.129;35.154.95.122"  
smsbox-port = 13001  
box-deny-ip = "*.*.*.*"  
box-allow-ip = "127.0.0.1;49.206.51.129;35.154.95.122"  
log-file = "/var/log/kannel/bearerbox.log"  
log-level = 1  
access-log = "/var/log/kannel/access.log"  
dlr-storage = internal  

group = smsc  
smsc = smpp  
smsc-id = SMPPSim  
allowed-smsc-id = SMPPSim  
preferred-smsc-id = SMPPSim  
host = 172.31.18.204  
port = 2775  
transceiver-mode = yes  
smsc-username = smppclient1  
smsc-password = Test12  
system-type=default  

group = smsbox  
bearerbox-host = 127.0.0.1  
sendsms-port = 13013  
global-sender = 13013  
sendsms-chars = "0123456789 +-"  
log-file = "/var/log/kannel/smsbox.log"  
log-level = 0  
access-log = "/var/log/kannel/access.log"  

group = sendsms-user  
username = chandramouli  
password = Test12  
user-deny-ip = "*.*.*.*"  
user-allow-ip = "127.0.0.1;49.206.51.129;35.154.95.122"  

group = sms-service  
keyword = nop  
text = "You asked nothing and I did it!"  

group = sms-service  
keyword = default  
text = "No service specified"  

### Setting up the SMSC server:  
*******************************
cd /home/ubuntu  
sudo apt-get update  
sudo apt-get upgrade  
sudo apt-get install build-essential wget libssl-dev libncurses5-dev libnewt-dev  libxml2-dev linux-headers-$(uname -r) libsqlite3-dev uuid-dev  
sudo apt-get build-dep linux-image-$(uname -r)  
sudo apt-get install kernel-package  

sudo add-apt-repository ppa:webupd8team/java  
sudo apt-get update  
sudo apt-get install oracle-java8-installer  
	(Testing: javac -version)  
sudo apt-get install oracle-java8-set-default  
sudo update-alternatives --config java  
(Note: Select Auto mode)  
sudo update-alternatives --config java  
(Note: Copy the path something like /usr/lib/jvm/java-8-oracle)  
sudo nano /etc/environment  
	(Add the below line at end of the file)  
	JAVA_HOME="/usr/lib/jvm/java-8-oracle"  
source /etc/environment  
	(Testing: echo $JAVA_HOME)  

Note: Download "SMPPSim.tar.gz" from "http://www.seleniumsoftware.com/downloads.html" web site and save it in to your home folder.  
cd /home/ubuntu  
tar -zxvf SMPPSim.tar.gz  

### Deploy the REST API code (Non-Developer approach) in Kannel server:
***********************************************************************
cd /home/ubuntu  
git clone https://github.com/ReachCMP/plivosms.git  

### Deploy the REST API code (Developer approach) in Kannel server:
*******************************************************************
Login to the Kannel server  
sudo apt-get install git-all  
cd /usr/local/  
sudo wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz  
sudo tar zxvf go1.8.linux-amd64.tar.gz  
export PATH=$PATH:/usr/local/go/bin  
cd  
sudo nano .profile  
	export PATH=$PATH:/usr/local/go/bin  
	export GOROOT=/usr/local/go  
	export GOPATH=/home/ubuntu/go  
source ~/.profile  
cd /home/ubuntu/  
mkdir go  
cd go  
mkdir src  
cd src  
git clone https://github.com/ReachCMP/plivosms.git  


### Run the services:
*********************
Login to the SMSC server and execute the below commands to start SMSC simulator:  
- cd /home/ubuntu/SMPPSim  
- sudo sh startsmppsim.sh  

Open another terminal and login to the Kannel server and execute the below commands to start BearerBox:  
- cd /usr/local/sbin  
- sudo bearerbox -v 0 smskannel.conf  

Open another terminal and login to the Kannel server and execute the below commands to start SMSBox:  
- cd /usr/local/sbin  
- sudo smsbox -v 0 smskannel.conf  

Open another terminal and login to the Kannel server and execute the below commands to start REST API service (Non-Developer approach):  
- cd /home/ubuntu/plivosms  
- ./testproject  

Open another terminal and login to the Kannel server and execute the below commands to start REST API service (Developer approach):  
- cd $GOPATH/src/plivosms  
- go get github.com/LK4D4/vndr  
- go get github.com/goji/httpauth  
- go get github.com/yazver/gsmmodem/pdu  
- go build main.go  
- make build  
- ./testproject  

Alternatively, you can run the unit tests using the below command:  
- make cover  

And run the project using the below command:  
- make run  

### Testing with out using API:
*******************************
Below are the some of the URLs to do some operations from the web browser:  
- Check the status and access the SMPPSim Control Panel from a web browser (non-editable mode only): http://35.154.65.222:88/
- Check the status of the Kannel service from the Kannel server CLI: lynx -dump "http://localhost:13000/status?password=Test12"
- Check the status of the Kannel service from a web browser: http://35.154.95.122:13000/status?password=Test12
- Shutdown the Kannel service from a web browser: http://35.154.95.122:13000/shutdown?password=Test12
- Send SMS from a web browser: 
http://35.154.95.122:13013/cgi-bin/sendsms?username=chandramouli&password=Test12&smsc=SMPPSim&to=1234567891&text=Test SMS from Chandramouli  

### Testing with REST API:
**************************
- You can do different tests as you asked
- You can use any REST client to test it (For example, "Advanced REST client" extension in Google Chrome web browser)
- Below are the values that you have to give in the REST client:  
URL: http://35.154.95.122:3000/outbound/sms  
Content-Type: application/json  
Raw payload:  
{  
  "from": "1234567891",  
  "to": "3336661999",  
  "text": "Sending SMS using REST API"  
}  
Basic Auth:  
- username: chandramouli  
- password: <Confidential>  






