package main

import (
  "bufio"
  "os"
  "fmt"
  "math/rand"
  "strings"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{ Layout: "layout", }))
  m.Use(martini.Static("assets"))

  m.Get("/", func(r render.Render) {
    name, err := generateName()
    if err == nil {
      r.HTML(200, "main", name)
    }
  })

  m.Run()
}

func generateName() (string, error) {
  names, err := readLines("words.txt")
  if err != nil {
    return "", err
  }
  name := fmt.Sprintf("%s %s", names[rand.Intn(len(names))], names[rand.Intn(len(names))])
  return strings.Title(name), nil
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
