package stdimports

// _ "somepackage"  the _ will allow any init() function and variables/constants within that package to load, these are called side effects
// e.g. a driver package might only be useful for it's side effects (it's init() functions) so the package is imported as _ "driverpackage"
// Remember that init() is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.
// see https://stackoverflow.com/questions/32465453/golang-what-is-import-side-effect for some more information
