package atc 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Atc struct {

	/*DevelopmentMode - Descr: Loosen up security for development purposes. This allows the ATC to be
configured with no authentication methods.
 Default: false
*/
	DevelopmentMode interface{} `yaml:"development_mode,omitempty"`

	/*OldResourceGracePeriod - Descr: How long to cache the result of a get step after a newer version of the
resource is found. Use Go duration format (1m = 1 minute).
 Default: 5m
*/
	OldResourceGracePeriod interface{} `yaml:"old_resource_grace_period,omitempty"`

	/*Postgresql - Descr: Name of the database to use.
 Default: atc
*/
	Postgresql Postgresql `yaml:"postgresql,omitempty"`

	/*Riemann - Descr: Port of the Riemann server to emit events to.
 Default: 5555
*/
	Riemann Riemann `yaml:"riemann,omitempty"`

	/*GithubAuth - Descr: An array of different criteria to check for when authorizing a GitHub
user. If empty, GitHub authorization is effectively disabled.
 Default: []
*/
	GithubAuth GithubAuth `yaml:"github_auth,omitempty"`

	/*BasicAuthUsername - Descr: Username for HTTP basic auth.
 Default: 
*/
	BasicAuthUsername interface{} `yaml:"basic_auth_username,omitempty"`

	/*BindIp - Descr: IP address on which the ATC should listen for HTTP traffic.
 Default: 0.0.0.0
*/
	BindIp interface{} `yaml:"bind_ip,omitempty"`

	/*DefaultCheckInterval - Descr: The interval, in Go duration format (1m = 1 minute), on which to check
for new versions of resources.

This can also be specified on a per-resource basis by specifying
`check_every` on the resource config.
 Default: 1m
*/
	DefaultCheckInterval interface{} `yaml:"default_check_interval,omitempty"`

	/*Retention - Descr: The duration to keep a failed step's containers before expiring them.
 Default: 1h
*/
	Retention Retention `yaml:"retention,omitempty"`

	/*BindPort - Descr: Port on which the ATC should listen for HTTP traffic.
 Default: 8080
*/
	BindPort interface{} `yaml:"bind_port,omitempty"`

	/*BasicAuthPassword - Descr: Password for HTTP basic auth, in plaintext.
 Default: 
*/
	BasicAuthPassword interface{} `yaml:"basic_auth_password,omitempty"`

	/*PostgresqlDatabase - Descr: Name of the database to use from the `postgresql` link.
 Default: <nil>
*/
	PostgresqlDatabase interface{} `yaml:"postgresql_database,omitempty"`

	/*Yeller - Descr: Environment name to specify for errors emitted to Yeller.
 Default: 
*/
	Yeller Yeller `yaml:"yeller,omitempty"`

	/*ResourceCacheCleanupInterval - Descr: The interval, in Go duration format (1m = 1 minute), on which to check
for and release old caches of resource versions.
 Default: 30s
*/
	ResourceCacheCleanupInterval interface{} `yaml:"resource_cache_cleanup_interval,omitempty"`

	/*PubliclyViewable - Descr: Allow viewing of pipelines as an anonymous user. Destructive operations
still require auth, and the output of builds will only be visible if
their job is configured with `public: true`.

This is useful for open-source projects, or as a convenience to make
monitoring your pipeline status easier.
 Default: false
*/
	PubliclyViewable interface{} `yaml:"publicly_viewable,omitempty"`

	/*ExternalUrl - Descr: Externally reachable URL of the ATCs. Required for OAuth.

Typically this is the URL that you as a user would use to reach your CI.
For multiple ATCs it would go to some sort of load balancer.
 Default: <nil>
*/
	ExternalUrl interface{} `yaml:"external_url,omitempty"`

	/*PeerUrl - Descr: Address used internally to reach the ATC. This will be auto-generated
using the IP of each ATC VM if not specified.

Note that this refers to an *individual ATC*, not the whole cluster. This
property is only useful if you're deploying in a way that cannot
autodetect its own IP, e.g. a `bosh-init` deployment.

You should otherwise leave this value blank.
 Default: <nil>
*/
	PeerUrl interface{} `yaml:"peer_url,omitempty"`

}