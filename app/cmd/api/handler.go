package api

import (
	"io"
	"net/http"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(r.Method, "http://localhost:2024/register", r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for name, values := range r.Header {
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }


    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    
    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body) 
}
func HandleBookmarketGetBooks(w http.ResponseWriter, r *http.Request) {

    req, err := http.NewRequest(r.Method, "http://localhost:2022/getbooks", r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    for name, values := range r.Header {
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body) 
}
func HandleAuth(w http.ResponseWriter, r *http.Request) {

    req, err := http.NewRequest(r.Method, "http://localhost:2022/auth", r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    for name, values := range r.Header {
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body) 
}
func HandleGetPurshaes(w http.ResponseWriter, r *http.Request) {

    req, err := http.NewRequest(r.Method, "http://localhost:2023/getpursh", r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    for name, values := range r.Header {
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body) 
}
func HandleForFrontend(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./src/index.html")
}