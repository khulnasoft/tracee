![Tracee Logo](docs/images/tracee.png)

<!-- links that differ between docs and readme -->
[installation]:https://khulnasoft.github.io/tracee/latest/docs/install/
[docker-guide]:https://khulnasoft.github.io/tracee/latest/docs/install/docker/
[kubernetes-guide]:https://khulnasoft.github.io/tracee/latest/docs/install/kubernetes/
[prereqs]:https://khulnasoft.github.io/tracee/latest/docs/install/prerequisites/
[macfaq]:https://khulnasoft.github.io/tracee/latest/docs/advanced/mac/

Before moving on, please consider giving us a GitHub star ⭐️. Thank you!

## About Tracee

Tracee is a runtime security and observability tool that helps you understand how your system and applications behave.  
It is using [eBPF technology](https://ebpf.io/what-is-ebpf/) to tap into your system and expose that information as events that you can consume.  
Events range from factual system activity events to sophisticated security events that detect suspicious behavioral patterns.

To learn more about Tracee, check out the [documentation](https://khulnasoft.github.io/tracee/).

## Quickstart

To quickly try Tracee use one of the following snippets. For a more complete installation guide, check out the [Installation section][installation].  
Tracee should run on most common Linux distributions and kernels. For compatibility information see the [Prerequisites][prereqs] page.  Mac users, please read [this FAQ][macfaq].

### Using Docker

```shell
docker run --name tracee -it --rm \
  --pid=host --cgroupns=host --privileged \
  -v /etc/os-release:/etc/os-release-host:ro \
  -v /var/run:/var/run:ro \
  khulnasoft/tracee:latest
```

For a complete walkthrough please see the [Docker getting started guide][docker-guide].

### On Kubernetes

```shell
helm repo add aqua https://khulnasoft.github.io/helm-charts/
helm repo update
helm install tracee aqua/tracee --namespace tracee --create-namespace
```

```shell
kubectl logs --follow --namespace tracee daemonset/tracee
```

For a complete walkthrough please see the [Kubernetes getting started guide][kubernetes-guide].

## Contributing
  
Join the community, and talk to us about any matter in the [GitHub Discussions](https://github.com/khulnasoft/tracee/discussions) or [Slack](https://slack.khulnasoft.com).  
If you run into any trouble using Tracee or you would like to give use user feedback, please [create an issue.](https://github.com/khulnasoft/tracee/issues)

Find more information on [contribution documentation](https://khulnasoft.github.io/tracee/latest/contributing/overview/).

## More about Khulnasoft Security

Tracee is an [Khulnasoft Security](https://khulnasoft.com) open source project.  
Learn about our open source work and portfolio [here](https://www.khulnasoft.com/products/open-source-projects/).
