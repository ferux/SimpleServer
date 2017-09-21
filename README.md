# Simple Server
An application written on Go for hosting static files.

## Install:

```bash
git clone github.com/ferux/SimpleServer
```

## Usage

Runs from command line and by default uses the following parameters:
Listen address: **0.0.0.0:8080**
Static Files loads from root directory, e.g. **"./"**
Debug mode is **off**.

If you would like to specify parameters use the following flags in command line:
**-d** to enable debug mode;
**-l** "address:port" to specify the address and port;
**-l** ":port" to specify port only;
**-a** "/path/to/assets/" to specify the directory of files;
**-h** to show help info.

## Example

##For WIndows##
simpleServer.exe -l "192.168.10.11:80" -a "C:\www\static\" -d

##For Linux##
simpleServer.exe -l "192.168.10.11:80" -a "/home/usr/dev/www/static/" -d

## License

MIT