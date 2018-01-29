This little program's *Welcome* message says it all:

    REPL for our demo 'TinyCalc'
    language, consisting only of:
    float operands, parens and the 4
    most basic arithmetic operators
    (with no precedence: use parens).

    Enter:
    · Q to quit
    · A to toggle between:
      · "ADT" interpreter approach (default)
      · "Alt" interpreter approach
    · <expr> to parse-and-prettyprint-and-eval

## Purpose:

A self-contained (stdlib-only, no deps) interactive
Read-Eval-Print-Loop (REPL) of a minimalist tiny uni-typed
language with multiple (principally / architecturally
"pluggable") interpreters.

**That is: a minimal working skeleton for developing
(in Go) custom languages to be lexed-parsed-and-interpreted.**

Two interpreters are built-in:

- **eval** — arithmetic reduction of parsed expression tree to final numeric result
- **pretty-print** — string-formatting of parsed expression tree

(Other "interpreters" —more so in the general case,
less so for the mini-language at hand— could be
optimizers, simplifiers, byte-code generators,
transpilers, compilers etc..)

Furthermore, **two approaches to interpretation** of the
syntax-tree (each sporting its own *eval* and own
*pretty-print* implementation) are included:

- `approach-adt-interp.go` — used by default and
probably the more intuitive, idiomatic, common
approach — also in retrospect, at least in Go,
the terser and more comprehensible one;

- `approach-alt-interp.go` — inspired by
http://okmij.org/ftp/tagless-final/course/lecture.pdf
(chapter 2 only) — while this approach would be much
more desirable in a language such as Haskell (*and*
when targeting *embedded* DSLs), in Go (and targeting
lexed-and-parsed instead of embedded languages) it soon
necessitates tedious-read-and-write code bloat — as it
turns out, much of the "almost-magic" convenience of
tagless-final is truly afforded by the expressive power
of Haskell's type-classes and parametric polymorphism,
and so to transfer the idea to a very-low-level language
would amount to furnishing a code-generator not far in
its capabilities from a Haskell compiler! Not on, for now.
