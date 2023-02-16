package main

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	chartPath := "/Users/nianyu/Dropbox/Mackup/kubeconfigs/longhorn-1.4.0/chart"
	chart, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}

	kubeconfigPath := "/Users/nianyu/Dropbox/Mackup/kubeconfigs/test-helm.kubeconfig"
	releaseName := "longhorn"
	releaseNamespace := "longhorn-system"
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfigPath, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}

	client := action.NewInstall(actionConfig)
	client.CreateNamespace = true
	client.Namespace = releaseNamespace
	client.ReleaseName = releaseName
	rel, err := client.Run(chart, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully installed release: ", rel.Name)
}
