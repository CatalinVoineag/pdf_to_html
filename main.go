package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/pdfcrowd/pdfcrowd-go"
)

func main() {
  // https://pdfcrowd.com/api/pdf-to-html-go/
  filename := flag.String("file", "", "PDF file to convert to HTML")
  //skipPreview := flag.Bool("s", false, "Skip auto-preview")
  flag.Parse()

  if *filename == "" {
    flag.Usage()
    os.Exit(1)
  }

  if err := run(*filename); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func run(filename string) error {
  client := pdfcrowd.NewPdfToHtmlClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

  tmpFile := "/tmp/sample.html"
  err := client.ConvertFileToFile(filename, tmpFile)

  defer os.Remove(tmpFile)

  preview(tmpFile)

  return err
}

func preview(filename string) error {
  cName := ""
  cParams := []string{}

  switch runtime.GOOS {
  case "linux":
    cName = "xdg-open"
  case "windows":
    cName = "cmd.exe"
    cParams = []string{"/C", "start"}
  case "darwin":
    cName = "open"
  default:
    return fmt.Errorf("OS not supported")
  }

  // Append filename to param slice
  cParams = append(cParams, filename)
  // Locate executable in PATH
  cPath, err := exec.LookPath(cName)

  if err != nil {
    return err
  }

  // Open the file using default program
  err = exec.Command(cPath, cParams...).Run()

  time.Sleep(2 * time.Second)
  return err
}
