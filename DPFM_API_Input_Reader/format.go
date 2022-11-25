package dpfm_api_input_reader

import (
	"data-platform-api-language-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPartnerFunction() *requests.Language {
	data := sdc.Language
	return &requests.Language{
		Language: data.Language,
	}
}
