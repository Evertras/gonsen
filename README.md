# gonsen

Sometimes, you want to build a service with a basic web interface, but you don't
want to have to deal with building a full API and spin up a SPA with some heavy
framework, etc.

Enter gonsen!

![Todo
example](https://user-images.githubusercontent.com/5923958/165977983-86e2452e-1277-41e4-9844-91a9018aceb7.png)

## What is it?

Gonsen is a simple templating system that lets you write HTML pages using Go
templates, and serve them with type safety using generics.  It's just a simple
wrapper around Go's existing templating system.

Gonsen works with either a file system on disk or using the `embed` package.
Using `embed` is highly recommended with Gonsen as it will bundle the entire
site into the compiled binary, making your service easily portable.

## What is it not?

Gonsen is not a full routing framework.  It is intended to be used with
something like Gorilla Mux or other routing frameworks that already exist.

## How it works

For a code example to follow, see [the basic example](./examples/basic).

Gonsen is intended to be used with the `embed` standard package.  You supply a
basic skeletal structure of what your pages should look like, and then each page
can be defined as simpler components similar to something like Vue.

The skeletal structure can include any number of subcomponents.  The basic
example includes subcomponents for style, scripts, and the main HTML body, but
you could add other subcomponents for a more dynamic menu, etc.

The given filesystem must be in the following structure:

```
site
 |-> pages/
 |   |-> page1.html
 |   |-> another.html
 |-> base.html
```

The `base.html` file is used as the base template, and all files found in
`pages` will use `base.html`.  Any other directories are ignored by Gonsen, but
it can be useful to include static assets and such in other directories as the
example shows.

## Trying it out

`make` will run [the basic example](./examples/basic).

## Where this could go

This is really just a simple wrapper for now, and it's almost more a proof of
concept than a proper library.  Is there room for it to grow into something
more?  Let me know in the Discussions tab if you have any thoughts!

## What's with the name?

I didn't want a [SPA](https://developer.mozilla.org/en-US/docs/Glossary/SPA),
I wanted to build an
[温泉](https://en.wikipedia.org/wiki/Onsen) instead.
