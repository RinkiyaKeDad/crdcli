# CRD CLI

Creating a baisc CLI to produce policy report resources based on CIS Benchmarks.

## Getting Started

1. Clone the repo.
2. Add the CRDs to your cluster:

```
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/wgpolicyk8s.io_policyreports.yaml

```

3. Run the following command in the project directory to generate your policy report resource yaml file.

```
go run main.go produce
```

4. Create the policy report resource

```
kubectl create -f actual-resource.yaml

```

5. View policy report resources

```
kubectl get policyreports

```
