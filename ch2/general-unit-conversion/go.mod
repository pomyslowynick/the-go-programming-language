module example.com/guc

go 1.16

replace example.com/tempconv => ../tempconv

require (
	example.com/lengthconv v0.0.0-00010101000000-000000000000
	example.com/tempconv v0.0.0-00010101000000-000000000000
	example.com/weightconv v0.0.0-00010101000000-000000000000
)

replace example.com/weightconv => ../weightconv

replace example.com/lengthconv => ../lengthconv
