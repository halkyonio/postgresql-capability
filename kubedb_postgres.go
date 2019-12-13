package main

import (
	"fmt"
	"github.com/appscode/go/encoding/json/types"
	capability "halkyon.io/api/capability/v1beta1"
	framework "halkyon.io/operator-framework"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ofst "kmodules.xyz/offshoot-api/api/v1"
	kubedbv1 "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
)

type postgres struct {
	*framework.DependentResourceHelper
}

func (res postgres) Update(toUpdate runtime.Object) (bool, error) {
	return false, nil
}

func (res *postgres) SetOwner(owner framework.Resource) {
	resource := framework.NewDependentResource(&kubedbv1.Postgres{}, owner)
	res.DependentResourceHelper = resource
	resource.SetDelegate(res)
}

func (res postgres) GetGroupVersionKind() schema.GroupVersionKind {
	return kubedbv1.SchemeGroupVersion.WithKind(kubedbv1.ResourceKindPostgres)
}

func newPostgres(owner framework.Resource) *postgres {
	p := &postgres{}
	p.SetOwner(owner)
	return p
}

func (res postgres) ownerAsCapability() *capability.Capability {
	return ownerAsCapability(res)
}

func (res postgres) Name() string {
	return PostgresName(res.Owner())
}

//buildSecret returns the postgres resource
func (res postgres) Build() (runtime.Object, error) {
	c := res.ownerAsCapability()
	ls := getAppLabels(c.Name)
	paramsMap := parametersAsMap(c.Spec.Parameters)

	postgres := &kubedbv1.Postgres{
		ObjectMeta: metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.Namespace,
			Labels:    ls,
		},
		Spec: kubedbv1.PostgresSpec{
			Version:  SetDefaultDatabaseVersionIfEmpty(c.Spec.Version),
			Replicas: replicaNumber(1),
			UpdateStrategy: apps.StatefulSetUpdateStrategy{
				Type: apps.RollingUpdateStatefulSetStrategyType,
			},
			DatabaseSecret: &core.SecretVolumeSource{
				SecretName: SetDefaultSecretNameIfEmpty(c.Name, paramsMap[DbConfigName]),
			},
			StorageType:       kubedbv1.StorageTypeEphemeral,
			TerminationPolicy: kubedbv1.TerminationPolicyDelete,
			PodTemplate: ofst.PodTemplateSpec{
				Spec: ofst.PodSpec{
					Env: []core.EnvVar{
						{Name: KubedbPgDatabaseName, Value: SetDefaultDatabaseName(paramsMap[DbName])},
					},
				},
			},
		},
	}
	return postgres, nil
}

func (postgres) ShouldBeCheckedForReadiness() bool {
	return true
}

func (res postgres) OwnerStatusField() string {
	//return res.ownerAsCapability().DependentStatusFieldName()
	return "PodName" // todo: find a way to compute this as above instead of hardcoding it
}

func (res postgres) IsReady(underlying runtime.Object) (ready bool, message string) {
	psql := underlying.(*kubedbv1.Postgres)
	ready = psql.Status.Phase == kubedbv1.DatabasePhaseRunning
	if !ready {
		msg := ""
		reason := psql.Status.Reason
		if len(reason) > 0 {
			msg = ": " + reason
		}
		message = fmt.Sprintf("%s is not ready%s", psql.Name, msg)
	}
	return
}

func (res postgres) NameFrom(underlying runtime.Object) string {
	return underlying.(*kubedbv1.Postgres).Name
}

func replicaNumber(num int) *int32 {
	q := int32(num)
	return &q
}

func SetDefaultDatabaseVersionIfEmpty(version string) types.StrYo {
	if version == "10.6-v2" {
		return types.StrYo("10.6")
	} else {
		// Map DB Version with the KubeDB Version
		switch version {
		case "9":
			return types.StrYo("9.6-v4")
		case "10":
			return types.StrYo("10.6-v2")
		case "11":
			return types.StrYo("11.2")
		default:
			return types.StrYo("10.6-v2")
		}
	}
}

/*
	// https://github.com/kubernetes/client-go/tree/master/examples/dynamic-create-update-delete-deployment
	// Approach to create dynamically tyhe object without type imported
	postgresRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	postgres := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
		},
	}
	// Create Postgres DB
		fmt.Println("Creating Postgres DB ...")
		result, err := client.Resource(postgresRes).Namespace(namespace).Create(postgres, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
*/