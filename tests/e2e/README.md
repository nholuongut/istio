# End-to-End Testing

This directory contains Istio end-to-end tests and associated test framework.

E2E tests are meant for ensure functional correcttness in an E2E environment to make sure Istio works with one or more deployments. For now, these tests run with GKE in Prow and Minikube in CircleCI in both pre-submit and post-submit stages. Their results can be found in https://prow.istio.io/ and https://k8s-testgrid.appspot.com/istio.

Developers, on the other hand, are recommended to run the tests locally before sending out any PR.


# Running E2E Tests

## Using a local VM
E2E tests can be run on your local machines. It helps local testing and debuging. You can use one of the following to set up a local testing environment.

1. See [vagrant/README](local/vagrant/README.md) for instructions to set up a local Vagrant VM environment to run E2E tests.

2. See [minikube/README](local/minikube/README.md) for instructions to set up a Minikube VM environment to run E2E tests.

All lcoal testing options requires the `--use-local-cluster` flag so the framework will not create a LoadBalancer and talk directly to the Pod running istio-ingress.


## Using GKE
Optionally, you can set up a GKE environment to run the E2E tests. See [instructions](UsingGKE.md).


## Using CI (PR pre-submit stage)
You can send a PR to trigger all E2E tests in CI, but you should run these tests locally before sending it out to avoid wasting valuable and shared CI resources.

By default, CI does not run all GKE based E2E tests in pre-submit, but they are be triggered manually using the following commands after "prow/istio-presubmit" completes and has generated the required artifacts for testing.

`/test e2e-suite-rbac-no_auth`

`/test e2e-suite-rbac-auth`

`/test e2e-cluster_wide-auth`


## Debugging
The requests from the containers deployed in tests are performed by `client` program.
For example, to perform a call from a deployed test container to https://console.bluemix.net/, run:

```bash
kubectl exec -it <test pod> -n <test namespace> -c app -- client -url https://console.bluemix.net/
```

To see its usage run:

```
kubectl exec -it <test pod> -n <test namespace> -c app -- client -h
```


# Adding New E2E Tests
[demo_test.go](tests/bookinfo/demo_test.go) is a sample test that covers four cases: default routing, versiong routing, failt delay, and version migration.
Each case applies traffic rules and then clean up after the test. It can serve as a reference for building new test cases.

See below for guidelines for creating a new E2E testsr.
1. Create a new commonConfig for framework and add app used for this test in setTestConfig().
   Each test file has a `testConfig` handling framework and test configuration.
   `testConfig` is a cleanable structure which has  `Setup` and `Teardown`. `Setup` will run before all tests and `Teardown`
   is going to clean up after all tests.
2. Framework would handle all setting up: install and setup istio, deploy app.
3. Setup test-specific environment, like generate rule files from templates and apply routing rules.
   These could be done in `testConfig.Setup()` and would be executed by cleanup register right after framework setup.
4. Write a test. Test case name should start with 'Test' and using 't *testing.T' to log test failures.
   There is no guarantee for running order


# Test Framework Notes

The E2E test framework defines and creates structures and processes for creating cleanable test environments:
install and setup istio modules and clean up afterward.

Writing new tests doesn't require knowledge of the framework.

- framework.go: `Cleanable` is a interface defined with setup() and teardown(). While initialization, framework calls setup() from all registered cleanable
structures and calls teardown() while framework cleanup. The cleanable register works like a stack, first setup, last teardown.

- kubernetes.go: `KubeInfo` handles interactions between tests and kubectl, installs istioctl and apply istio module. Module yaml files are in store at
[install/kubernetes/templates](../../install/kubernetes/templates) and will finally use all-in-one yaml [istio.yaml](../../install/kubernetes/istio.yaml)

- appManager.go: gather apps required for test into a array and deploy them while setup()

