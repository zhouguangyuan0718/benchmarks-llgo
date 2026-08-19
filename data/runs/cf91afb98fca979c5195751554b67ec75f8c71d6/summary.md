# LLGo binary-size CI
All values are ELF file sizes in bytes, collected by Bent `benchsize`.

| Benchmark | Go | LLGoNoLTO | LLGoDeadcodeDrop | LLGoFullLTONoGlobalDCE | LLGoFullLTOGlobalDCE | LLGoFullLTOGlobalDCEPlugin |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| Aws_restjson | 14649316 | 14203368 | 10522264 | 14102472 | 11911288 | 11911288 |
| Dustin_humanize | 4998276 | 5232288 | 3581408 | 5082624 | 3777752 | 3777752 |
| Etcdctl | 25910623 | 24552184 | 23258432 | 24234888 | 23639880 | 23639880 |
| Gorm_schema | 9430131 | 7769376 | 5238024 | 7655320 | 7394200 | 5992824 |
| K8s_workqueue | 10177137 | 11808568 | 7321216 | 11730752 | 11568744 | 11568744 |
| Toml | 7217918 | 6633592 | 4740720 | 6497728 | 5472792 | 5472792 |
| Uber_zap | 8738630 | 12643656 | 7975056 | 12581856 | 10159120 | 10159120 |
| XGo | 18671613 | 19366288 | 17009296 | 19392112 | 18928656 | 18928656 |
