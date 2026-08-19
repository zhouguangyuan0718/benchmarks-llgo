## Build timing diagnostics

Native Bent `-report-build-time` records, sorted by CPU time (`user + sys`, slowest first). Wall time is diagnostic only.

| Benchmark | Configuration | CPU (user + sys) | User | Sys | Wall (reference) |
| --- | --- | ---: | ---: | ---: | ---: |
| Etcdctl | LLGoFullLTOGlobalDCE | 327547.8 ms | 322346.3 ms | 5201.5 ms | 214520.8 ms |
| Etcdctl | LLGoFullLTONoGlobalDCE | 316092.3 ms | 311527.2 ms | 4565.1 ms | 214702.6 ms |
| Etcdctl | LLGoFullLTOGlobalDCEPlugin | 315064.5 ms | 310264.5 ms | 4800.0 ms | 210435.4 ms |
| Aws_restjson | LLGoFullLTONoGlobalDCE | 231117.9 ms | 226529.8 ms | 4588.0 ms | 137085.5 ms |
| Etcdctl | LLGoDeadcodeDrop | 200213.9 ms | 196272.9 ms | 3941.0 ms | 60845.6 ms |
| XGo | LLGoFullLTONoGlobalDCE | 189871.0 ms | 186269.7 ms | 3601.4 ms | 141921.0 ms |
| XGo | LLGoFullLTOGlobalDCE | 189549.7 ms | 185741.2 ms | 3808.5 ms | 136867.4 ms |
| Uber_zap | LLGoFullLTOGlobalDCEPlugin | 187881.5 ms | 182902.0 ms | 4979.5 ms | 101806.5 ms |
| XGo | LLGoFullLTOGlobalDCEPlugin | 183328.9 ms | 179704.1 ms | 3624.8 ms | 134193.8 ms |
| Uber_zap | LLGoDeadcodeDrop | 177693.3 ms | 173897.3 ms | 3796.0 ms | 54770.4 ms |
| Uber_zap | LLGoFullLTOGlobalDCE | 164651.4 ms | 160744.6 ms | 3906.8 ms | 90709.6 ms |
| Aws_restjson | LLGoFullLTOGlobalDCE | 134860.2 ms | 131987.4 ms | 2872.8 ms | 102478.4 ms |
| Aws_restjson | LLGoFullLTOGlobalDCEPlugin | 131256.9 ms | 128373.6 ms | 2883.3 ms | 100214.8 ms |
| Uber_zap | LLGoFullLTONoGlobalDCE | 110889.5 ms | 108178.4 ms | 2711.1 ms | 89150.7 ms |
| K8s_workqueue | LLGoFullLTOGlobalDCEPlugin | 99070.8 ms | 96303.1 ms | 2767.7 ms | 80700.8 ms |
| K8s_workqueue | LLGoFullLTOGlobalDCE | 96919.7 ms | 94259.7 ms | 2660.0 ms | 79498.6 ms |
| K8s_workqueue | LLGoFullLTONoGlobalDCE | 96760.1 ms | 94091.4 ms | 2668.7 ms | 79839.2 ms |
| Etcdctl | LLGoNoLTO | 82997.3 ms | 79654.9 ms | 3342.5 ms | 28420.5 ms |
| Gorm_schema | LLGoFullLTONoGlobalDCE | 67980.1 ms | 66211.8 ms | 1768.2 ms | 51592.1 ms |
| Gorm_schema | LLGoFullLTOGlobalDCE | 66473.2 ms | 64762.9 ms | 1710.3 ms | 50068.4 ms |
| Aws_restjson | LLGoNoLTO | 65777.9 ms | 63293.7 ms | 2484.1 ms | 26531.9 ms |
| Aws_restjson | LLGoDeadcodeDrop | 64992.3 ms | 62317.9 ms | 2674.5 ms | 28445.2 ms |
| XGo | LLGoDeadcodeDrop | 62521.2 ms | 59779.6 ms | 2741.6 ms | 22759.7 ms |
| XGo | LLGoNoLTO | 60994.7 ms | 58277.7 ms | 2717.0 ms | 21813.0 ms |
| Gorm_schema | LLGoFullLTOGlobalDCEPlugin | 54200.4 ms | 52474.5 ms | 1725.8 ms | 38057.9 ms |
| Toml | LLGoFullLTONoGlobalDCE | 53658.6 ms | 52134.3 ms | 1524.3 ms | 43834.5 ms |
| Toml | LLGoFullLTOGlobalDCEPlugin | 45844.0 ms | 44292.5 ms | 1551.5 ms | 35395.2 ms |
| Toml | LLGoFullLTOGlobalDCE | 45269.5 ms | 43813.4 ms | 1456.2 ms | 35109.0 ms |
| Dustin_humanize | LLGoFullLTONoGlobalDCE | 42234.7 ms | 40882.5 ms | 1352.2 ms | 35665.6 ms |
| Gorm_schema | LLGoDeadcodeDrop | 39746.4 ms | 37984.5 ms | 1761.9 ms | 12987.6 ms |
| Dustin_humanize | LLGoFullLTOGlobalDCE | 33288.8 ms | 31935.4 ms | 1353.4 ms | 26585.1 ms |
| Dustin_humanize | LLGoFullLTOGlobalDCEPlugin | 32854.5 ms | 31484.3 ms | 1370.2 ms | 25876.4 ms |
| Etcdctl | Go | 31527.1 ms | 29638.5 ms | 1888.6 ms | 9778.5 ms |
| K8s_workqueue | LLGoNoLTO | 28136.1 ms | 25988.3 ms | 2147.8 ms | 9916.4 ms |
| K8s_workqueue | LLGoDeadcodeDrop | 28068.3 ms | 25977.1 ms | 2091.2 ms | 10015.3 ms |
| Uber_zap | LLGoNoLTO | 25740.4 ms | 23617.3 ms | 2123.1 ms | 9430.3 ms |
| Toml | LLGoDeadcodeDrop | 20486.6 ms | 19083.4 ms | 1403.2 ms | 7646.3 ms |
| XGo | Go | 15501.6 ms | 14618.6 ms | 883.0 ms | 4765.6 ms |
| Gorm_schema | LLGoNoLTO | 15017.8 ms | 13586.1 ms | 1431.8 ms | 5691.7 ms |
| Toml | LLGoNoLTO | 12600.3 ms | 11335.1 ms | 1265.2 ms | 4887.0 ms |
| Dustin_humanize | LLGoDeadcodeDrop | 11601.9 ms | 10347.2 ms | 1254.7 ms | 4910.2 ms |
| Dustin_humanize | LLGoNoLTO | 10727.0 ms | 9503.8 ms | 1223.2 ms | 4231.6 ms |
| Aws_restjson | Go | 4858.2 ms | 4426.5 ms | 431.7 ms | 1870.7 ms |
| Gorm_schema | Go | 3374.5 ms | 3171.9 ms | 202.7 ms | 1460.0 ms |
| Uber_zap | Go | 2999.9 ms | 2777.5 ms | 222.4 ms | 1316.4 ms |
| K8s_workqueue | Go | 1772.5 ms | 1572.8 ms | 199.7 ms | 817.0 ms |
| Dustin_humanize | Go | 783.3 ms | 633.2 ms | 150.1 ms | 381.8 ms |
| Toml | Go | 547.4 ms | 425.8 ms | 121.6 ms | 307.3 ms |

### Configuration totals

| Configuration | Total CPU (user + sys) | Total wall (reference) | Cases |
| --- | ---: | ---: | ---: |
| LLGoFullLTONoGlobalDCE | 1108604.2 ms | 793791.1 ms | 8 |
| LLGoFullLTOGlobalDCE | 1058560.4 ms | 735837.2 ms | 8 |
| LLGoFullLTOGlobalDCEPlugin | 1049501.5 ms | 726680.7 ms | 8 |
| LLGoDeadcodeDrop | 605323.9 ms | 202380.4 ms | 8 |
| LLGoNoLTO | 301991.5 ms | 110922.4 ms | 8 |
| Go | 61364.5 ms | 20697.3 ms | 8 |

Dependency download details are in `download-timings.log`.
