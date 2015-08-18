package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"path"
	"bytes"
	"strings"
	"net/http"
	"path/filepath"
	"github.com/emicklei/go-restful"
)

const STORE string = "/usr/share/liborio"

func get(req *restful.Request, resp *restful.Response){
	appName := req.PathParameter("app-name")
	filename := req.PathParameter("filename")

	pathname := path.Join(STORE, fmt.Sprint(appName, "/", filename))

	log.Println("Getting file... ", pathname)

	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		pathname)
}

func upload(req *restful.Request, resp *restful.Response){
	appName := req.PathParameter("app-name")
	filename := req.PathParameter("filename")

	pathname := fmt.Sprint(STORE, "/", appName, "/", filename)

	log.Println("Uploading file... ", pathname)

	file, _ := os.Create(pathname)
	io.Copy(file, req.Request.Body)
	defer file.Close()
}

func findAll(req *restful.Request, resp *restful.Response) {
    io.WriteString(resp, walk(STORE, STORE))
}

func find(req *restful.Request, resp *restful.Response) {
	appName := req.PathParameter("app-name")
	io.WriteString(resp, walk(STORE, path.Join(STORE, appName)))
}

func walk(parent, source string) string {
	fileList := []string{}
    err := filepath.Walk(source, func(path string, f os.FileInfo, err error) error {
        if path != parent {
          fileList = append(fileList, path)
        }
        return nil
    })

    if err != nil {
        panic(err)
    }

    var cur_module string
    var buffer bytes.Buffer   
    for _, filename := range fileList {

        file, _ := os.Open(filename)
        fi, _ := file.Stat()

        if fi.IsDir() {
            i, j := strings.LastIndex(filename, "/"), strings.LastIndex(filename, path.Ext(filename))
            module_name := filename[i:j]
            if cur_module != module_name {
                buffer.WriteString(fmt.Sprint("* ", module_name, "\n"))
            }   
        } else {
            buffer.WriteString(fmt.Sprint("   - ", filepath.Base(filename), " (", fi.Size(), ")", "\n"))     
        }
    }

    return buffer.String()
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	encoded := req.Request.Header.Get("Authorization")
	if len(encoded) == 0 || "Basic YWRtaW46bGlib3Jpbw==" != encoded {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}

func main() {
	ws := new(restful.WebService)

	ws.Route(ws.GET("/").To(findAll))
	ws.Route(ws.GET("/{app-name}").To(find))
	ws.Route(ws.GET("/{app-name}/{filename}").To(get))
	ws.Route(ws.POST("/{app-name}/{filename}").
				Filter(basicAuthenticate).
	            To(upload))
	restful.Add(ws)

	log.Println("[liborio] serving files on http://localhost:8080/{app}/{filename} from local ", STORE)
	http.ListenAndServe(":8080", nil)
}