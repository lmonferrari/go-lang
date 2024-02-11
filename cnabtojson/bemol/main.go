package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type CnabHeader struct {
	IdRegistro                   string `json:"id_registro"`
	IdArquivoRemessa             string `json:"id_arquivo_remessa"`
	LitRemessa                   string `json:"lit_remessa"`
	CodigoServico                string `json:"codigo_servico"`
	LitServico                   string `json:"lit_servico"`
	CodigoCedente                string `json:"codigo_cedente"`
	NomeCedente                  string `json:"nome_cedente"`
	NumeroBancoCamaraCompensacao string `json:"numero_banco_camara_compensacao"`
	NomeBanco                    string `json:"nome_banco"`
	DataGravacaoArquivo          string `json:"data_gravacao_arquivo"`
	Branco1                      string `json:"branco_1"`
	IdSistema                    string `json:"id_sistema"`
	NumSeqArquivo                string `json:"num_seq_arquivo"`
	DataCessao                   string `json:"data_cessao"`
	Coobrigacao                  string `json:"coobrigacao"`
	NaturezaOperacao             string `json:"natureza_operacao"`
	Segmento                     string `json:"segmento"`
	TipoCalculo                  string `json:"tipo_calculo"`
	Branco2                      string `json:"branco_2"`
	NumSeqRegistro               string `json:"num_seq_registro"`
}

type CnabTrailer struct {
	IdRegistro     string `json:"id_registro"`
	Branco         string `json:"branco"`
	NumSeqRegistro string `json:"num_seq_registro"`
}

type CnabStatements struct {
	IdRegistro                   string `json:"id_registro"`
	TipoInscricaoCedente         string `json:"tipo_inscricao_cedente"`
	NumInscricaoCedente          string `json:"num_inscricao_cedente"`
	NumContrato                  string `json:"num_contrato"`
	NumParcela                   string `json:"num_parcela"`
	TotalParcelas                string `json:"total_parcelas"`
	Brancos1                     string `json:"brancos_1"`
	SeuNumero                    string `json:"seu_numero"`
	Zeros1                       string `json:"zeros_1"`
	IdTituloBanco                string `json:"id_titulo_banco"`
	DigitoNossoNumero            string `json:"digito_nosso_numero"`
	ValorPago                    string `json:"valor_pago"`
	DataLiquidacao               string `json:"data_liquidacao"`
	DataObito                    string `json:"data_obito"`
	Brancos2                     string `json:"brancos_2"`
	CodigoLoja                   string `json:"codigo_loja"`
	CodigoProduto                string `json:"codigo_produto"`
	CodigoTipoFinanciamento      string `json:"codigo_tipo_financiamento"`
	CodigoSetorSacado            string `json:"codigo_setor_sacado"`
	IdOcorrencia                 string `json:"id_ocorrencia"`
	DataVencimentoTitulo         string `json:"data_vencimento_titulo"`
	ValorTitulo                  string `json:"valor_titulo"`
	BancoEncarregadoCobranca     string `json:"banco_encarregado_cobranca"`
	AgenciaDepositaria           string `json:"agencia_depositaria"`
	SeqRenegociacao              string `json:"seq_renegociacao"`
	Brancos3                     string `json:"brancos_3"`
	EspecieTitulo                string `json:"especie_titulo"`
	Identificacao                string `json:"identificacao"`
	DataEmissaoTitulo            string `json:"data_emissao_titulo"`
	PrimeiraInstrucao            string `json:"primeira_instrucao"`
	SegundaInstrucao             string `json:"segunda_instrucao"`
	ValorDiaAtraso               string `json:"valor_dia_atraso"`
	DataLimiteDesconto           string `json:"data_limite_desconto"`
	ValorDesconto                string `json:"valor_desconto"`
	ValorPresenteParcela         string `json:"valor_presente_parcela"`
	ValorAbatimento              string `json:"valor_abatimento"`
	Indice                       string `json:"indice"`
	CorrecaoIndice               string `json:"correcao_indice"`
	IdTipoInscricaoSacado        string `json:"id_tipo_inscricao_sacado"`
	NumInscricaoSacado           string `json:"num_inscricao_sacado"`
	NomeSacado                   string `json:"nome_sacado"`
	EnderecoCompleto             string `json:"endereco_completo"`
	PrimeiraMensagem             string `json:"primeira_mensagem"`
	Cep                          string `json:"cep"`
	DataNascimentoSacado         string `json:"data_nascimento_sacado"`
	IdadeSacado                  string `json:"idade_sacado"`
	TipoInscricaoSacadorAvalista string `json:"tipo_inscricao_sacador_avalista"`
	NumInscricaoSacadoAvalista   string `json:"num_inscricao_sacado_avalista"`
	NomeSacadorAvalista          string `json:"nome_sacador_avalista"`
	ValorFinanciado              string `json:"valor_financiado"`
	MatriculaSacado              string `json:"matricula_sacado"`
	CnpjEnteConsignante          string `json:"cnpj_ente_consignante"`
	TipoContrato                 string `json:"tipo_contrato"`
	TipoAtivo                    string `json:"tipo_ativo"`
	Nfe                          string `json:"nfe"`
	Brancos4                     string `json:"brancos_4"`
	NumSeqRegistro               string `json:"num_seq_registro"`
}

type Cnab struct {
	Header     CnabHeader
	Trailer    CnabTrailer
	Statements []CnabStatements
}

