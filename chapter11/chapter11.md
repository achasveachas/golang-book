# Problems

- Why do we use packages?

To help keep our code DRY.

- What is the difference between an identifier that starts with a capital letter and one which doesn’t? (`Average` vs `average`)

Capita letters get exported.

- What is a package alias? How do you make one?

If you need to import two packages with the same name you can set an alias for one (or both) of them like this:

```golang
import alias "package"
```

- We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

// TODO

- How would you document the functions you created in #3?

By writing a comment before each one so that `godoc` picks it up.