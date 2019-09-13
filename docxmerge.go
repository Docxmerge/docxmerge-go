package docxmerge_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Docxmerge struct {
	baseUrl string
	apiKey  string
	tenant  string
}

type DocxmergeOptions struct {
	BaseUrl string
	ApiKey  string
	Tenant  string
}
type Data map[string]interface{}

func newMultipartFile(reader io.Reader, body *bytes.Buffer) (*multipart.Writer, error) {
	fileContents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file.docx")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return writer, nil
}

// Creates a new file upload http request with optional extra params
func newMultipartData(reader io.Reader, data Data, body *bytes.Buffer) (*multipart.Writer, error) {
	fileContents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file.docx")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(data)
	part, err = writer.CreateFormField("data")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(jsonBytes)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return writer, nil
}

func (d *Docxmerge) TransformDocument(reader io.Reader) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/TransformFile", d.baseUrl)
	body := new(bytes.Buffer)
	w, err := newMultipartFile(reader, body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}

func (d *Docxmerge) hidrateRequest(request *http.Request) {
	request.Header.Set("api-key", d.apiKey)
	request.Header.Set("x-tenant", d.tenant)
}

func (d *Docxmerge) TransformTemplate(templateName string) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/TransformTemplate?template=%s", d.baseUrl, templateName)
	body := new(bytes.Buffer)
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}

func (d *Docxmerge) MergeDocument(reader io.Reader, data Data) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/MergeFile", d.baseUrl)
	body := new(bytes.Buffer)
	w, err := newMultipartData(reader, data, body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}

func (d *Docxmerge) MergeTemplate(templateName string, data Data) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/MergeTemplate?template=%s", d.baseUrl, templateName)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}

func (d *Docxmerge) MergeAndTransformDocument(reader io.Reader, data Data) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/MergeAndTransform", d.baseUrl)
	body := new(bytes.Buffer)
	w, err := newMultipartData(reader, data, body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}
func (d *Docxmerge) MergeAndTransformTemplate(templateName string, data Data) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/MergeAndTransformTemplatePost?template=%s", d.baseUrl, templateName)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}

func NewDocxmerge(options DocxmergeOptions) *Docxmerge {
	return &Docxmerge{
		baseUrl: options.BaseUrl,
		apiKey:  options.ApiKey,
		tenant:  options.Tenant,
	}
}
