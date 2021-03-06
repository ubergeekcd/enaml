package groundcrew 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Tsa struct {

	/*Host - Descr: IP or DNS address of the TSA server to register with.

If not specified, the `tsa` link is used.
 Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

	/*PrivateKey - Descr: Private key to use when authenticating with the TSA.
If not specified, a deployment-scoped default is used.
 Default: 
*/
	PrivateKey interface{} `yaml:"private_key,omitempty"`

	/*Port - Descr: Port of the TSA server to register with.

Only used when `tsa.host` is also specified. Otherwise the port is
autodiscovered via the `tsa` link.
 Default: 2222
*/
	Port interface{} `yaml:"port,omitempty"`

	/*HostPublicKey - Descr: Host key to verify for the TSA server.
If not specified, a deployment-scoped default is used.
 Default: 
*/
	HostPublicKey interface{} `yaml:"host_public_key,omitempty"`

}