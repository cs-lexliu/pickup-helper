# Pickup Helper
---

## Execution Instructions

### 1. Forward the port from k8s pod
* forward the ordermanager port to local port 10001
* forward the shopping-checkout port to local port 10002

The example command is as follows:
```bash
    # change pod name accordingly, it will be changed after each deployment.
    kubectl port-forward -n ordermanager pod/ordermanager-cffd48745-l99ml 10001:9281
    kubectl port-forward -n shopping-checkout pod/shopping-checkout-85bc4c64b7-t4fcz 10002:9281
```

### 2. Change the order id
* Change the order id in the main.go file

### 3. Execute the program
* Run the program with the following command:
```bash
    go run main.go
```
