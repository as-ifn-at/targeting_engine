package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/models"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
)

var _ = Describe("campaign test cases", func() {
	var (
		logger zerolog.Logger
		r      *gin.Engine
	)
	BeforeEach(func() {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		r = gin.Default()
		testCampaignHandler := NewDeliverHandler(config.Config{}, logger)
		routerG := r.Group("/v1/delivery")
		routerG.GET("", testCampaignHandler.Get)
	})

	When("Get is called", func() {
		Campaigns["spotify"] = models.Campaign{
			CampaignId: "spotify",
			Name:       "Duolingo: Best way to learn ",
			Image:      "https://somelink2",
			CTA:        "Install",
			Status:     "ACTIVE",
		}

		Rules["spotify"] = models.TargetRules{
			CampaignId: "spotify",
			Rules: models.RuleSet{
				ExcludeCountry: []string{"us"},
				IncludeOS:      []string{"android"},
			},
		}
		It("should return the matching list of applicable campaigns", func() {

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/delivery?app=com.abc.xyz&country=germany&os=android", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			data, _ := io.ReadAll(w.Body)

			fmt.Printf("data: %s\n", (data))

			// Expect(data).To(BeEquivalentTo(expectedData))
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("should return error if the url is missing app", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/delivery?country=germany&os=android", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			data, _ := io.ReadAll(w.Body)

			fmt.Printf("data: %s\n", (data))

			// Expect(data).To(BeEquivalentTo(expectedData))
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should return error if the url is missing country", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/delivery?app=com.abc.xyz&os=android", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			data, _ := io.ReadAll(w.Body)

			fmt.Printf("data: %s\n", (data))

			// Expect(data).To(BeEquivalentTo(expectedData))
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should return error if the url is missing os", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/delivery?app=com.abc.xyz&country=germany", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			data, _ := io.ReadAll(w.Body)

			fmt.Printf("data: %s\n", (data))

			// Expect(data).To(BeEquivalentTo(expectedData))
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should return empty resp if there is no matching campaign", func() {
			Campaigns["subwaysurfer"] = models.Campaign{
				CampaignId: "subwaysurfer",
				Name:       "Duolingo: Best way to learn ",
				Image:      "https://somelink2",
				CTA:        "Install",
				Status:     "PENDING",
			}
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/delivery?app=com.abc.xyz&country=us&os=android", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			dataBytes, _ := io.ReadAll(w.Body)
			newData := []models.DeliverResponse{}
			_ = json.Unmarshal(dataBytes, &newData)

			expectedData := []models.DeliverResponse{}
			Expect(newData).To(BeEquivalentTo(expectedData))

			Expect(w.Code).To(Equal(http.StatusNoContent))
		})
	})
})
