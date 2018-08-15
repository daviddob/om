package commands_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
)

var _ = Describe("DeleteSSLCertificate", func() {
	var (
		fakeService *fakes.DeleteSSLCertificateService
		fakeLogger  *fakes.Logger
		command     commands.DeleteSSLCertificate
	)

	BeforeEach(func() {
		fakeService = &fakes.DeleteSSLCertificateService{}
		fakeLogger = &fakes.Logger{}
		command = commands.NewDeleteSSLCertificate(fakeService, fakeLogger)
	})

	Describe("Execute", func() {
		It("deletes the custom ssl certificate", func() {
			err := command.Execute([]string{})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeService.DeleteSSLCertificateCallCount()).To(Equal(1))

			Expect(fakeLogger.PrintfCallCount()).To(Equal(2))
			format, content := fakeLogger.PrintfArgsForCall(0)
			Expect(fmt.Sprintf(format, content...)).To(Equal("Successfully deleted custom SSL Certificate and reverted to the provided self-signed SSL certificate.\n"))
			format, content = fakeLogger.PrintfArgsForCall(1)
			Expect(fmt.Sprintf(format, content...)).To(Equal("Please allow about 1 min for the new certificate to take effect.\n"))
		})

		Context("failure cases", func() {
			Context("when the service fails to delete the custom certificate", func() {
				It("returns an error", func() {
					fakeService.DeleteSSLCertificateReturns(errors.New("failed to delete certificate"))

					err := command.Execute([]string{})
					Expect(err).To(MatchError("failed to delete certificate"))
				})
			})
			Context("when an unknown flag is provided", func() {
				It("returns an error", func() {
					err := command.Execute([]string{"--badflag"})
					Expect(err).To(MatchError("could not parse delete-ssl-certificate flags: flag provided but not defined: -badflag"))
				})
			})
		})
	})

	Describe("Usage", func() {
		It("returns usage info", func() {
			usage := command.Usage()
			Expect(usage).To(Equal(jhanda.Usage{
				Description:      "This authenticated command deletes a custom certificate applied to Ops Manager and reverts to the auto-generated cert",
				ShortDescription: "deletes certificate applied to Ops Manager",
				Flags:            command.Options,
			}))
		})
	})
})
