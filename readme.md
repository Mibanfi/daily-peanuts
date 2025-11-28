This is a little tool that hosts the daily Peanuts strip from *gogocomics* at `localhost:8080`.

## Building
**Prerequisites:** node, npm, go

```
git clone https://github.com/Mibanfi/daily-peanuts.git
cd daily-peanuts
npm -i
go build main.go
```

### Installing as a service (on Linux)
Edit `peanuts.service` with your username and the path to this repository.

Then, run:

```
sudo cp peanuts.service /etc/systemd/system/peanuts.service
sudo systemctl enable peanuts.service
sudo systemctl start peanuts.service
```

Whereas, to uninstall:

```
sudo rm /etc/systemd/system/peanuts.service
sudo systemctl disable peanuts.service
sudo systemctl stop peanuts.service
```


## Running
```
./main.go
```