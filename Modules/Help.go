package CastHunter

import "fmt"

/*
he string \x1b[H\x1b[2J\x1b[3J is a series of escape codes that can be used to control the terminal or console.

The \x1b is an escape character (ASCII code 27). It is used to introduce a control sequence, which is a series of characters that control the formatting or behavior of the terminal or console.

The H control code moves the cursor to the specified position (coordinates). For example, \x1b[10;10H would move the cursor to the 10th row and 10th column.

The 2J control code clears the screen and moves the cursor to the top left corner of the screen.

The 3J control code clears the screen and moves the cursor to the top left corner of the screen, but it also removes all lines from the scrollback buffer. This means that you will not be able to use the scrollback function to view previous lines that were cleared by this control code.

So, the full sequence \x1b[H\x1b[2J\x1b[3J would clear the screen, move the cursor to the top left corner of the screen, and remove all lines from the scrollback buffer.
*/

func HelpMenu() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	/*
		|BAR0|  <\x1b> (C27*ASCII*) 				|- introduce control sequence
		|BAR1|  <H>    (CC *Control Code*)  		|- Move cursor to the specified POS For example, \x1b[10;10H would move the cursor to the 10th row and 10th column.
		|BAR2|  <2J>   (CCC *Control Code Clear*)   |- Clears the screen, places cursor to the top left corner
		|BAR3|  <3J>   (CRAL *Control Code Remove)  |- Like CCC but clears the screen and removes all lines from the scrollback buffer, meaning no scrollback or extra text

		--> C27+CC+CCC+CRAL 0-> \x1b[H\x1b[2J\x1b[3J
	*/
	Banner()
	fmt.Print("\x1b[0m") // Reset color
	fmt.Println(
		`
					Caster V1.0 STABLE
	------------------------------------------------------------
	Welcome to caster, the ghost that hunts down IoT devices for 
	Enumeration. This is the help menu for casters start up, please
	make sure all flags you create are specified with arguments.

	FORMAT OF PROGRAM
	-----------------
	Command: sudo ./buildfile [FLAGS]
	example: sudo ./main --arp=true --ssdp=true --server=true
	
	flag syntax is as follows (FLAG, DATA TYPE, Description, Options)
	-----------------------------------------------------------------

	FLAGS:
		--arp     | bool | Scan for devices using ARP      | true, false
		--ssdp    | bool | SSDP capture mode               | true, false
		--single  | bool | use a single interface for ARP  | true, false
		--help    | bool | this menu                       | true, false
		--server  | bool | start a web server for caster   | true, false 
		--errors  | bool | output errors, if false skip    | true, false
		--trace   | bool | traceback from errors or fatals | true, false
		--helpe   | bool | Loads a more descriptive page   | true, false
		--send    | int  | Delimeter in seconds for ARP    | (1,2,3 etc..) 
	
	for more information onto how and why certain flags are needed goto http://localhost:5429
	`)
}

