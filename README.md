# dockerlang
a turing complete language based on docker containers

[![Build Status](https://travis-ci.org/connorwalsh/dockerlang.svg?branch=dev)](https://travis-ci.org/connorwalsh/dockerlang)

## ❦ Contents
- [How to use Dockerlang](##how-to-use-dockerlang)
- [Development guide](#development-guide)

---

## ❦ How to use Dockerlang

H e l l o and welcome, you've made it: you're in Dockerlang, a dockerland where everything is docker.
Let us begin! by creating a dockerlang file! The file extension is `.doc`.

### Hello, wirld

To print text in Dockerlang, all you have to do is use the `#` symbol, which means printing, followed with the location to print to.
If you want to print to standard out, use `!`, or to print to a string, put a string after the `#`.

```dockerlang
(#! "hellow wourld')
```

### (In)variables

Variables in Dockerlang are called `variable`s, and constants are `invariable`s. Here's how to declare a variable!
```dockerlang
(variable x = 0)
(invariable y = 'why")
```

Dockerlang doesn't worry about types and neither should you. Be free.

### Control structures

#### Conditions

- `‽` Equal to: `(2 ‽ 3) slashslash False`
- `¿` Not equal to: `(2 ¿ 3) slashslash True`
- `☛` Greater than: `(0 ☛ 1) slashslash False`
- `☚` Leess than `(0 ☚ 1) slashslash True`

#### Function declaration

Use the `❦` operator to begin a function. Don't forget to type it, then delete it, then type it again, for best (and most auspicious) results.

---

## ❦ Development guide

Dockerlang is written in Go.

### Setting up your development environment
To get started writing Dockerlang, the first thing you neeed to do is put on some Gamelan music.
You will need to specify what you are listening to in your commits.

- Set up your git environment to use the commit template

```bash
git config commit.template .gitmessage
```
