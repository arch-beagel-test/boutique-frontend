# Runbook: boutique-frontend

- **Language:** Go
- **Owner:** @arch-beagel-test/storefront
- **Calls:** productcatalogservice, currencyservice, cartservice, recommendationservice, shippingservice, checkoutservice

## It's down

1. `kubectl logs deploy/frontend --tail=200`
2. `kubectl rollout restart deploy/frontend`
3. If a dependency is failing, its runbook is in that repository's `docs/runbook.md`.

## Before changing who it calls

Read boutique-deploy `adr/cross/0002` and `0003`. A new call is a change to `protos/demo.proto` first.

## Payment status on the confirmation page (issue #3)

Not built here on purpose. The frontend must not call paymentservice (boutique-deploy `adr/cross/0003`: only checkoutservice does; `0002`: a new call starts as a change to `protos/demo.proto`). `OrderResult` carries no payment fields today. To show "Payment: <status> (<transaction id>)":

1. boutique-deploy: add payment status and transaction ID to `OrderResult` in `protos/demo.proto` (platform review).
2. boutique-checkoutservice: fill them from the `Charge` response.
3. Here: regenerate `genproto` and render them in `templates/order.html`.