func Help2() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	Banner()
	fmt.Print("\x1b[0m")
	fmt.Println(
		`
					Caster V1.0 STABLE
	------------------------------------------------------------
	this is a more descriptive help menu for the Caster Framework
	if you used option --helpe then you are in the right place.
	this next brick will show commands and why they are needed.
	think of it like a MAN page except shorter.

	Program syntax | (sudo ./main [FLAGS])
	Example        | sudo ./main --arp=true

	FLAGS 

	the caster framework uses and depends on small amount of flags
	most of them in fact are just here to give you a decent experience.
	Something to note is when using the caster framework you can not do 
	as much as the framework is capable with without the use of flags. 
	Below you will find a list of flags and their use cases, descriptions, 
	data types and more.

	----------------------------------------------------------------------

	--ARP: the --arp(-a) flag tells caster to start its arp module, this is 
	set to standard by default. This module allows for command options such as 
	enumerate* or view which allow you to enumerate all hosts caster has found 
	and seperated into unique lists, or view which allows you to view hosts on the
	network that were found through listeners or discovery. This flag is highly suggested 
	as it makes the target option and setting as well as enumeration process easier.

	this flag is a BOOL value         |(true/false)
	this flag has options like so     |true, false
	this is how you can use it        | 
									  |__ --arp=true  (enable ARP) 
									  	  --arp=false (disable ARP)
										  -a 		  (enable ARP)


	--ssdp: the --ssdp(-d) flag tells caster to start listening for incoming 
	SSDP (Simple Service Discovery Protocol) packets with SSDP-NOTIFY to parse 
	special UUID's that may belong to devices such as amazonfire TV's. When using 
	this program the devs assume that you do not know the UUID of the device ( for obvious reasons )
	so setting a custom UUID does not exist. However when you go to enumerate a fireTV or 
	device that may require a UUID to enumerate or send data to via services like UPnP 
	the host or target you set will be the verification. If a user set a target say 

	10.0.0.9 and it was an amazon device, if casters SSDP module has not found any UUID's 
	for that host, it will say that the module returned empty and it can not enumerate. This is 
	because some device's like fireTV devices use UUID's to grab system information or to grab 
	and make requests to the server, these are sent in the form of SSDP. In order for caster 
	to find these UUID's then it will have to listen for these packets. This is also why --ssdp 
	is a standard true flag but you can disable it ( not suggested )

	this flag is a BOOL value         |(true/false)
	this flag has options like so     |true, false
	this is how you can use it        | 
									  |__ --ssdp=true  (enable SSDP) 
									  	  --ssdp=false (disable SSDP)
										  -d           (enable SSDP)
	
	
	--send: the --send(-e) flag tells caster how often to send out ARP packets to discover hosts.
	This flag accepts any data in the form of integer, however just know that if you put 20 it will 
	send packets out EVERY 20 SECONDS! It may also be important to note that this also dictates how 
	often the data within the local server is updated. This flag is really only used for and along side 
	of the SERVER AND ARP flag.


	this flag is a INT  value         |(1,2,3,4,5,6,7 etc...)
	this is how you can use it        | 
									  |__ --send=20    (Set timer for 20 seconds) 
										  -e=20        (Set timer for 20 seconds)

	
	--server: the --server(-l) flag tells caster to load a local HTTP server, this server by defualt 
	loads and runs in the background of caster as its running its main functions. This server will contain 
	documentation, hostnames found that are seperated and also give you links to go to as far as author creds

	this flag is a BOOL value         |(true/false)
	this flag has options like so     |true, false
	this is how you can use it        | 
									  |__ --server=true  (enable server) 
									  	  --server=false (disable server)
										  -l           (enable server)
										  
										 

	--single: the --single(-s) flag is a VERY IMPORTANT FLAG FOR CERTAIN CASES. Caster ARP module and SSDP module 
	work with a networking library to find all current interfaces on the host or machine, within this list is all
	the interfaces which will be used (LO [localhost], ETH [Ethernet]) will ALL BE SKIPPED AND NOT USED. However,
	despite caster being pretty stable and nearly finished to the fullest as it will continue to expand, there have 
	been known issues that may arrise with setting --single to false. In a basic stance this flag when set to true 
	will use only one single interface that is not LO or ETH. If this flag is turned off Caster will use every single
	interface that it finds on the machine even if it is in use. This means that when caster finds say the following list 

		[lo, eth, wlan1, wlan2, wlan3, wlan4, wlan5, wlan6, wlan7]

	it will use every card within that list OTHER THAN LO AND ETH for ARP scanning and SSDP listening. This option 
	can cause arrays to glitch out, this is an issue because in the future say two cards listen and find the same host 
	which had the same MAC of another host, the arrays will collide in a sense and make output a bit wackier. This is not 
	super dangerous or anything but due to certain network structures it may prevent it from formatting fully. Out of all 
	100+ tests during development on the framework this only happened twice. It is suggested that you use --single and set it 
	to true so you can prevent such a major waste of networking systems. Caster is thread safe and trys its best to make every 
	single listener, sender, reciever and much more on seperate threads and channels making caster much more safer to use and 
	faster to run background processes such as ARP modules and servers. HOWEVER IN THE CASE THAT YOU USE EVERY INTERFACE DESPITE 
	RESULTS FOR HOSTS AND ENUMERATION BEING BETTER THE FRAMEWORK IS NOT GUARENTEED TO SAVE RESOURCES.


										  `)
}
