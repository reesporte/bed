default: build

build: editor.go go.mod
	go build -ldflags "-w -s" .  

run: build 
	./bed

clean:
	rm bed

install: build
	cp bed /usr/local/bin

uninstall:
	rm /usr/local/bin/bed
