# Dictionary in Go

This is a simple dictionary application implemented in Go. It uses BadgerDB to store word definitions.

## Features

* **Add:** Add a new word and its definition to the dictionary.

* **Get:** Retrieve the definition of a word from the dictionary.

* **List:** List all words and their definitions in the dictionary.

* **Remove:** Remove a word from the dictionary.

## How to Use

1.  **Clone the repository:**

    ```
    git clone git@github.com:axellelanca/dictionary-CLI.git
    ```

2.  **Navigate to the project directory:**

    ```
    cd dictionary-CLI
    ```

3.  **Run the main application:**

    ```
    go run main.go 
    ```

    Where `<action>` can be one of the following:

    * `list`: Lists all words and their definitions.

    * `add`: Adds a new word and definition (requires two arguments: word and definition).

    * `define`: Retrieves the definition of a word (requires one argument: word).

    * `remove`: Removes a word from the dictionary (requires one argument: word).

## Code Structure

The project is structured as follows:

* `main.go`: Contains the main application logic and command-line argument parsing.

* `dictionary/dictionary.go`: Defines the `Dictionary` struct and its methods for interacting with the BadgerDB.

* `dictionary/actions.go`: Implements the actions performed on the dictionary.

## Details

### `main.go`

* Parses command-line arguments using the `flag` package.

* Creates a new `Dictionary` instance using `dictionary.New("./badger")`.

* Handles errors using the `handleErr` function.

* Calls the appropriate function based on the provided action.

### `dictionary/dictionary.go`

* Defines the `Dictionary` struct, which holds a pointer to a BadgerDB instance.

* Defines the `Entry` struct to represent a word and its definition.

* Provides helper functions, notably `String()`, for displaying entries.

* The `New` function opens a new BadgerDB database.

* The `Close` function closes the database connection.

### `dictionary/actions.go`

* Implements the following functions:

    * `Add`: Adds a new word and definition to the dictionary.

        * The word is converted to title case using `golang.org/x/text/cases`.

        * The `Entry` is encoded using `encoding/gob` and stored in the database.

    * `Get`: Retrieves the definition of a word from the dictionary.

        * The entry is retrieved from the database and decoded using `encoding/gob`.

    * `List`: Retrieves all entries from the dictionary, sorts the words alphabetically, and returns the sorted list of words and the map of entries.

    * `Remove`: Removes a word from the dictionary.

* Helper functions:

    * `sortedKeys`: Returns a sorted slice of keys from a map.

    * `getEntry`: Decodes a database item into an `Entry` struct.

### Error Handling

The `handleErr` function in `main.go` prints any errors that occur during dictionary operations.

## Dependencies

* [dgraph-io/badger/v3](https://github.com/dgraph-io/badger/v3)

* [golang.org/x/text/cases](https://pkg.go.dev/golang.org/x/text/cases)

* [golang.org/x/text/language](https://pkg.go.dev/golang.org/x/text/language)
