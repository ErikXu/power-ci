package istio

import (
	"fmt"
	"os"
	"path"
	"power-ci/consts"
	"power-ci/utils"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	istioInstallCmd.Flags().StringVarP(&Domain, "domain", "d", "example.com", "Domain, eg: example.com")
}

var Domain string

var template = `#!/bin/bash
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.15.0 TARGET_ARCH=x86_64 sh -

cd istio-1.15.0

./bin/istioctl install -f {ISTIO.YAML} -y

kubectl create ns bookinfo
kubectl label namespace bookinfo istio-injection=enabled
kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml -n bookinfo

kubectl apply -f samples/addons/prometheus.yaml
kubectl apply -f samples/addons/grafana.yaml
kubectl apply -f samples/addons/kiali.yaml
kubectl apply -f samples/addons/jaeger.yaml

kubectl apply -f {ISTIO-GATEWAY.YAML}
kubectl apply -f {BOOKINFO-VS.YAML}
kubectl apply -f {PROMETHEUS-VS.YAML}
kubectl apply -f {GRAFANA-VS.YAML}
kubectl apply -f {KIALI-VS.YAML}
kubectl apply -f {JAEGER-QUERY.YAML}
kubectl apply -f {JAEGER-VS.YAML}`

var istioYaml = `---
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
spec:
  meshConfig:
    enableTracing: true
    defaultConfig:
      tracing:
        sampling: 100
        zipkin:
          address: jaeger-collector.istio-system:9411
    enablePrometheusMerge: true

  components:
    # Istio Gateway feature
    ingressGateways:
    - name: istio-ingressgateway
      enabled: true
      k8s:
        service:
          ports:
          - name: status-port
            nodePort: 30021
            port: 15021
            protocol: TCP
            targetPort: 15021
          - name: http2
            nodePort: 30080
            port: 80
            protocol: TCP
            targetPort: 8080
          - name: https
            nodePort: 30443
            port: 443
            protocol: TCP
            targetPort: 8443`

var istioGatewayTemplate = `---
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: istio-gateway
  namespace: istio-system
spec:
  selector:
    istio: ingressgateway
  servers:
  - hosts:
    - '*.{DOMAIN}'
    port:
      name: http
      number: 80
      protocol: HTTP`

var bookinfoVsTemplate = `---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: bookinfo-vs
  namespace: bookinfo
spec:
  hosts:
  - 'bookinfo.{DOMAIN}'
  gateways:
  - istio-system/istio-gateway
  http:
  - match:
    - uri:
        exact: /productpage
    - uri:
        prefix: /static
    - uri:
        exact: /login
    - uri:
        exact: /logout
    - uri:
        prefix: /api/v1/products
    route:
    - destination:
        host: productpage.bookinfo.svc.cluster.local
        port:
          number: 9080`

var prometheusVsTemplate = `---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: prometheus-vs
  namespace: istio-system
spec:
  hosts:
    - 'prometheus.{DOMAIN}'
  gateways:
    - istio-system/istio-gateway
  http:
    - match:
        - uri:
            prefix: /
      route:
        - destination:
            host: prometheus.istio-system.svc.cluster.local
            port:
              number: 9090`

var grafanaVsTemplate = `---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: grafana-vs
  namespace: istio-system
spec:
  hosts:
    - 'grafana.{DOMAIN}'
  gateways:
    - istio-system/istio-gateway
  http:
    - match:
        - uri:
            prefix: /
      route:
        - destination:
            host: grafana.istio-system.svc.cluster.local
            port:
              number: 3000`

var kialiVsTemplate = `---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: kiali-vs
  namespace: istio-system
spec:
  hosts:
    - 'kiali.{DOMAIN}'
  gateways:
    - istio-system/istio-gateway
  http:
    - match:
        - uri:
            prefix: /
      route:
        - destination:
            host: kiali.istio-system.svc.cluster.local
            port:
              number: 20001`

var jaegeQuerySvcYaml = `---
apiVersion: v1
kind: Service
metadata:
  name: jaeger-query
  namespace: istio-system
  labels:
    app: jaeger
spec:
  type: ClusterIP
  ports:
  - name: http
    port: 16686
    targetPort: 16686
    protocol: TCP
  selector:
    app: jaeger`

var jaegerVsTemplate = `---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: jaeger-vs
  namespace: istio-system
spec:
  hosts:
    - 'jaeger.{DOMAIN}'
  gateways:
    - istio-system/istio-gateway
  http:
    - match:
        - uri:
            prefix: /
      route:
        - destination:
            host: jaeger-query.istio-system.svc.cluster.local
            port:
              number: 16686`

func writeYaml(istioWorkspace string, template string, domain string, filename string) string {
	yaml := strings.Replace(template, "{DOMAIN}", domain, -1)
	yamlPath := path.Join(istioWorkspace, filename)
	f, _ := os.Create(yamlPath)
	f.WriteString(yaml)
	return yamlPath
}

var istioInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install istio",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		istioWorkspace := path.Join(homeDir, consts.Workspace, "istio")
		os.MkdirAll(istioWorkspace, os.ModePerm)

		istioYamlPath := writeYaml(istioWorkspace, istioYaml, Domain, "istio.yaml")
		istioGatewayYamlPath := writeYaml(istioWorkspace, istioGatewayTemplate, Domain, "istio-gateway.yaml")
		bookinfoVsYamlPath := writeYaml(istioWorkspace, bookinfoVsTemplate, Domain, "bookinfo-vs.yaml")
		prometheusVsYamlPath := writeYaml(istioWorkspace, prometheusVsTemplate, Domain, "prometheus-vs.yaml")
		grafanaVsYamlPath := writeYaml(istioWorkspace, grafanaVsTemplate, Domain, "grafana-vs.yaml")
		kialiVsYamlPath := writeYaml(istioWorkspace, kialiVsTemplate, Domain, "kiali-vs.yaml")
		jaegerQuerySvcYamlPath := writeYaml(istioWorkspace, jaegeQuerySvcYaml, Domain, "jaeger-query.yaml")
		jaegerVsYamlPath := writeYaml(istioWorkspace, jaegerVsTemplate, Domain, "jaeger-vs.yaml")

		script := strings.Replace(template, "{ISTIO.YAML}", istioYamlPath, -1)
		script = strings.Replace(script, "{ISTIO-GATEWAY.YAML}", istioGatewayYamlPath, -1)
		script = strings.Replace(script, "{BOOKINFO-VS.YAML}", bookinfoVsYamlPath, -1)
		script = strings.Replace(script, "{PROMETHEUS-VS.YAML}", prometheusVsYamlPath, -1)
		script = strings.Replace(script, "{GRAFANA-VS.YAML}", grafanaVsYamlPath, -1)
		script = strings.Replace(script, "{KIALI-VS.YAML}", kialiVsYamlPath, -1)
		script = strings.Replace(script, "{JAEGER-QUERY.YAML}", jaegerQuerySvcYamlPath, -1)
		script = strings.Replace(script, "{JAEGER-VS.YAML}", jaegerVsYamlPath, -1)

		filepath := utils.WriteScript("install-istio.sh", script)
		utils.ExecuteScript(filepath)

		fmt.Println("Install success. More info please refer https://istio.io/latest/docs/setup/getting-started/")
	},
}
