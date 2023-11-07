# GARNI Weather Station Server

This server collects meteo data from Garni weather station to provide them to other services such as home automation. The data is stored only in memory as the latest snapshot.

It was tested with GARNI 1025 Arcus but should work with any weather station that supports sending data to custom server.

In order to work you need to configure the weather station to send data to this server's IP address and port:
```http://[IP_ADDRESS]:8080```


## How to run it on Raspberry Pi

I am running this meteo server on Raspberry Pi Zero 2.

You will need to:

1) Install (Go)[https://go.dev/] on your computer if you don't have it already
2) Compile the source
```
GOARCH=arm GOARM=7 GOOS=linux go build meteo-server.go
```
3) Copy the binary to the Raspberry Pi
```
scp server username@IP_ADDRESS:/home/username/meteo
```
4) Log in to Raspberry and run the server
```
./meteo-server
```
5) To run the server automatically on startup, add this line to `/etc/rc.local`:
```
sudo /home/username/meteo/meteo-server &

```
