package docxmerge_go

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func newDocxmerge() *Docxmerge {
	return NewDocxmerge(DocxmergeOptions{
		BaseUrl: "http://localhost:5101",
		Tenant:  "default",
		ApiKey:  "vdnpUV4ZTLeYYrcyvF3XcKe4ZuToY5",
	})
}
func TestTransformDocument(t *testing.T) {
	docxmerge := newDocxmerge()
	reader, err := os.Open("./fixtures/helloworld.docx")
	if err != nil {
		t.Fatalf("Fallo al abrir el documento %v", err)
	}
	pdf, err := docxmerge.TransformDocument(reader)
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

func TestTransformTemplate(t *testing.T) {
	docxmerge := newDocxmerge()

	pdf, err := docxmerge.TransformTemplate("example-invoice")
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
func TestMergeDocument(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	reader, err := os.Open("./fixtures/helloworld.docx")
	if err != nil {
		t.Fatalf("Fallo al abrir el documento %v", err)
	}
	pdf, err := docxmerge.MergeDocument(reader, data)
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

func TestMergeTemplate(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	pdf, err := docxmerge.MergeTemplate("helloworld", data)
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

func TestMergeAndTransformDocument(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	reader, err := os.Open("./fixtures/helloworld.docx")
	if err != nil {
		t.Fatalf("Fallo al abrir el documento %v", err)
	}
	pdf, err := docxmerge.MergeAndTransformDocument(reader, data)
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

func TestMergeAndTransformTemplate(t *testing.T) {
	data := Data{
		"hello_world": "Hola mundo",
	}
	docxmerge := newDocxmerge()
	pdf, err := docxmerge.MergeAndTransformTemplate("helloworld", data)
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
