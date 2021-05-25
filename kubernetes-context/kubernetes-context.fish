function kubernetes-context
  if [ -n "$KUBECONFIG" ]
    echo -n "\$KUBECONFIG is not supported"
    return
  end

  set -l context (yq e '.current-context // ""' ~/.kube/config)

  if [ -z "$context" ]
    return
  end

  set -l namespace (
    yq e "(.contexts[] | select(.name == \"$context\").context.namespace) // \"\"" \
      ~/.kube/config
  )

  echo -n "$context"

  if [ -n "$namespace" ]
    echo -n " ($namespace)"
  end
end
