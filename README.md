# Kuberhealthy CRDs

This repository contains the CRD packages and the code generation for Kuberhealthy's custom resource definitions (CRDs) for Kubernetes.  This package is included in the upstream [Kuberhealthy Repository](https://github.com/kuberhealthy/kuberhealthy).

- The files that define the CRDs are located in the `api/v2/` directory of this repository.
- The output YAMLs for these CRDs are placed in `config/crd/bases/` after running `make`.
- The controller code for listening for and acting on events is generated in `internal/controller/` where it is then copied to [github.com/kuberhealthy/kuberhealthy/internal/controller](github.com/kuberhealthy/kuberhealthy/internal/controller).


## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

