Testing Library/Framework: Go standard library "testing".
Rationale: Align with common Go project conventions without introducing new dependencies.
This suite covers:
- Success path (ReadinessCheck returns nil)
- Failure path with error message propagation
- Context propagation semantics
- Contract that handler returns nil error and encodes status in ManagementResult