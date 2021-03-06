package k8sutils_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kubevirt/kubevirt-template-validator/internal/pkg/k8sutils"
)

var _ = Describe("TlsInfo", func() {
	var certDir string

	BeforeEach(func() {
		var err error
		certDir, err = ioutil.TempDir("", "certs")
		Expect(err).ToNot(HaveOccurred())
	})

	It("should return nil if no certificate exists", func() {
		tlsInfo := k8sutils.TLSInfo{CertsDirectory: certDir}
		tlsInfo.Init()
		defer tlsInfo.Clean()
		tlsConfig := tlsInfo.CrateTlsConfig()

		Consistently(func() *tls.Certificate {
			cert, _ := tlsConfig.GetCertificate(nil)
			return cert
		}, time.Second).Should(BeNil())
	})

	It("should load certificate", func() {
		writeCertificate(certDir)
		tlsInfo := k8sutils.TLSInfo{CertsDirectory: certDir}
		tlsInfo.Init()
		defer tlsInfo.Clean()
		tlsConfig := tlsInfo.CrateTlsConfig()

		Eventually(func() (*tls.Certificate, error) {
			return tlsConfig.GetCertificate(nil)
		}, time.Second).ShouldNot(BeNil())
	})

	It("should reload new certificate", func() {
		tlsInfo := k8sutils.TLSInfo{CertsDirectory: certDir}
		tlsInfo.Init()
		defer tlsInfo.Clean()
		tlsConfig := tlsInfo.CrateTlsConfig()

		Consistently(func() *tls.Certificate {
			cert, _ := tlsConfig.GetCertificate(nil)
			return cert
		}, time.Second).Should(BeNil())

		writeCertificate(certDir)

		Eventually(func() (*tls.Certificate, error) {
			return tlsConfig.GetCertificate(nil)
		}, time.Second).ShouldNot(BeNil())
	})

	It("should keep old certificate on failure", func() {
		writeCertificate(certDir)
		tlsInfo := k8sutils.TLSInfo{CertsDirectory: certDir}
		tlsInfo.Init()
		defer tlsInfo.Clean()
		tlsConfig := tlsInfo.CrateTlsConfig()

		Eventually(func() (*tls.Certificate, error) {
			return tlsConfig.GetCertificate(nil)
		}, time.Second).ShouldNot(BeNil())

		Expect(ioutil.WriteFile(filepath.Join(certDir, k8sutils.CertFilename), []byte{}, 0777)).To(Succeed())

		Consistently(func() (*tls.Certificate, error) {
			return tlsConfig.GetCertificate(nil)
		}, time.Second).ShouldNot(BeNil())
	})

	AfterEach(func() {
		os.RemoveAll(certDir)
	})
})

func writeCertificate(dir string) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		Fail(err.Error())
	}

	now := time.Now()
	template := x509.Certificate{
		SerialNumber:          new(big.Int).SetInt64(0),
		Subject:               pkix.Name{CommonName: "test.kubevirt.io"},
		NotBefore:             now.UTC(),
		NotAfter:              now.Add(24 * time.Hour).UTC(),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, key.Public(), key)
	if err != nil {
		Fail(err.Error())
	}

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	Expect(ioutil.WriteFile(filepath.Join(dir, k8sutils.CertFilename), certPEM, 0777)).To(Succeed())
	Expect(ioutil.WriteFile(filepath.Join(dir, k8sutils.KeyFilename), keyPEM, 0777)).To(Succeed())
}

func TestTlsInfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Certificate Reload Suite")
}
