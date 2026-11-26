package money

import pb "github.com/GoogleCloudPlatform/microservices-demo/src/frontend/genproto"

// Percent returns pct percent of m, for showing a discount on the cart page.
func Percent(m pb.Money, pct int32) pb.Money {
	units := m.GetUnits() * int64(pct) / 100
	nanos := m.GetNanos() * pct / 100
	return pb.Money{CurrencyCode: m.GetCurrencyCode(), Units: units, Nanos: nanos}
}
