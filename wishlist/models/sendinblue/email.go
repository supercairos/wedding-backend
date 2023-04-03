package sendinblue

import (
	"context"
	"encoding/json"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"github.com/supercairos/wedding-backend/wishlist/utils"
	"go.uber.org/zap"
)

type EmailService struct {
	transac *sendinblue.TransactionalEmailsApiService
	log     *zap.Logger
}

type envRecipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Check it implements the interface
var _ models.EmailService = &EmailService{}

// NewPersonService creates the person service using the given
// connection pool to a postgres DB.
func NewEmailService(logger *zap.Logger, sib *sendinblue.APIClient) (*EmailService, error) {
	return &EmailService{
		transac: sib.TransactionalEmailsApi,
		log:     logger,
	}, nil
}

func (s *EmailService) Send(email *models.Email) error {
	s.log.Info("Sending email", zap.Any("email", email))

	recipients := []envRecipient{}
	recipientsJson := utils.GetEnv("EMAIL_RECIPIENTS", "[]")
	if err := json.Unmarshal([]byte(recipientsJson), &recipients); err != nil {
		return err
	}

	to := []sendinblue.SendSmtpEmailTo{}
	for _, recipient := range recipients {
		to = append(to, sendinblue.SendSmtpEmailTo{
			Email: recipient.Email,
			Name:  recipient.Name,
		})
	}

	ce, http, err := s.transac.SendTransacEmail(
		context.Background(),
		sendinblue.SendSmtpEmail{
			To: to,
			Sender: &sendinblue.SendSmtpEmailSender{
				Email: email.From.Email,
				Name:  email.From.Name,
			},
			ReplyTo: &sendinblue.SendSmtpEmailReplyTo{
				Email: email.From.Email,
				Name:  email.From.Name,
			},
			Subject:     email.Subject,
			TextContent: email.Body,
		},
	)
	if err != nil {
		return err
	}

	if http.StatusCode != 201 {
		return models.ErrEmailNotSent
	}

	s.log.Info("Email sent", zap.Any("email", ce))
	return nil
}
