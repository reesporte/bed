default: build

build: editor.go go.mod
	go build -ldflags "-w -s" .  

run: build 
	./bed

clean:
	rm bed

