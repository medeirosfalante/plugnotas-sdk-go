package plugnotas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
	Token  string
}

type Endereco struct {
	TipoLogradouro  string `json:"tipoLogradouro"`
	Logradouro      string `json:"logradouro"`
	Numero          string `json:"numero"`
	Complemento     string `json:"complemento"`
	TipoBairro      string `json:"tipoBairro"`
	Bairro          string `json:"bairro"`
	CodigoPais      string `json:"codigoPais"`
	CodigoCidade    string `json:"codigoCidade"`
	DescricaoCidade string `json:"descricaoCidade"`
	Estado          string `json:"estado"`
	Cep             string `json:"cep"`
}

type Tomador struct {
	CpfCnpj            string `json:"cpfCnpj"`
	InscricaoMunicipal string `json:"inscricaoMunicipal"`
	InscricaoEstadual  string `json:"inscricaoEstadual"`
	RazaoSocial        string `json:"razaoSocial"`
	NomeFantasia       string `json:"nomeFantasia"`
}

type Prestador struct {
	CpfCnpj                  string    `json:"cpfCnpj"`
	InscricaoMunicipal       string    `json:"inscricaoMunicipal"`
	InscricaoEstadual        string    `json:"inscricaoEstadual"`
	RazaoSocial              string    `json:"razaoSocial"`
	NomeFantasia             string    `json:"nomeFantasia"`
	Endereco                 *Endereco `json:"endereco"`
	SimplesNacional          bool      `json:"simplesNacional"`
	IncentivadorCultural     bool      `json:"incentivadorCultural"`
	IncentivoFiscal          bool      `json:"incentivoFiscal"`
	RegimeTributario         int       `json:"regimeTributario"`
	RegimeTributarioEspecial int       `json:"regimeTributarioEspecial"`
}

type Iss struct {
	Aliquota      int     `json:"aliquota"`
	Exigibilidade int     `json:"exigibilidade"`
	Valor         float64 `json:"valor"`
	ValorRetido   float64 `json:"valorRetido"`
}

type Valor struct {
	Servico                float32 `json:"servico"`
	BaseCalculo            float32 `json:"baseCalculo"`
	Deducoes               float32 `json:"deducoes"`
	DescontoCondicionado   float32 `json:"descontoCondicionado"`
	DescontoIncondicionado float32 `json:"descontoIncondicionado"`
	Liquido                float32 `json:"liquido"`
}

type ErrorResponse struct {
	Error *Message `json:"error"`
}

type Message struct {
	Message string      `json:"message"`
	Data    interface{} `json:"Data"`
}

func NewClient(token string) *Client {
	return &Client{
		client: &http.Client{Timeout: 10 * time.Second},
		Token:  token,
	}
}

func (plugnotas *Client) devProd() string {
	if os.Getenv("ENV") == "develop" {
		return "https://api.sandbox.plugnotas.com.br"
	}
	return "https://api.plugnotas.com.br"
}

func (plugnotas *Client) Request(method, action string, body []byte, out interface{}) (error, *ErrorResponse) {
	if plugnotas.client == nil {
		plugnotas.client = &http.Client{Timeout: 10 * time.Second}
	}
	endpoint := fmt.Sprintf("%s%s", plugnotas.devProd(), action)
	fmt.Printf("body %s\n\n", string(body))
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	if plugnotas.Token != "" {
		req.Header.Add("x-api-key", plugnotas.Token)
	}
	res, err := plugnotas.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	fmt.Printf("bodyResponse %s", bodyResponse)
	if res.StatusCode > 200 {
		var errAPI ErrorResponse
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return err, nil
		}
		return nil, &errAPI
	}
	err = json.Unmarshal(bodyResponse, out)
	if err != nil {
		return err, nil
	}

	return nil, nil

}
