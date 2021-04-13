# temporal-go ExecuteLocalActivity nil struct issue

When registering activites using a struct as such;

## Activities  
```go
type BeautifulActivities struct {
}

var MyBeautifulActivities *BeautifulActivities

func NewBeautifulActivities() *BeautifulActivities {
	return &BeautifulActivities{}
}
func (a *BeautifulActivities) ComposeGreeting(name string) (string, error) {
	if a == nil {
		// This happens when
		fmt.Println("Ermaghd me pointer is nil!")
		panic("Ermaghd me pointer is nil!")
	} else {
		fmt.Println("Nothing to see here, everything is A-OK!")
	}
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
```
## Registering Activities  
```go
act := app.NewBeautifulActivities()
w.RegisterActivity(act)
``` 
Results in the pointer to **\*BeautifulActivities** being nil if the activity is executed locally.  
```go 
func (a *BeautifulActivities) ComposeGreeting(name string) (string, error) {
// a is nil - panic
}
```  

```go
localOptions := workflow.LocalActivityOptions{
		StartToCloseTimeout: time.Second * 5,
}
ctx = workflow.WithLocalActivityOptions(ctx, localOptions)
err = workflow.ExecuteLocalActivity(ctx, MyBeautifulActivities.ComposeGreeting, name).Get(ctx, &result)
```

This does not happen when the activity is exeucted normally.  
```go
options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
}
ctx = workflow.WithActivityOptions(ctx, options)
var result string
err := workflow.ExecuteActivity(ctx, MyBeautifulActivities.ComposeGreeting, name).Get(ctx, &result)
```
