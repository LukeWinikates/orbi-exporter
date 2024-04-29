# orbi-exporter
Prometheus Exporter for Orbi Metrics


# Origin Story

My Orbi Router doesn't expose Prometheus metrics, but it does have a page that provides them.

this is built to run via the container service on my synology NAS

## Configuration

The exporter listens on port `6724` (on a telephone keypad, that's `O:6 R:7 B:2 I:4`)

The exporter expects the following environment variables:

```
ORBI_HOST=http://192.168.1.1
ORBI_BASICAUTH_USER=admin
ORBI_BASICAUTH_PASSWORD=<your password>
```

## Compatibility

Tested against an Orbi `RBR20`. Router Firmware Version `V2.7.4.24`



## Interpreting

WAN:

LAN:

WLAN b/g/n: this is the 2.6ghz wifi network?
WLAN a/n/ac: this is the 5.0ghz wifi network
WLAN Backhaul


## the JS Vars

var sys_uptime="3069182";

System updtime in seconds; 3069182 / 60 / 60 / 24 -> ~35.5 days



var lan_status="Link up";


# packets transmitted over lan
var lan_txpkts="88670935";

# packets received over lan
var lan_rxpkts="43826450";
var lan_collisions="0";

# transmitted B/s
var lan_txbs="1345";

# received B/s
var lan_rxbs="729";

# uptime
var lan_systime="3069182";



var wan_status="1000M/Full";
var wan_txpkts="121869567";
var wan_rxpkts="262693070";
var wan_collisions="0";
var wan_txbs="1399";
var wan_rxbs="1399";
var wan_systime="2695";


var bgn_status="400M";
var bgn_txpkts="31060602";
var bgn_rxpkts="60138";
var bgn_collisions="0";
var bgn_txbs="32";
var bgn_rxbs="6";
var bgn_systime="3069110";
var an_status="866M";
var an_txpkts="17510305";
var an_rxpkts="0";
var an_collisions="0";
var an_txbs="0";
var an_rxbs="0";
var an_systime="3069093";
var bh_status="866M";
var bh_txpkts="32365100";
var bh_rxpkts="4094";
var bh_collisions="0";
var bh_txbs="0";
var bh_rxbs="0";
var bh_systime="3069093";
var wwan0_status="no service";
var wwan0_txpkts="0";
var wwan0_rxpkts="0";
var wwan0_collisions="0";
var wwan0_txbs="0";
var wwan0_rxbs="0";
var wwan0_systime="0";
var lan_status0="1000M/Full"
var lan0_systime = "3069133"
var lan_status1="Link down"
var lan1_systime = "0"
var lan_status2="Link down"
var lan2_systime = "0"
var lan_status3="Link down"
var lan3_systime = "0"
