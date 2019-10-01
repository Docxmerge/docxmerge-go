package docxmerge_go

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const version string = ""

func newDocxmerge() *Docxmerge {
	return NewDocxmerge(DocxmergeOptions{
		BaseUrl: "http://localhost:5101",
		Tenant:  "default",
		ApiKey:  "26JZ5iPpD4U3b9z7lqkXeB2OGsbdF7",
	})
}

func TestRenderFile(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	reader, err := os.Open("./fixtures/helloworld.docx")
	if err != nil {
		t.Fatalf("Fallo al abrir el documento %v", err)
	}
	pdf, err := docxmerge.RenderFile(reader, data, "PDF")
	if err != nil {
		t.Fatalf("Fallo al transformar el documento %v", err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(pdf)
	if err != nil {
		t.Fatalf("Fallo al copiar el stream %v", err)
	}
	log.Printf("Pdf %d", buf.Len())
}

func TestRenderTemplate(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	pdf, err := docxmerge.RenderTemplate("hello_world2", data, "PDF", version)
	if err != nil {
		t.Fatalf("Fallo al transformar el documento %v", err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(pdf)
	if err != nil {
		t.Fatalf("Fallo al copiar el stream %v", err)
	}
	log.Printf("Pdf %d", buf.Len())
}
func TestRenderUrl(t *testing.T) {
	data := Data{
		"name": "David",
		"logo": "https://docxmerge.com/assets/android-chrome-512x512.png",
	}
	docxmerge := newDocxmerge()
	url := "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3&download=true"
	pdf, err := docxmerge.RenderUrl(url, data, "PDF")
	if err != nil {
		t.Fatalf("Fallo al transformar el documento %v", err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(pdf)
	if err != nil {
		t.Fatalf("Fallo al copiar el stream %v", err)
	}
	log.Printf("Pdf %d", buf.Len())
	ioutil.WriteFile("./tmp/render_url.pdf", buf.Bytes(), 0640)
}

func TestRenderWithVersionTemplate(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	pdf, err := docxmerge.RenderTemplate("hello_world2", data, "PDF", "")
	if err != nil {
		t.Fatalf("Fallo al transformar el documento %v", err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(pdf)
	if err != nil {
		t.Fatalf("Fallo al copiar el stream %v", err)
	}
	log.Printf("Pdf %d", buf.Len())
}
