function kubernetes-context {
  local context namespace

  if [ -n "$KUBECONFIG" ]; then
    echo -n "\$KUBECONFIG is not supported"
    return
  fi

  context=$(yq e '.current-context // ""' ~/.kube/config)

  if [ -z "$context" ]; then
    return
  fi

  namespace=$(
    yq e "(.contexts[] | select(.name == \"$context\").context.namespace) // \"\"" \
      ~/.kube/config
  )

  echo -n "$context"

  if [ -n "$namespace" ]; then
    echo -n " ($namespace)"
  fi
}