func cnabToJson(filename string) ([]byte, error) {
	var cnabHeader CnabHeader
	var cnabTrailer CnabTrailer
	var cnabStatements CnabStatements
	var cnabStatementsList []CnabStatements

	// Read the entire file into memory
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		// Ignore empty lines
		if len(line) == 0 {
			continue
		}

		if i == 0 {
			cnabHeader.IdRegistro = line[0:1]
			cnabHeader.IdArquivoRemessa = line[1:2]
			cnabHeader.LitRemessa = line[2:9]
			cnabHeader.CodigoServico = line[9:11]
			cnabHeader.LitServico = line[11:26]
			cnabHeader.CodigoCedente = line[26:46]
			cnabHeader.NomeCedente = line[46:76]
			cnabHeader.NumeroBancoCamaraCompensacao = line[76:79]
			cnabHeader.NomeBanco = line[79:94]
			cnabHeader.DataGravacaoArquivo = line[94:102]
			cnabHeader.Branco1 = line[102:107]
			cnabHeader.IdSistema = line[107:110]
			cnabHeader.NumSeqArquivo = line[110:117]
			cnabHeader.DataCessao = line[117:125]
			cnabHeader.Coobrigacao = line[125:126]
			cnabHeader.NaturezaOperacao = line[126:127]
			cnabHeader.Segmento = line[127:130]
			cnabHeader.TipoCalculo = line[130:131]
			cnabHeader.Branco2 = line[131:494]
			cnabHeader.NumSeqRegistro = line[494:500]
		} else if line[0:1] == "9" {
			cnabTrailer.IdRegistro = line[0:1]
			cnabTrailer.Branco = line[1:494]
			cnabTrailer.NumSeqRegistro = line[494:500]

		} else {
			if line[0:1] != "9" {
				cnabStatements.IdRegistro = line[0:1]
				cnabStatements.TipoInscricaoCedente = line[1:3]
				cnabStatements.NumInscricaoCedente = line[3:17]
				cnabStatements.NumContrato = line[17:29]
				cnabStatements.NumParcela = line[29:32]
				cnabStatements.TotalParcelas = line[32:35]
				cnabStatements.Brancos1 = line[35:47]
				cnabStatements.SeuNumero = line[47:72]
				cnabStatements.Zeros1 = line[72:83]
				cnabStatements.IdTituloBanco = line[83:94]
				cnabStatements.DigitoNossoNumero = line[94:95]
				cnabStatements.ValorPago = line[95:108]
				cnabStatements.DataLiquidacao = line[108:116]
				cnabStatements.DataObito = line[116:124]
				cnabStatements.Brancos2 = line[124:126]
				cnabStatements.CodigoLoja = line[126:129]
				cnabStatements.CodigoProduto = line[129:132]
				cnabStatements.CodigoTipoFinanciamento = line[132:135]
				cnabStatements.CodigoSetorSacado = line[135:138]
				cnabStatements.IdOcorrencia = line[138:140]
				cnabStatements.DataVencimentoTitulo = line[140:148]
				cnabStatements.ValorTitulo = line[148:161]
				cnabStatements.BancoEncarregadoCobranca = line[161:164]
				cnabStatements.AgenciaDepositaria = line[164:169]
				cnabStatements.SeqRenegociacao = line[169:177]
				cnabStatements.Brancos3 = line[177:179]
				cnabStatements.EspecieTitulo = line[179:181]
				cnabStatements.Identificacao = line[181:182] // codigo_aceite
				cnabStatements.DataEmissaoTitulo = line[182:190]
				cnabStatements.PrimeiraInstrucao = line[190:192]
				cnabStatements.SegundaInstrucao = line[192:194]
				cnabStatements.ValorDiaAtraso = line[194:207]
				cnabStatements.DataLimiteDesconto = line[208:215]
				cnabStatements.ValorDesconto = line[215:228]
				cnabStatements.ValorPresenteParcela = line[228:241]
				cnabStatements.ValorAbatimento = line[241:254]
				cnabStatements.Indice = line[254:257]
				cnabStatements.CorrecaoIndice = line[257:264]
				cnabStatements.IdTipoInscricaoSacado = line[264:266]
				cnabStatements.NumInscricaoSacado = line[266:280]
				cnabStatements.NomeSacado = line[280:320]
				cnabStatements.EnderecoCompleto = line[320:360]
				cnabStatements.PrimeiraMensagem = line[360:372]
				cnabStatements.Cep = line[372:380]
				cnabStatements.DataNascimentoSacado = line[380:388]
				cnabStatements.IdadeSacado = line[388:391]
				cnabStatements.TipoInscricaoSacadorAvalista = line[391:393]
				cnabStatements.NumInscricaoSacadoAvalista = line[393:407]
				cnabStatements.NomeSacadorAvalista = line[407:447]
				cnabStatements.ValorFinanciado = line[447:460]
				cnabStatements.MatriculaSacado = line[460:472]
				cnabStatements.CnpjEnteConsignante = line[472:486]
				cnabStatements.TipoContrato = line[486:489]
				cnabStatements.TipoAtivo = line[489:490]
				cnabStatements.Brancos4 = line[490:494]
				cnabStatements.NumSeqRegistro = line[494:500]

				cnabStatementsList = append(cnabStatementsList, cnabStatements)
			}
		}
	}

	// Parse the line into a Cnab struct
	var cnab Cnab
	cnab.Header = cnabHeader
	cnab.Trailer = cnabTrailer
	cnab.Statements = cnabStatementsList

	jsonBytes, err := json.Marshal(cnab)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func main() {
	fmt.Println("Bemol")
	jsonBytes, err := cnabToJson("CON04042023085614.rem")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("CON04042023085614.json", jsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
