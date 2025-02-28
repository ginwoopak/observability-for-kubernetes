/*
Copyright 2022.

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

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"k8s.io/client-go/discovery"

	"go.uber.org/zap/zapcore"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	wavefrontcomv1alpha1 "github.com/wavefronthq/observability-for-kubernetes/operator/api/v1alpha1"
	"github.com/wavefronthq/observability-for-kubernetes/operator/controllers"
	//+kubebuilder:scaffold:imports
)

var (
	version  string // populated via ldflags at build time
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(wavefrontcomv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func namespace() string {
	return getEnvOrDie("NAMESPACE")
}

func getComponentVersion(component string) string {
	return getEnvOrDie(component)
}

func getEnvOrDie(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		panic(fmt.Sprintf("%s must be set in environment", key))
	}

	return val
}

func main() {
	var probeAddr string
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	opts := zap.Options{
		Development: true, // Developer true, defaults to console writer instead of JSON
		Level:       zapcore.InfoLevel,
		TimeEncoder: zapcore.RFC3339TimeEncoder,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	config := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(config, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     "0",
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		Namespace:              namespace(),
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	objClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		setupLog.Error(err, "error creating reconciler client")
		os.Exit(1)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		setupLog.Error(err, "error creating kubernetes discovery client")
		os.Exit(1)
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		setupLog.Error(err, "error creating reconciler client")
		os.Exit(1)
	}

	defaultNS, err := cs.CoreV1().Namespaces().Get(context.Background(), "default", v12.GetOptions{})
	clusterUUID := string(defaultNS.UID)
	if err != nil {
		log.Log.Error(err, "error reading default namespace: %s")
	} else {
		log.Log.Info(fmt.Sprintf("*********************** Setting cluster uud: %s", clusterUUID))
	}

	controller, err := controllers.NewWavefrontReconciler(
		controllers.Versions{
			OperatorVersion:  version,
			CollectorVersion: getComponentVersion("COLLECTOR_VERSION"),
			ProxyVersion:     getComponentVersion("PROXY_VERSION"),
			LoggingVersion:   getComponentVersion("LOGGING_VERSION"),
		},
		objClient,
		discoveryClient,
		clusterUUID,
	)

	setupLog.Info(fmt.Sprintf("Versions %+v", controller.Versions))

	if err != nil {
		setupLog.Error(err, "error creating wavefront operator reconciler")
		os.Exit(1)
	}

	if err = controller.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to setup manager", "controller", namespace())
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
