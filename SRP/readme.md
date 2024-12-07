
# SRP 
```
SRP states a type should have one primary responsibility and as a result it should have one reason to change
```

- Lets say we have a journal. We have some methods to add entries and remove entries from our journal.

```
type Journal struct {
Entries []string
}

func (j *Journal) AddEntry(text string) int

func (j *Journal) RemoveEntry(idx int) int {


lets say we need to save our journal to disk

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}
This breaks SRP as there is no sepration of concern. (GOD OBJECT)

```

## Components and Their Responsibilities

### 1. Journal Package (`journal.go`)
- **Single Responsibility**: Managing journal entries
- **Key Functions**:
  - `AddEntry()`: Adds new entries to the journal
  - `RemoveEntry()`: Removes entries by index
  - `String()`: Provides string representation of entries
- **Why It's SRP Compliant**: 
  - Only focuses on journal entry management
  - Doesn't handle persistence or external operations
  - Changes to entry management won't affect other components

### 2. Persistence Package (`persistence.go`)
- **Single Responsibility**: Handling data persistence
- **Key Functions**:
  - `SaveToFile()`: Saves journal data to a file
- **Why It's SRP Compliant**:
  - Only handles saving/loading operations
  - Can be modified without affecting journal logic
  - Can be extended to support different storage methods without changing journal code

### 3. Main Package (`main.go`)
- **Responsibility**: Composition and orchestration
- Acts as a client that brings together the different components
- Demonstrates how separate responsibilities can work together

## Benefits of This Implementation
1. **Maintainability**: Each package can be maintained independently
2. **Testability**: Components can be tested in isolation
3. **Flexibility**: Easy to modify or extend individual components
4. **Clarity**: Clear separation of concerns makes code easier to understand

## Example Usage
```
func main() {

// Create journal instance

j := journal.Journal{}

// Create persistence instance

p := persistence.Persistence{LineSeparator: "\n"}

// Use journal functions

j.AddEntry("I am sad")

j.AddEntry("I am dead")

j.RemoveEntry(1)

// Use persistence to save

p.SaveToFile(&j, "journal11.txt")

}
```

## Violation Examples
To understand SRP better, here are examples of what would violate SRP:

1. **Bad Practice**: Putting file operations in Journal struct
2. **Bad Practice**: Mixing persistence logic with entry management