/*
Copyright 2020 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v1 "github.com/rancher/rancher-operator/pkg/apis/rancher.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type RoleTemplateBindingHandler func(string, *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error)

type RoleTemplateBindingController interface {
	generic.ControllerMeta
	RoleTemplateBindingClient

	OnChange(ctx context.Context, name string, sync RoleTemplateBindingHandler)
	OnRemove(ctx context.Context, name string, sync RoleTemplateBindingHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() RoleTemplateBindingCache
}

type RoleTemplateBindingClient interface {
	Create(*v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error)
	Update(*v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error)
	UpdateStatus(*v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.RoleTemplateBinding, error)
	List(namespace string, opts metav1.ListOptions) (*v1.RoleTemplateBindingList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RoleTemplateBinding, err error)
}

type RoleTemplateBindingCache interface {
	Get(namespace, name string) (*v1.RoleTemplateBinding, error)
	List(namespace string, selector labels.Selector) ([]*v1.RoleTemplateBinding, error)

	AddIndexer(indexName string, indexer RoleTemplateBindingIndexer)
	GetByIndex(indexName, key string) ([]*v1.RoleTemplateBinding, error)
}

type RoleTemplateBindingIndexer func(obj *v1.RoleTemplateBinding) ([]string, error)

type roleTemplateBindingController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewRoleTemplateBindingController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) RoleTemplateBindingController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &roleTemplateBindingController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromRoleTemplateBindingHandlerToHandler(sync RoleTemplateBindingHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.RoleTemplateBinding
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.RoleTemplateBinding))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *roleTemplateBindingController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.RoleTemplateBinding))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateRoleTemplateBindingDeepCopyOnChange(client RoleTemplateBindingClient, obj *v1.RoleTemplateBinding, handler func(obj *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error)) (*v1.RoleTemplateBinding, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *roleTemplateBindingController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *roleTemplateBindingController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *roleTemplateBindingController) OnChange(ctx context.Context, name string, sync RoleTemplateBindingHandler) {
	c.AddGenericHandler(ctx, name, FromRoleTemplateBindingHandlerToHandler(sync))
}

func (c *roleTemplateBindingController) OnRemove(ctx context.Context, name string, sync RoleTemplateBindingHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromRoleTemplateBindingHandlerToHandler(sync)))
}

func (c *roleTemplateBindingController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *roleTemplateBindingController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *roleTemplateBindingController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *roleTemplateBindingController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *roleTemplateBindingController) Cache() RoleTemplateBindingCache {
	return &roleTemplateBindingCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *roleTemplateBindingController) Create(obj *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error) {
	result := &v1.RoleTemplateBinding{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *roleTemplateBindingController) Update(obj *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error) {
	result := &v1.RoleTemplateBinding{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *roleTemplateBindingController) UpdateStatus(obj *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error) {
	result := &v1.RoleTemplateBinding{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *roleTemplateBindingController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *roleTemplateBindingController) Get(namespace, name string, options metav1.GetOptions) (*v1.RoleTemplateBinding, error) {
	result := &v1.RoleTemplateBinding{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *roleTemplateBindingController) List(namespace string, opts metav1.ListOptions) (*v1.RoleTemplateBindingList, error) {
	result := &v1.RoleTemplateBindingList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *roleTemplateBindingController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *roleTemplateBindingController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.RoleTemplateBinding, error) {
	result := &v1.RoleTemplateBinding{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type roleTemplateBindingCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *roleTemplateBindingCache) Get(namespace, name string) (*v1.RoleTemplateBinding, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.RoleTemplateBinding), nil
}

func (c *roleTemplateBindingCache) List(namespace string, selector labels.Selector) (ret []*v1.RoleTemplateBinding, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RoleTemplateBinding))
	})

	return ret, err
}

func (c *roleTemplateBindingCache) AddIndexer(indexName string, indexer RoleTemplateBindingIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.RoleTemplateBinding))
		},
	}))
}

func (c *roleTemplateBindingCache) GetByIndex(indexName, key string) (result []*v1.RoleTemplateBinding, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.RoleTemplateBinding, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.RoleTemplateBinding))
	}
	return result, nil
}

type RoleTemplateBindingStatusHandler func(obj *v1.RoleTemplateBinding, status v1.RoleTemplateBindingStatus) (v1.RoleTemplateBindingStatus, error)

type RoleTemplateBindingGeneratingHandler func(obj *v1.RoleTemplateBinding, status v1.RoleTemplateBindingStatus) ([]runtime.Object, v1.RoleTemplateBindingStatus, error)

func RegisterRoleTemplateBindingStatusHandler(ctx context.Context, controller RoleTemplateBindingController, condition condition.Cond, name string, handler RoleTemplateBindingStatusHandler) {
	statusHandler := &roleTemplateBindingStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromRoleTemplateBindingHandlerToHandler(statusHandler.sync))
}

func RegisterRoleTemplateBindingGeneratingHandler(ctx context.Context, controller RoleTemplateBindingController, apply apply.Apply,
	condition condition.Cond, name string, handler RoleTemplateBindingGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &roleTemplateBindingGeneratingHandler{
		RoleTemplateBindingGeneratingHandler: handler,
		apply:                                apply,
		name:                                 name,
		gvk:                                  controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterRoleTemplateBindingStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type roleTemplateBindingStatusHandler struct {
	client    RoleTemplateBindingClient
	condition condition.Cond
	handler   RoleTemplateBindingStatusHandler
}

func (a *roleTemplateBindingStatusHandler) sync(key string, obj *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type roleTemplateBindingGeneratingHandler struct {
	RoleTemplateBindingGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *roleTemplateBindingGeneratingHandler) Remove(key string, obj *v1.RoleTemplateBinding) (*v1.RoleTemplateBinding, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.RoleTemplateBinding{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *roleTemplateBindingGeneratingHandler) Handle(obj *v1.RoleTemplateBinding, status v1.RoleTemplateBindingStatus) (v1.RoleTemplateBindingStatus, error) {
	objs, newStatus, err := a.RoleTemplateBindingGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
