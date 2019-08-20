package plugnotas

import (
	"encoding/json"
	"fmt"
)

// NfseResponse - Struct para definir o objeto NfseResponse
type NfseResponse struct {
	Documents []*Nfse `json:"documents"`
	Protocol  string  `json:"protocol"`
}

// Servico - Struct para definir o objeto serviço
type Servico struct {
	Codigo                    string `json:"codigo"`
	IDIntegracao              string `json:"idIntegracao"`
	Discriminacao             string `json:"discriminacao"`
	CodigoTributacao          string `json:"codigoTributacao"`
	Cnae                      string `json:"cnae"`
	CodigoCidadeIncidencia    string `json:"codigoCidadeIncidencia"`
	DescricaoCidadeIncidencia string `json:"descricaoCidadeIncidencia"`
	Iss                       *Iss   `json:"iss"`
	Valor                     *Valor `json:"valor"`
}

// Nfse - Struct para definir o objeto Nfse
type Nfse struct {
	IDIntegracao string     `json:"IdIntegracao"`
	EnviarEmail  bool       `json:"enviarEmail"`
	Prestador    *Prestador `json:"prestador"`
	Tomador      *Tomador   `json:"tomador"`
	Servico      *Servico   `json:"servico"`
	ID           string     `json:"id"`
}
type ResumoNfse struct {
	ID                string  `json:"id"`
	IDIntegracao      string  `json:"idIntegracao"`
	Emissao           string  `json:"emissao"`
	TipoAutorizacao   string  `json:"tipoAutorizacao"`
	Situacao          string  `json:"situacao"`
	Prestador         string  `json:"prestador"`
	Tomador           string  `json:"tomador"`
	ValorServico      float32 `json:"valorServico"`
	NumeroNfse        string  `json:"numeroNfse"`
	Serie             string  `json:"serie"`
	Lote              string  `json:"lote"`
	CodigoVerificacao string  `json:"codigoVerificacao"`
	Autorizacao       string  `json:"autorizacao"`
	Mensagem          string  `json:"mensagem"`
	Pdf               string  `json:"pdf"`
	XML               string  `json:"xml"`
	Cancelamento      string  `json:"cancelamento"`
}
type ResumoNfseList []*ResumoNfse

// CreateNfse - enviar uma lista de notas
func (plugnotas *Client) CreateNfse(req []*Nfse) (*NfseResponse, *ErrorResponse) {

	data, err := json.Marshal(req)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Message{
				Message: err.Error(),
			},
		}
	}
	var result = &NfseResponse{}
	err, errAPI := plugnotas.Request("POST", "/nfse", data, result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Message{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return result, nil

}

// GetNfseByID buscar nota por id
func (plugnotas *Client) GetNfseByID(id string) (*Nfse, *ErrorResponse) {
	var result = &Nfse{}
	err, errAPI := plugnotas.Request("GET", fmt.Sprintf("/nfse/%s", id), nil, result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Message{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return result, nil

}

// GetNfseByID buscar nota por id
func (plugnotas *Client) ConsultarNfse(id string) (ResumoNfseList, *ErrorResponse) {
	var result ResumoNfseList
	err, errAPI := plugnotas.Request("GET", fmt.Sprintf("/nfse/consultar/%s", id), nil, &result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Message{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return result, nil

}

// GetNfseByID buscar nota por id
func (plugnotas *Client) CancelarNfse(id string) (*Message, *ErrorResponse) {
	var result Message
	err, errAPI := plugnotas.Request("POST", fmt.Sprintf("/nfse/cancelar/%s", id), []byte("{}"), &result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Message{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return &result, nil

}
