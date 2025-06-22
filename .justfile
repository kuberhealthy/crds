setup-mac:
	#brew install operator-sdk -y
	operator-sdk init --domain kuberhealthy.github.io --repo github.com/kuberhealthy/crds --owner "Kuberhealthy Authors"
	operator-sdk create api --group kuberhealthy.github.io --version v4 --kind KuberhealthyCheck --resource --controller
