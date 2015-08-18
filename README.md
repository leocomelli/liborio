# Libório

This is a very simple example of file storage in the file system using Go. This example in particular stores files and uses simple HTTP basic authentication.

Because this is an example, the authentication credentials aren't really a secret and can be setup in the source code.

By default, files are stored in ```/usr/share/liborio```, but can be changed directly in the source code.

## Storage Structure:

The structure is composed of an application and some files, something like that...
```
* /myapp
   - mynewfile.zip (604)
   - myfile.zip (5)
```

Where ```myapp``` is my application and ```mynewfile.zip``` and ```myfile.zip``` are my files.

## Usage:

### server
[Download](https://github.com/leocomelli/liborio/releases) the latest version and run...

```
./main
```

### curl

List files
```
curl -u admin:liborio http://localhost:8080/
```

List files by application
```
curl -u admin:liborio http://localhost:8080/myapp
```

Download file
```
curl -u admin:liborio http://localhost:8080/myapp/myfile.zip
```

Upload file
```
curl -XPOST -u admin:liborio http://localhost:8080/myapp/mynewfile.zip -T /dev/mynewfile.zip
```

## An easy way to contrib

1. Clone the Libório repository
```git clone https://github.com/leocomelli/liborio.git```
2. Use docker
```docker run -ti -p 8080:8080 -v /dev/liborio:/go/src/liborio golang```
3. Go to the liborio source directory
```cd /go/src/liborio```
4. Get the go-retful lib
```go get github.com/emicklei/go-restful```
5. Running
```go run main.go```
6. Happy hacking! \o/

## License

(The MIT License)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.