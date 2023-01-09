
caster:
	sudo apt-get install libpcap-dev
	go build -o caster main.go
	rm main.go 
	rm -rf Modules
	