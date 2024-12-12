setup-mac:
	#brew install operator-sdk -y
	operator-sdk init --domain comcast.github.io --repo github.com/kuberhealthy/crds
	operator-sdk create api --group comcast.github.io --version v1 --kind KuberhealthyCheck --resource --controller
	#operator-sdk create api --group comcast.github.io --version v1 --kind KuberhealthyJob --resource --controller
	#operator-sdk create api --group comcast.github.io --version v1 --kind KuberhealthyState --resource --controller
