# Go

## Golang structs

### Encapsulation

```
package common

type TableItem struct {
    id        int
    task_name string
    status    string
}

// NewTableItem creates a new TableItem with the given parameters.
func NewTableItem(id int, taskName, status string) TableItem {
    return TableItem{
        id:        id,
        task_name: taskName,
        status:    status,
    }
}
```

### Export directly (public)

```
package common

type TableItem struct {
    ID        int    // Exported
    TaskName  string // Exported
    Status    string // Exported
}

```
