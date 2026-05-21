### Mocking
<p align="center"><em>The definition of refactoring is that the code changes but the behaviour stays the same</em></p>

lot of mocking &rarr; bad abstraction

bad design &rarr; poor test

- too many dependencies to mock
  - break the module apart
- dependencies are too fine-grained
  - group dependencies together into modules
- too concerned with implementation
  - test expected behaviour rather than implementation
