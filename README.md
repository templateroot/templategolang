# templategolang


2025.0314:
  add 
```GO
if !modUtility.EnsureSingleInstance() {
		return
	}

err := modUtility.Utility_Initialize()
	if err != nil {
		fmt.Println("utility initialize failed: ", err)
		return
	}
```

if allowed one instance only in one system; call it at the first; like above
if allowed one instance each token, call it after config init; like below

```GO
err := modUtility.Utility_Initialize()
	if err != nil {
		fmt.Println("utility initialize failed: ", err)
		return
	}

`if !modUtility.EnsureSingleInstance() {
		return
	}
```
