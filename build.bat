windres -o main.syso main.rc
go build -trimpath -ldflags "-w -s -H windowsgui"