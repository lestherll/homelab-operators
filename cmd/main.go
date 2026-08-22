package main

import (
	"flag"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	janitorv1alpha1 "github.com/lestherll/homelab-operators/api/janitor/v1alpha1"
	janitorcontroller "github.com/lestherll/homelab-operators/internal/controller/janitor"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")

	// registry maps a controller name to the func that wires it into the manager.
	registry = map[string]func(ctrl.Manager) error{
		"janitor": func(mgr ctrl.Manager) error {
			return (&janitorcontroller.StalePodPolicyReconciler{}).SetupWithManager(mgr)
		},
	}
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = janitorv1alpha1.AddToScheme(scheme)
}

func main() {
	var controllers string
	var metricsAddr string
	flag.StringVar(&controllers, "controllers", "", "comma-separated list of controllers to run (default: all)")
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "address the metrics endpoint binds to")
	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	for _, name := range selected(controllers) {
		setup, ok := registry[name]
		if !ok {
			setupLog.Error(nil, "unknown controller", "name", name)
			os.Exit(1)
		}
		if err := setup(mgr); err != nil {
			setupLog.Error(err, "unable to set up controller", "name", name)
			os.Exit(1)
		}
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// selected returns the controller names to enable, defaulting to every
// registered controller when the flag is empty.
func selected(flagValue string) []string {
	if flagValue == "" {
		names := make([]string, 0, len(registry))
		for name := range registry {
			names = append(names, name)
		}
		return names
	}
	return strings.Split(flagValue, ",")
}
