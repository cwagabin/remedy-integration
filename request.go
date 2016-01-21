package main

import "errors"
import "fmt"
import "io"
import "io/ioutil"
import "net/http"

type HttpRequest struct {
  Method  string
  Url     *string
  Data    io.Reader
}

func httpRequest(c *Config, r *HttpRequest) (response_body []byte, err error) {
  req, err := http.NewRequest(r.Method, *r.Url, r.Data)
  if err != nil {
    return
  }

  // Add authentication headers
  token := fmt.Sprintf("Token token=%s", c.ApiKey)
  req.Header.Set("Authorization", token)
  req.Header.Set("Content-type", "application/json")

  // Fire request
  client := &http.Client{}
  response, err := client.Do(req)
  if err != nil {
    return
  }
  defer response.Body.Close()

  // Only 200's allowed
  if response.Status != "200 OK" {
    message := fmt.Sprintf("Got HTTP status code %s", response.Status)
    err = errors.New(message)
    return
  }


  response_body, err = ioutil.ReadAll(response.Body)
  return
}
