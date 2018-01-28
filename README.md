This little program's *Welcome* message says it all:

    REPL for our demo 'NanoCalc'
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
Read-Eval-Print-Loop (REPL) of a minimalist tiny language
with multiple (principally / architecturally "pluggable")
interpreters.

**That is: a minimal working skeleton for developing
(in Go) custom languages to be lexed-parsed-and-interpreted.**

Two interpreters are built-in:

- **eval** — arithmetic reduction of parsed expression tree to final numeric result
- **pretty-print** — string-formatting of parsed expression tree

(Other "interpreters" could be optimizers, byte-code
generators, transpilers, compilers etc.)

Furthermore, **two approaches to interpretation** of the
syntax-tree (each sporting its own *eval* and own
*pretty-print* implementation) are included:

- `approach-adt-interp.go` — used by default and
probably the more intuitive, idiomatic, common
approach — also in retrospect, at least in Go,
the terser and more comprehensible one;

- `approach-alt-interp.go` — an attempt to
implement the "tagless final" approach cited
in more detail below. While this would be
the much more desirable approach in a language
such as Haskell, in Go it immediately necessitates
such endless plumbing and hard-coding that any
theoretical gains are overshadowed by code bloat.
(Especially so if one attempted to progress
*beyond* `lecture.pdf`'s chapter 2...)


## Original intention:

~~Primarily, this "toy-lang REPL"s commit history
mirrors my chapter-by-chapter walk through
http://okmij.org/ftp/tagless-final/course/lecture.pdf~~

~~Many of of the various benefits outlined by Oleg
would technically never *really materialize* in Go
as powerfully as they would in (for example) Haskell,
of course. I don't care about actual, compile-time
type-checked EDSLs-for-Go here anyway, just the
general parse+eval strategy still arising from it all.
Much of the elegance and terseness a Haskell
implementation would enjoy, will naturally be lost
in Go: no type-variables, no parametric polymorphism,
no equivalent (in power) to Haskell type-classes etc.~~

~~However, the notion of an alternative approach
to ADTs (aka. "countless type-switches" in Go),
promising "more pluggable" evaluators and
more-conveniently-extensible dialects, sounds
worth exploring and tinkering with.~~

~~It could also well be that outside of a language
like Haskell, the gains are too small and the costs
too high. Only one way to find out! (Well, only-one
for *me*, that is.)~~
