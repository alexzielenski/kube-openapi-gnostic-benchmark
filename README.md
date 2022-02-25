# kube-openapi-gnostic-benchmark
benchmarks for kube-openapi gnostic conversion

## OpenAPI V2

### JS
```
goos: darwin
goarch: amd64
pkg: github.com/alexzielenski/parsebench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkSlowConversion
BenchmarkSlowConversion/json->swagger
BenchmarkSlowConversion/json->swagger-8         	       2	 559821998 ns/op	95734604 B/op	 1381629 allocs/op
BenchmarkSlowConversion/swagger->json
BenchmarkSlowConversion/swagger->json-8         	       8	 175216241 ns/op	70629431 B/op	  274052 allocs/op
BenchmarkSlowConversion/json->gnostic
BenchmarkSlowConversion/json->gnostic-8         	       5	 311571265 ns/op	81431456 B/op	 1248391 allocs/op
BenchmarkSlowConversion/gnostic->pb
BenchmarkSlowConversion/gnostic->pb-8           	     102	  10260869 ns/op	 2899970 B/op	       1 allocs/op
BenchmarkSlowConversion/pb->gnostic
BenchmarkSlowConversion/pb->gnostic-8           	      92	  13718802 ns/op	 9480703 B/op	  123829 allocs/op
BenchmarkSlowConversion/gnostic->yaml
BenchmarkSlowConversion/gnostic->yaml-8         	      38	  45272785 ns/op	32855103 B/op	  264561 allocs/op
BenchmarkSlowConversion/yaml->json
BenchmarkSlowConversion/yaml->json-8            	      15	  69750500 ns/op	24687740 B/op	  463782 allocs/op
BenchmarkSlowConversion/json->swagger#01
BenchmarkSlowConversion/json->swagger#01-8      	       2	 617799926 ns/op	95791084 B/op	 1381634 allocs/op
PASS
```

### YAML
```
goos: darwin
goarch: amd64
pkg: github.com/alexzielenski/parsebench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkFastConversion
BenchmarkFastConversion/json->swagger
BenchmarkFastConversion/json->swagger-8         	       2	 566188894 ns/op	95734740 B/op	 1381635 allocs/op
BenchmarkFastConversion/swagger->json
BenchmarkFastConversion/swagger->json-8         	       7	 151444487 ns/op	73029682 B/op	  274082 allocs/op
BenchmarkFastConversion/json->gnostic
BenchmarkFastConversion/json->gnostic-8         	       5	 222287071 ns/op	81491691 B/op	 1248403 allocs/op
BenchmarkFastConversion/gnostic->pb
BenchmarkFastConversion/gnostic->pb-8           	     100	  10269276 ns/op	 2899971 B/op	       1 allocs/op
BenchmarkFastConversion/pb->gnostic
BenchmarkFastConversion/pb->gnostic-8           	     100	  12733826 ns/op	 9480702 B/op	  123829 allocs/op
BenchmarkFastConversion/gnostic->yaml
BenchmarkFastConversion/gnostic->yaml-8         	      37	  30927474 ns/op	32855214 B/op	  264562 allocs/op
BenchmarkFastConversion/yaml->swagger
BenchmarkFastConversion/yaml->swagger-8         	      19	  67597440 ns/op	23686302 B/op	  510368 allocs/op
PASS
ok  	github.com/alexzielenski/parsebench	12.068s
```

## OpenAPI V3

### JSON
```
goos: darwin
goarch: amd64
pkg: github.com/alexzielenski/parsebench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkSlowConversionV3
BenchmarkSlowConversionV3/json->kube
BenchmarkSlowConversionV3/json->kube-8         	       2	 677217190 ns/op	149247920 B/op	 1835203 allocs/op
BenchmarkSlowConversionV3/json->gnostic
BenchmarkSlowConversionV3/json->gnostic-8      	       3	 337665597 ns/op	116477536 B/op	 1528741 allocs/op
BenchmarkSlowConversionV3/gnostic->pb
BenchmarkSlowConversionV3/gnostic->pb-8        	      79	  14470868 ns/op	 4304898 B/op	      47 allocs/op
BenchmarkSlowConversionV3/pb->gnostic
BenchmarkSlowConversionV3/pb->gnostic-8        	      61	  18952843 ns/op	14967267 B/op	  170123 allocs/op
BenchmarkSlowConversionV3/gnostic->yaml
BenchmarkSlowConversionV3/gnostic->yaml-8      	      33	  41073664 ns/op	37846709 B/op	  388852 allocs/op
BenchmarkSlowConversionV3/gnostic->json
BenchmarkSlowConversionV3/gnostic->json-8      	      49	  22973946 ns/op	 5869037 B/op	      47 allocs/op
BenchmarkSlowConversionV3/json->kube#01
BenchmarkSlowConversionV3/json->kube#01-8      	       2	 617506573 ns/op	149449976 B/op	 1835730 allocs/op
PASS
ok  	github.com/alexzielenski/parsebench	13.407s
```

### YAML
```
goos: darwin
goarch: amd64
pkg: github.com/alexzielenski/parsebench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkFastConversionV3
BenchmarkFastConversionV3/json->kube
BenchmarkFastConversionV3/json->kube-8         	       2	 635909210 ns/op	149248800 B/op	 1835207 allocs/op
BenchmarkFastConversionV3/json->gnostic
BenchmarkFastConversionV3/json->gnostic-8      	       4	 318940632 ns/op	116516990 B/op	 1528742 allocs/op
BenchmarkFastConversionV3/gnostic->pb
BenchmarkFastConversionV3/gnostic->pb-8        	      73	  16141946 ns/op	 4304898 B/op	      47 allocs/op
BenchmarkFastConversionV3/pb->gnostic
BenchmarkFastConversionV3/pb->gnostic-8        	      60	  20002848 ns/op	14967266 B/op	  170123 allocs/op
BenchmarkFastConversionV3/gnostic->yaml
BenchmarkFastConversionV3/gnostic->yaml-8      	      30	  45300212 ns/op	37846700 B/op	  388852 allocs/op
BenchmarkFastConversionV3/yaml->kube
BenchmarkFastConversionV3/yaml->kube-8         	      12	  91225051 ns/op	37919650 B/op	  733865 allocs/op
PASS
ok  	github.com/alexzielenski/parsebench	12.126s
```