package setup

import (
	_ "divsperf/script/addon/parallel"
	// addon
	_ "divsperf/script/addon/use"
	// analyzer
	_ "divsperf/script/analyzer/dnsanalyzer"
	// conner
	_ "divsperf/script/conner/udp"
	// reporter
	_ "divsperf/script/reporter/dnsreporter"
)