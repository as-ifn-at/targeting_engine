package handlers

import (
	"bytes"
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
		testCampaignHandler := NewRuleHandler(config.Config{}, logger)
		routerG := r.Group("/v1/rule")
		routerG.POST("/create", testCampaignHandler.Save)
		routerG.GET("/:id", testCampaignHandler.Get)
	})

	When("save is called", func() {
		It("should give success response with created status", func() {
			Campaigns["duolingo"] = models.Campaign{
				CampaignId: "duolingo",
				Name:       "Duolingo: Best way to learn ",
				Image:      "",
				CTA:        "",
				Status:     "ACTIVE",
			}
			payload := []byte(`{
  						"cid": "duolingo",
  						"rules": {
    					"exclude_country": ["US"],
    					"include_os": ["Android"]
  						}
					}`)
			bodyReader := bytes.NewReader(payload)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/rule/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusCreated))
		})
		It("should give error response if there is no campaign for the rule", func() {
			payload := []byte(`{
  						"cid": "invalid",
  						"rules": {
    					"exclude_country": ["US"],
    					"include_os": ["Android"]
  						}
					}`)
			bodyReader := bytes.NewReader(payload)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/rule/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})

		It("should give error response if the rule violates the exclusion/inclusion", func() {
			payload := []byte(`{
  						"cid": "duolingo",
  						"rules": {
    					"exclude_country": ["US"],
						"include_country": ["IND"],
    					"include_os": ["Android"]
  						}
					}`)
			bodyReader := bytes.NewReader(payload)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/rule/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should give error response if the rule violates the exclusion/inclusion", func() {
			payload := []byte(`{
  						"cid": "duolingo",
  						"rules": {
    					"exclude_country": ["US"],
    					"include_os": ["Android"],
    					"exclude_os": ["Android"]
  						}
					}`)
			bodyReader := bytes.NewReader(payload)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/rule/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should give error response if the rule violates the exclusion/inclusion", func() {
			payload := []byte(`{
  						"cid": "duolingo",
  						"rules": {
    					"exclude_country": ["US"],
    					"include_app": ["app.com"],
    					"exclude_app": ["app.c"]
  						}
					}`)
			bodyReader := bytes.NewReader(payload)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/rule/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
	})

	When("payload is invalid", func() {
		It("should give error response", func() {
			invalid := []byte(`{
  						"cid": 1,
  						"rules": {
    					"exclude_country": ["US"],
    					"include_os": ["Android"]
  						}
					}`)
			bodyReader := bytes.NewReader(invalid)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/rule/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
	})

	When("Get is called", func() {
		It("should return the campaign", func() {
			Rules[""] = models.TargetRules{
				CampaignId: "duolingo",
				Rules: models.RuleSet{
					IncludeCountry: []string{"US"},
				},
			}

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/rule/duolingo", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			data, _ := io.ReadAll(w.Body)

			fmt.Printf("data: %s\n", (data))

			// Expect(data).To(BeEquivalentTo(expectedData))
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("should return error if campaign is absent", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/rule/invalid", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
})
