.PHONY:	opc

all: opc

opc:
	go build -o opc main.go
clean:
	rm opc
