module halkyon.io/kubedb-capability

go 1.13

require (
	github.com/appscode/go v0.0.0-20191119085241-0887d8ec2ecc
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	halkyon.io/api v1.0.0-rc.5.0.20200312223923-071629b2929c
	halkyon.io/operator-framework v1.0.0-beta.6.0.20200311123519-8276f263c10b
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.17.0
	kmodules.xyz/offshoot-api v0.0.0-20190901210649-de049192326c
	kubedb.dev/apimachinery v0.13.0-rc.2
	sigs.k8s.io/controller-runtime v0.3.0
)

replace (
	github.com/census-instrumentation/opencensus-proto => github.com/census-instrumentation/opencensus-proto v0.2.1
	github.com/go-check/check => github.com/go-check/check v0.0.0-20180628173108-788fd7840127 // needed to avoid goreleaser failure
	gomodules.xyz/cert => gomodules.xyz/cert v1.0.1 // needed to avoid error from checksum server
	gomodules.xyz/jsonpatch/v2 => gomodules.xyz/jsonpatch/v2 v2.0.1
	k8s.io/api => k8s.io/api v0.0.0-20190805182251-6c9aa3caf3d6 // kubernetes-1.14.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190805182715-88a2adca7e76+incompatible
)
