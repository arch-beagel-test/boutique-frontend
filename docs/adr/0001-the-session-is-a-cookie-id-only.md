# The session is a cookie holding an ID only

- **Status:** Accepted
- **Date:** 2026-03-25
- **Applies to:** boutique-frontend
- **Deciders:** @arch-beagel-test/storefront

## Context

The frontend needs to know whose cart to show.

## Decision

The `shop_session-id` cookie holds a random ID and nothing else. The cart lives in cartservice under that ID.

## Consequences

No personal data in cookies.
