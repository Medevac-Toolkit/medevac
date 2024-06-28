package strategies

import (
	"fmt"
	"plugin"

	v1 "k8s.io/api/core/v1"
)

type Strategy interface {
	DetectIssues(pod v1.Pod) string
}

type RestartDetector struct {
	Threshold int
}

func (r RestartDetector) DetectIssues(pod v1.Pod) string {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		fmt.Println(containerStatus)
	}
	return ""
}

type PodStatusDetector struct{}

func (p PodStatusDetector) DetectIssues(pod v1.Pod) string {
	if pod.Status.Phase == v1.PodPending {
		return fmt.Sprintf("Pod %s is pending", pod.Name)
	}
	if pod.Status.Phase == v1.PodFailed {
		return fmt.Sprintf("Pod %s has failed", pod.Name)
	}
	for _, condition := range pod.Status.Conditions {
		if condition.Type == v1.PodReady && condition.Status != v1.ConditionTrue {
			return fmt.Sprintf("Pod %s is not ready", pod.Name)
		}
	}
	return ""
}

func LoadCustomStrategy(path string) (Strategy, error) {
	// Placeholder for future implementation of loading custom strategies
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	sym, err := p.Lookup("CustomDetectorPlugin")
	if err != nil {
		return nil, err
	}
	strategy, ok := sym.(Strategy)
	if !ok {
		return nil, fmt.Errorf("unexpected type from module symbol")
	}
	return strategy, nil
}
