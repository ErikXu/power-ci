package istio

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"power-ci/consts"
	"strings"

	"github.com/creack/pty"
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

var istioInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install istio",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		istioWorkspace := path.Join(homeDir, consts.Workspace, "istio")
		os.MkdirAll(istioWorkspace, os.ModePerm)

		istioYamlPath := path.Join(istioWorkspace, "istio.yaml")
		f, _ := os.Create(istioYamlPath)
		f.WriteString(istioYaml)

		istioGatewayYaml := strings.Replace(istioGatewayTemplate, "{DOMAIN}", Domain, -1)
		istioGatewayYamlPath := path.Join(istioWorkspace, "istio-gateway.yaml")
		f, _ = os.Create(istioGatewayYamlPath)
		f.WriteString(istioGatewayYaml)

		bookinfoVsYaml := strings.Replace(bookinfoVsTemplate, "{DOMAIN}", Domain, -1)
		bookinfoVsYamlPath := path.Join(istioWorkspace, "bookinfo-vs.yaml")
		f, _ = os.Create(bookinfoVsYamlPath)
		f.WriteString(bookinfoVsYaml)

		prometheusVsYaml := strings.Replace(prometheusVsTemplate, "{DOMAIN}", Domain, -1)
		prometheusVsYamlPath := path.Join(istioWorkspace, "prometheus-vs.yaml")
		f, _ = os.Create(prometheusVsYamlPath)
		f.WriteString(prometheusVsYaml)

		grafanaVsYaml := strings.Replace(grafanaVsTemplate, "{DOMAIN}", Domain, -1)
		grafanaVsYamlPath := path.Join(istioWorkspace, "grafana-vs.yaml")
		f, _ = os.Create(grafanaVsYamlPath)
		f.WriteString(grafanaVsYaml)

		kialiVsYaml := strings.Replace(kialiVsTemplate, "{DOMAIN}", Domain, -1)
		kialiVsYamlPath := path.Join(istioWorkspace, "kiali-vs.yaml")
		f, _ = os.Create(kialiVsYamlPath)
		f.WriteString(kialiVsYaml)

		jaegerQuerySvcYamlPath := path.Join(istioWorkspace, "jaeger-query.yaml")
		f, _ = os.Create(jaegerQuerySvcYamlPath)
		f.WriteString(jaegeQuerySvcYaml)

		jaegerVsYaml := strings.Replace(jaegerVsTemplate, "{DOMAIN}", Domain, -1)
		jaegerVsYamlPath := path.Join(istioWorkspace, "jaeger-vs.yaml")
		f, _ = os.Create(jaegerVsYamlPath)
		f.WriteString(jaegerVsYaml)

		script := strings.Replace(template, "{ISTIO.YAML}", istioYamlPath, -1)
		script = strings.Replace(script, "{ISTIO-GATEWAY.YAML}", istioGatewayYamlPath, -1)
		script = strings.Replace(script, "{BOOKINFO-VS.YAML}", bookinfoVsYamlPath, -1)
		script = strings.Replace(script, "{PROMETHEUS-VS.YAML}", prometheusVsYamlPath, -1)
		script = strings.Replace(script, "{GRAFANA-VS.YAML}", grafanaVsYamlPath, -1)
		script = strings.Replace(script, "{KIALI-VS.YAML}", kialiVsYamlPath, -1)
		script = strings.Replace(script, "{JAEGER-QUERY.YAML}", jaegerQuerySvcYamlPath, -1)
		script = strings.Replace(script, "{JAEGER-VS.YAML}", jaegerVsYamlPath, -1)

		filepath := path.Join(homeDir, consts.Workspace, "install-istio.sh")
		f, _ = os.Create(filepath)
		f.WriteString(script)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			fmt.Println("Install failed")
			return
		}
		io.Copy(os.Stdout, f)

		fmt.Println("Install success, more info please refer https://istio.io/latest/docs/setup/getting-started/")
	},
}
