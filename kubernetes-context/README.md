# Kubernetes context

Display the current kubernetes context and the namespace set in that context. The shell functions require [yq](https://github.com/mikefarah/yq), compiled from source for best performance. All versions assume that `KUBECONFIG` is not set, and print an error otherwise.
