// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewUpstream(namespace, name string) *Upstream {
	return &Upstream{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Upstream) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Upstream) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type UpstreamList []*Upstream
type UpstreamsByNamespace map[string]UpstreamList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list UpstreamList) Find(namespace, name string) (*Upstream, error) {
	for _, upstream := range list {
		if upstream.Metadata.Name == name {
			if namespace == "" || upstream.Metadata.Namespace == namespace {
				return upstream, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find upstream %v.%v", namespace, name)
}

func (list UpstreamList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) Names() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.Metadata.Name)
	}
	return names
}

func (list UpstreamList) NamespacesDotNames() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.Metadata.Namespace+"."+upstream.Metadata.Name)
	}
	return names
}

func (list UpstreamList) Sort() UpstreamList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list UpstreamList) Clone() UpstreamList {
	var upstreamList UpstreamList
	for _, upstream := range list {
		upstreamList = append(upstreamList, proto.Clone(upstream).(*Upstream))
	}
	return upstreamList
}

func (list UpstreamList) ByNamespace() UpstreamsByNamespace {
	byNamespace := make(UpstreamsByNamespace)
	for _, upstream := range list {
		byNamespace.Add(upstream)
	}
	return byNamespace
}

func (byNamespace UpstreamsByNamespace) Add(upstream ...*Upstream) {
	for _, item := range upstream {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace UpstreamsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace UpstreamsByNamespace) List() UpstreamList {
	var list UpstreamList
	for _, upstreamList := range byNamespace {
		list = append(list, upstreamList...)
	}
	return list.Sort()
}

func (byNamespace UpstreamsByNamespace) Clone() UpstreamsByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Upstream{}

// Kubernetes Adapter for Upstream

func (o *Upstream) GetObjectKind() schema.ObjectKind {
	t := UpstreamCrd.TypeMeta()
	return &t
}

func (o *Upstream) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Upstream)
}

var UpstreamCrd = crd.NewCrd("gloo.solo.io",
	"upstreams",
	"gloo.solo.io",
	"v1",
	"Upstream",
	"us",
	&Upstream{})
