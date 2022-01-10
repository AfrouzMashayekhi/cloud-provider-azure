/*
Copyright 2018 The Kubernetes Authors.

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

package utils

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	aznetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"

	v1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	serviceTimeout        = 5 * time.Minute
	serviceTimeoutBasicLB = 10 * time.Minute
)

// DeleteService deletes a service
func DeleteService(cs clientset.Interface, ns string, serviceName string) error {
	zero := int64(0)
	err := cs.CoreV1().Services(ns).Delete(context.TODO(), serviceName, metav1.DeleteOptions{GracePeriodSeconds: &zero})
	Logf("Deleting service %s in namespace %s", serviceName, ns)
	if err != nil {
		return err
	}
	return wait.PollImmediate(poll, deletionTimeout, func() (bool, error) {
		if _, err := cs.CoreV1().Services(ns).Get(context.TODO(), serviceName, metav1.GetOptions{}); err != nil {
			return apierrs.IsNotFound(err), nil
		}
		return false, nil
	})
}

// DeleteServiceIfExists deletes a service if it exists, return nil if not exists
func DeleteServiceIfExists(cs clientset.Interface, ns string, serviceName string) error {
	err := DeleteService(cs, ns, serviceName)
	if apierrs.IsNotFound(err) {
		Logf("Service %s does not exist, no need to delete", serviceName)
		return nil
	}
	return err
}

// GetServiceDomainName cat prefix and azure suffix
func GetServiceDomainName(prefix string) (ret string) {
	suffix := extractSuffix()
	ret = prefix + suffix
	Logf("Get domain name: %s", ret)
	return
}

// WaitServiceExposure returns ip of ingress
func WaitServiceExposure(cs clientset.Interface, namespace string, name string) (string, error) {
	var service *v1.Service
	var err error

	timeout := serviceTimeout
	if skuEnv := os.Getenv(LoadBalancerSkuEnv); skuEnv != "" {
		if strings.EqualFold(skuEnv, string(aznetwork.LoadBalancerSkuNameBasic)) {
			timeout = serviceTimeoutBasicLB
		}
	}

	if wait.PollImmediate(10*time.Second, timeout, func() (bool, error) {
		service, err = cs.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			if IsRetryableAPIError(err) {
				return false, nil
			}
			return false, err
		}

		IngressList := service.Status.LoadBalancer.Ingress
		if len(IngressList) == 0 {
			err = fmt.Errorf("Cannot find Ingress in limited time")
			Logf("Fail to find ingress, retry it in 10 seconds")
			return false, nil
		}
		return true, nil
	}) != nil {
		return "", err
	}
	ip := service.Status.LoadBalancer.Ingress[0].IP
	Logf("Exposure successfully, get external ip: %s", ip)
	return ip, nil
}

// WaitUpdateServiceExposure returns ip of ingress
func WaitUpdateServiceExposure(cs clientset.Interface, namespace string, name string, targetIP string, expectSame bool) error {
	var service *v1.Service
	var err error
	poll := 10 * time.Second
	timeout := 10 * time.Minute

	return wait.PollImmediate(poll, timeout, func() (bool, error) {
		service, err = cs.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			if IsRetryableAPIError(err) {
				return false, nil
			}
			return false, err
		}

		IngressList := service.Status.LoadBalancer.Ingress
		if len(IngressList) == 0 {
			err = fmt.Errorf("Cannot find Ingress in limited time")
			Logf("Fail to get ingress, retry it in %v seconds", poll)
			return false, nil
		}
		if targetIP != service.Status.LoadBalancer.Ingress[0].IP == expectSame {
			if expectSame {
				Logf("still unmatched external IP, retry it in %v seconds", poll)
			} else {
				Logf("External IP is still %s", targetIP)
			}
			return false, nil
		}
		Logf("Exposure successfully")
		return true, nil
	})
}

// extractSuffix obtains the server domain name suffix
func extractSuffix() string {
	c := obtainConfig()
	prefix := ExtractDNSPrefix()
	url := c.Clusters[prefix].Server
	suffix := url[strings.Index(url, "."):]
	if strings.Contains(suffix, ":") {
		suffix = suffix[:strings.Index(suffix, ":")]
	}
	return suffix
}
