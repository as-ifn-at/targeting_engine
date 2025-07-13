package config

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("config test cases", func() {
	When("envs are not passed", func() {
		It("will take default values", func() {
			testConfig := Load()
			expectedConfig := &Config{
				Port:         DefaultPort,
				DatabaseName: DefaultDbName,
				Dbpath:       DefaultDBPath,
			}
			
			Expect(testConfig).To(BeEquivalentTo(expectedConfig))
		})
	})

	When("envs are passed", func() {
		It("will take values from envs", func() {
			os.Setenv(PORT, "8080")
			os.Setenv(DBPATH, "testpath")
			os.Setenv(DBNAME, "test-db")
			testConfig := Load()
			expectedConfig := &Config{
				Port:         "8080",
				DatabaseName: "test-db",
				Dbpath:       "testpath",
			}

			Expect(testConfig).To(BeEquivalentTo(expectedConfig))
		})
	})
})
