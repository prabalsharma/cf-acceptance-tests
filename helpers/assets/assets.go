package assets

type Assets struct {
	AspClassic                 string
	BatchScript                string
	Catnip                     string
	CredHubEnabledApp          string
	CredHubServiceBroker       string
	Dora                       string
	DoraDroplet                string
	DoraZip                    string
	DotnetCore                 string
	Fuse                       string
	GoCallsRubyZip             string
	Golang                     string
	HelloWorld                 string
	HelloRouting               string
	Java                       string
	JavaSpringZip              string
	JavaUnwriteableZip         string
	LoggregatorLoadGenerator   string
	LoggregatorLoadGeneratorGo string
	Python                     string
	Node                       string
	NodeWithProcfile           string
	Nora                       string
	Php                        string
	Proxy                      string
	RubySimple                 string
	SecurityGroupBuildpack     string
	ServiceBroker              string
	Staticfile                 string
	SyslogDrainListener        string
	Binary                     string
	LoggingRouteService        string
	TCPListener                string
	Wcf                        string
	WindowsWebapp              string
	WindowsWorker              string
	WorkerApp                  string
	MultiPortApp               string
	SpringSleuthZip            string
}

func NewAssets() Assets {
	return Assets{
		AspClassic:                 "assets/asp-classic",
		BatchScript:                "assets/batch-script",
		Catnip:                     "assets/catnip/bin",
		CredHubEnabledApp:          "assets/credhub-enabled-app/credhub-enabled-app.jar",
		CredHubServiceBroker:       "assets/credhub-service-broker",
		Dora:                       "assets/dora",
		DoraDroplet:                "assets/dora-droplet.tar.gz",
		DoraZip:                    "assets/dora.zip",
		DotnetCore:                 "assets/dotnet-core",
		Fuse:                       "assets/fuse-mount",
		GoCallsRubyZip:             "assets/go_calls_ruby.zip",
		Golang:                     "assets/golang",
		HelloRouting:               "assets/hello-routing",
		HelloWorld:                 "assets/hello-world",
		Java:                       "assets/java",
		JavaSpringZip:              "assets/java-spring/java-spring.jar",
		JavaUnwriteableZip:         "assets/java-unwriteable-dir/java-unwriteable-dir.jar",
		LoggregatorLoadGenerator:   "assets/loggregator-load-generator",
		LoggregatorLoadGeneratorGo: "assets/loggregator-load-generator-go",
		Node:                       "assets/node",
		NodeWithProcfile:           "assets/node-with-procfile",
		Nora:                       "assets/nora/NoraPublished",
		Php:                        "assets/php",
		Proxy:                      "vendor/code.cloudfoundry.org/cf-networking-release/src/example-apps/proxy",
		Python:                     "assets/python",
		RubySimple:                 "assets/ruby_simple",
		SecurityGroupBuildpack:     "assets/security_group_buildpack.zip",
		ServiceBroker:              "assets/service_broker",
		Staticfile:                 "assets/staticfile",
		SyslogDrainListener:        "assets/syslog-drain-listener",
		Binary:                     "assets/binary",
		LoggingRouteService:        "assets/logging-route-service",
		TCPListener:                "assets/tcp-listener",
		Wcf:                        "assets/wcf/Hello.Service.IIS",
		WindowsWebapp:              "assets/webapp",
		WorkerApp:                  "assets/worker-app",
		WindowsWorker:              "assets/worker",
		MultiPortApp:               "assets/multi-port-app",
		SpringSleuthZip:            "assets/spring-sleuth/spring-sleuth.jar",
	}
}
