package pgx

import "github.com/lalternativefabrique/eventkit/pkg/consumer"

// Compile-time proof that *Store implements the consumer's IdempotencyStore.
var _ consumer.IdempotencyStore = (*Store)(nil)
