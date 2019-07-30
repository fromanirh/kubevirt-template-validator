module github.com/fromanirh/kubevirt-template-validator

go 1.12

require (
	github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/davecgh/go-spew v1.1.1
	github.com/emicklei/go-restful v2.6.0+incompatible
	github.com/fromanirh/okdutil v0.0.1
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-openapi/analysis v0.17.2
	github.com/go-openapi/errors v0.17.2
	github.com/go-openapi/jsonpointer v0.19.2
	github.com/go-openapi/jsonreference v0.19.2
	github.com/go-openapi/loads v0.17.2
	github.com/go-openapi/runtime v0.17.2
	github.com/go-openapi/spec v0.17.2
	github.com/go-openapi/strfmt v0.17.2
	github.com/go-openapi/swag v0.19.2
	github.com/go-openapi/validate v0.18.0
	github.com/go-stack/stack v1.8.0
	github.com/gogo/protobuf v1.1.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v0.0.0-20190713102442-dd8d2a22370e
	github.com/golang/protobuf v1.3.1
	github.com/google/btree v0.0.0-20160524151835-7d79101e329e
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/google/uuid v1.0.0
	github.com/googleapis/gnostic v0.2.0
	github.com/gorilla/websocket v0.0.0-20180228210902-0647012449a1
	github.com/gregjones/httpcache v0.0.0-20170728041850-787624de3eb7
	github.com/hashicorp/golang-lru v0.5.1
	github.com/imdario/mergo v0.3.5
	github.com/json-iterator/go v1.1.6
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v0.0.0-20181121151021-386d141f4c94
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
	github.com/mailru/easyjson v0.0.0-20190614124828-94de47d64c63
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.1-0.20190515112211-6a48b4839f85
	github.com/openshift/api v3.9.1-0.20190401220125-3a6077f1f910+incompatible
	github.com/openshift/client-go v0.0.0-20190401163519-84c2b942258a
	github.com/pborman/uuid v1.2.0
	github.com/peterbourgon/diskv v2.0.1+incompatible
	github.com/spf13/pflag v1.0.1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	golang.org/x/oauth2 v0.0.0-20170412232759-a6bd8cefa181
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20161028155119-f51c12702a4d
	google.golang.org/appengine v1.5.0
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apiextensions-apiserver v0.0.0-20190228180357-d002e88f6236
	k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go v0.0.0-20190228174230-b40b2a5939e4
	k8s.io/klog v0.3.0
	k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	kubevirt.io/client-go v0.19.0
	kubevirt.io/containerized-data-importer v1.9.5
	sigs.k8s.io/yaml v1.1.0
)

replace github.com/go-kit/kit => github.com/go-kit/kit v0.3.0
