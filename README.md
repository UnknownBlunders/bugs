# bug-tracker

## Usage:

To get your list of bugs:

``` bash
$ go run cmd/main.go
```

To add a bug to your list titled "Can't delete bugs":

``` bash
$ go run cmd/main.go "Can't delete bugs"
```

Output examples:

``` bash
$ go run cmd/main.go
# Status Title
=================
0 Open   tests broken
1 Open   Can't update status of bugs
2 Open   Can't delete bugs
```

Future Command list:

```
bt list
bt help
bt #(alias bt help)
```