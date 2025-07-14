package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/as-ifn-at/targeting_engine/internal/config"
	"github.com/as-ifn-at/targeting_engine/models"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

var _ = Describe("campaign test cases", func() {
	var (
		logger zerolog.Logger
		r      *gin.Engine
	)
	BeforeEach(func() {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		r = gin.Default()
		testCampaignHandler := NewCampaignHandler(config.Config{}, logger)
		routerG := r.Group("/v1/campaign")
		routerG.POST("/create", testCampaignHandler.Save)
		routerG.GET("/:id", testCampaignHandler.Get)
	})

	When("save is called", func() {
		It("should give success response with created status", func() {
			payload := []byte(`{
   			"cid" : "duolingo",
    		"name" : "Duolingo: Best way to learn ", 
    		"img" : "https://somelink2", 
    		"cta" : "Install",
    		"status" : "ACTIVE"
			}`)
			bodyReader := bytes.NewReader(payload)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/campaign/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusCreated))
		})
	})
	When("payload is invalid", func() {
		It("should give error response", func() {
			invalid := []byte(`{
   			"cid" : 1,
    		"name" : "Duolingo: Best way to learn ", 
    		"img" : "https://somelink2", 
    		"cta" : "Install",
    		"status" : "ACTIVE"
			}`)
			bodyReader := bytes.NewReader(invalid)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/v1/campaign/create", bodyReader)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
	})

	When("Get is called", func() {
		It("should return the campaign", func() {
			Campaigns["duolingo"] = models.Campaign{
				CampaignId: "duolingo",
				Name:       "Duolingo: Best way to learn ",
				Image:      "",
				CTA:        "",
				Status:     "ACTIVE",
			}

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/campaign/duolingo", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			data, _ := io.ReadAll(w.Body)

			fmt.Printf("data: %s\n", (data))

			// Expect(data).To(BeEquivalentTo(expectedData))
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("should return error if campaign is absent", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/v1/campaign/invalid", nil)
			Expect(err).To(BeNil())

			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
})
