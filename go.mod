module halkyon.io/postgresql-capability

go 1.13

require (
	github.com/appscode/go v0.0.0-20191025021232-311ac347b3ef
	github.com/natefinch/pie v0.0.0-20170715172608-9a0d72014007
	halkyon.io/api v1.0.0-beta.7
	halkyon.io/operator-framework v0.0.0-20191108175501-3d0a053bc383
	halkyon.io/plugins v0.0.0-20191108175809-0426164e2120
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	kmodules.xyz/offshoot-api v0.0.0-20190901210649-de049192326c
	kubedb.dev/apimachinery v0.13.0-rc.2
)

replace (
	gomodules.xyz/cert => gomodules.xyz/cert v1.0.1
	gomodules.xyz/jsonpatch/v2 => gomodules.xyz/jsonpatch/v2 v2.0.1
	halkyon.io/operator-framework => ../operator-framework
	halkyon.io/plugins => ../plugins
	k8s.io/api => k8s.io/api v0.0.0-20190805182251-6c9aa3caf3d6 // kubernetes-1.14.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190811223248-5a95b2df4348
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190805182715-88a2adca7e76+incompatible
)
