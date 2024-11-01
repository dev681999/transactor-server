# DB Schema

This package contains all our database schema.

### How to add a new table?

1. Create a new file help, just copy one of the existing files
2. Modify the type name and add fields, edges & indexes as required
3. Then run `go generate` or `make generate` from root of the project

### How to add a new field, relationship, index?

Just edit the required type and run `go generate` or `make generate` from root of the project!
