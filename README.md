# typed-tagless-final-interpreters
Exploration: how smoothly can it be done in Go?

Primarily, this "toy-lang REPL"s commit history
mirrors my chapter-by-chapter walk through
http://okmij.org/ftp/tagless-final/course/lecture.pdf

Many of of the various benefits outlined by Oleg
would technically never *really materialize* in Go
as powerfully as they would in (for example) Haskell,
of course:

However, the notion of an alternative approach
to ADTs (aka. "countless type-switches" in Go),
promising "more pluggable" evaluators and
more-conveniently-extensible dialects, sounds
worth exploring.