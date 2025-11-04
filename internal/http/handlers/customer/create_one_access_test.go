package customer

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/internal/http/middleware"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestCustomerHandler_CreateOneAccess_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	eventID := "d4d7402f-dcab-443d-a829-f1817085f8da"
	timestamp := time.Now().Unix()

	// ---- Create payload ----
	event := customer.OneAccessCreate{
		Process:   "one access creation",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: customer.OneAccountCreateData{
			OneAccount: customer.OneAccountCreateRequest{
				OneAccountID:        "GMO4GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				DealerCustomerID:    "ASTVAJMF00552",
				FirstName:           "Nkoc",
				LastName:            "Maf",
				PhoneNumber:         "081234567890",
				Email:               "nkoc.maf@example.com",
				BirthDate:           "1995-08-12",
				VerificationChannel: "SMS",
				KtpNumber:           "PRRJKAESWC086H",
				Occupation:          "Engineer",
				Gender:              "FEMALE",
				RegistrationChannel: "MTOYOTA",
				RegistrationDate:    timestamp,
				ConsentGiven:        true,
				ConsentGivenAt:      timestamp,
				ConsentGivenDuring:  "REGISTRATION",
				AddressLabel:        "Rumah",
				ResidenceAddress:    "52EXP0ADG Q9G3AAD T, WI YktLIL ONb2 gSLTo a6pP YXZJ, NIl1n",
				Province:            "DKI Jakarta",
				City:                "Jakarta Barat",
				District:            "Pulo Gadung",
				Subdistrict:         "Jati",
				PostalCode:          "13220",
				DetailAddress:       "gts89detiyt",
			},
		},
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(event)
	require.NoError(t, err)

	// ---- Compute HMAC signature ----
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/one-access-creation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	rr := httptest.NewRecorder()

	// ---- Mock expectations ----
	m.mockIdempotencySvc.EXPECT().Store(eventID).Return(nil)
	m.mockSvc.EXPECT().CreateOneAccount(gomock.Any(), event).Return(nil)

	// ---- Simulate middleware context ----
	webhookHeaders := webhookDto.Headers{
		ContentType: "application/json",
		APIKey:      m.Config.Webhook.APIKey,
		Signature:   signature,
		EventID:     eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
	}
	ctx := context.WithValue(req.Context(), middleware.WebhookHeadersKey{}, webhookHeaders)
	req = req.WithContext(ctx)

	// ---- Execute handler ----
	m.handler.CreateOneAccess(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusAccepted, rr.Code)

	var resp webhookDto.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "accepted", resp.Message)
	assert.Equal(t, eventID, resp.Data.EventID)
	assert.Equal(t, "RECEIVED", resp.Data.Status)
}
