package handlers

import (
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
