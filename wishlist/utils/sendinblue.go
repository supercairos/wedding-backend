package utils

import (
	"os"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"go.uber.org/zap"
)

func NewSibClient(logger *zap.Logger) (*sendinblue.APIClient, error) {
	cfg := sendinblue.NewConfiguration()

	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", os.Getenv("SENDINBLUE_API_KEY"))
	sib := sendinblue.NewAPIClient(cfg)

	return sib, nil
}
