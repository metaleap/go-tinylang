# typed-tagless-final-interpreters
Exploration: how smoothly can it be done in Go?

Primarily, this "toy-lang REPL"s commit history
mirrors my chapter-by-chapter walk through
http://okmij.org/ftp/tagless-final/course/lecture.pdf

Many of of the various benefits outlined by Oleg
would technically never *really materialize* in Go
as powerfully as they would in (for example) Haskell,
of course. I don't care about actual, compile-time
type-checked EDSLs-for-Go here anyway, just the
general parse+eval strategy still arising from it all.
Much of the elegance and terseness a Haskell
implementation would enjoy, will naturally be lost
in Go: no type-variables, no parametric polymorphism,
no equivalent (in power) to Haskell type-classes etc.

However, the notion of an alternative approach
to ADTs (aka. "countless type-switches" in Go),
promising "more pluggable" evaluators and
more-conveniently-extensible dialects, sounds
worth exploring and tinkering with.

It could also well be that outside of a language
like Haskell, the gains are too small and the costs
too high. Only one way to find out! (Well, only-one
for *me*, that is.)
