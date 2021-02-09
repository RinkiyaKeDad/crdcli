# CRD CLI

Creating a CLI to produce policy report resources based on CIS Benchmarks.

## Demo

[See On Google Drive: CRD CLI In Action](https://drive.google.com/file/d/1HLusazPybyfTSGKT2h2td5fQXbrJtL9p/view)

## Getting Started

1. Clone the repo.
2. Navigate to the project directory and run

```
./kube-bench --json > logs.json
```

3. Add the CRDs to your cluster:

```
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/wgpolicyk8s.io_policyreports.yaml

```

4. Run the following command in the project directory to generate your policy report resource yaml file.

```
go run main.go produce
```

5. Create the policy report resource

```
kubectl create -f actual-resource.yaml

```

6. View policy report resources

```
kubectl get policyreports

```

> Currently only result summary is being edited and rest is mock data.
