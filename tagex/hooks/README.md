# Lifecycle Hooks

Tagex provides two optional lifecycle interfaces: `PreProcessor` and `PostProcessor`.

You can implement these by adding the following methods to your struct:

* `Before() error` for `PreProcessor`
* `After() error` for `PostProcessor`

These methods are automatically invoked by `tagex.ProcessStruct()` in this order:

1. `Before()` — *before* tag processing
2. Tag processing
3. `After()` — *after* tag processing

This mechanism is ideal for running prerequisite or post-processing logic, such as:

* Generating slugs (e.g., from names or titles)
* Calculating conversions (e.g., metric ↔ imperial, currency rates)
* Looking up geo-coordinates from an address
* Deriving hashes
* Generating identifiers (e.g. keys, UUIDs)
* Linking foreign keys or resolving relationships
* Performing cleanup logic

By using these interfaces, you can keep derived data logic close to your data model while keeping tag parsing clean and 
declarative.

You can find a full working example [here](https://github.com/tedla-brandsema/examples/tree/main/tagex/hooks/indirect/main.go).


## Directly invoking the pre- and post-processor 

Pre- and post-processing of structs is a very handy tool to have in situations outside of tags. This is why *Tagex* 
accommodates for the direct invocation of the pre- and post-processor via the `tagex.InvokePreProcessor()` and 
`tagex.InvokePostProcessor()` functions, without the need to create a *tag* or *directive*.

You can find a full working example [here](https://github.com/tedla-brandsema/examples/tree/main/tagex/hooks/direct/main.go).
