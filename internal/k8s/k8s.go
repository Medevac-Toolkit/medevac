package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var knownFaultyStatuses = []string{
	"CrashLoopBackOff",
	"ImagePullBackOff",
	"ErrImagePull",
	"CreateContainerConfigError",
	"InvalidImageName",
	"Evicted",
	"OOMKilled",
	"ContainerCannotRun",
	"DeadlineExceeded",
}

func isFaultyStatus(status string) bool {
	for _, faultyStatus := range knownFaultyStatuses {
		if status == faultyStatus {
			return true
		}
	}
	return false
}

func GetKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
		config, err = kubeconfig.ClientConfig()
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func ListPods(clientset *kubernetes.Clientset) error {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %s, Name: %s\n", pod.Namespace, pod.Name)
	}

	return nil
}

func GetProblematicPods(clientset *kubernetes.Clientset) error {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, pod := range pods.Items {
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if isFaultyStatus(containerStatus.State.Waiting.Reason) {
				fmt.Printf("Namespace: %s, Pod: %s, Container: %s, Reason: %s\n",
					pod.Namespace, pod.Name, containerStatus.Name, containerStatus.State.Waiting.Reason)
			}
		}
	}
	return nil
}
