

install:
	sudo apt-get install libpcap-dev
	go build -o caster main.go


caster:
	sudo apt-get install libpcap-dev
	go build -o caster main.go
	rm main.go 
	rm -rf Modules
	rm go.mod go.sum Makefile README.md img DemoCaster.img
	
