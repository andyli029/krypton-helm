
# XenonDB

<<<<<<< HEAD
> English | [中文](README_zh.md) 

## What is XenonDB

[XnenoDB](https://github.com/radondb/xenondb) is a open-source, cloud-native, highly available cluster solutions that is based on [MySQL](https://MySQL.org) database. With the Raft protocol，XenonDB provide the faster failover performance without losing any transactions. 

## Architecture

- Achieving decentralized election through the Raft protocol.
=======
English | [中文](README_zh.md) 

## What is XenonDB？

[XnenoDB](https://github.com/radondb/xenondb) is a open-source, cloud-native, highly available cluster solutions that is based on [MySQL](https://MySQL.org) database. With the Raft protocol，XenonDB provide the faster failover performance without losing any transactions. 

At present, XenonDB has supported the deployment of MySQL high availability clusters on kubernetes and kubesphere platforms.

## Architecture

- Achieving decentralized selection through the Raft protocol.
>>>>>>> 4bcce44 (modify README.md and add README_zh.md #51)
- Synchronize data based on GTID mode through Semi-Sync.

![](docs/images/XenonDB_Architecture_1.png)

<<<<<<< HEAD
<<<<<<< HEAD
## Features

- High availability MySQL database
    - Non-centralized automatic leader election
=======
## Features

- High availability MySQL database
    - Non-centralized automatic leader selection
>>>>>>> 4bcce44 (modify README.md and add README_zh.md #51)
    - Second level switch leader to follower 
    - Strongly consistent data for cluster switching
- Cluster management
- Monitoring and alerting
- Logs
- Account management

## Installation

<<<<<<< HEAD
There are support that depoy XenonDB on the Kubernetes or Kubesphere platforms.
=======
There are support that depoy XneonDB on the Kubernetes or Kubesphere platforms.
>>>>>>> 4bcce44 (modify README.md and add README_zh.md #51)

- [Deploy xenonDB on Kubernetes](docs/Kubernetes/deploy_xenondb_on_kubernetes.md) 

- [Deploy xenonDB on the appstore of KubeSphere](docs/KubeSphere/deploy_xenondb_on_kubesphere.md)

## Release

<<<<<<< HEAD
| Release | Features  |
|------|--------|
| Helm | High availability <br> Non-centralized automatic leader election <br>  Second level switch <br> Strongly consistent data <br> Cluster management <br> Monitoring and alerting <br> Logs <br> Account management | 
| Operator 1.0 | Node management <br> Automatic expansion and shrinkage capacity <br> Upgrade <br> Backups and Restorations <br> Automatic failover <br> Automatic rebuild node <br> Automatic restart service（all or signal node）<br> Account management（API）<br> Migrating Data online | 
| Operator 2.0 | Automatic O&M <br> Multiple node roles <br> Disaster Recovery <br> SSL transmission encryption  | 

## Who are using XenonDB

![](docs/images/users.png)

## License

XenonDB is released under the Apache 2.0, see [LICENSE](./LICENSE).

=======
| Release | Features  | Mode |
|------|--------|--------|
| Helm | High availability <br> Non-centralized automatic leader selection <br>  Second level switch <br> Strongly consistent data <br> Cluster management <br> Monitoring and alerting <br> Logs <br> Account management | Helm |
| Operator 1.0 | Node management <br> Automatic expansion and shrinkage capacity <br> Upgrade <br> Backups and Restorations <br> Automatic failover <br> Automatic rebuild node <br> Automatic restart service（all or signal node）<br> Account management（API）<br> Migrating Data online | Operator |
| Operator 2.0 | Automatic O&M <br> Multiple node roles <br> Disaster Recovery <br> SSL transmission encryption  | Operator |
## Who are using XenonDB？

![](docs/images/users.png)
## License

XenonDB is released under the Apache 2.0, see [LICENSE](./LICENSE).
>>>>>>> 4bcce44 (modify README.md and add README_zh.md #51)
## Discussion and Community

- Forum
  
<<<<<<< HEAD
  The XenonDB topic in [Kubesphere Community](https://github.com/kubesphere/community).

- WeChat group
   
   ![](docs/images/wechat_group.png)

=======
  The XenonDB topic in [Kubesphere Community](https://kubesphere.com.cn/forum/).

- WeChat group

![](docs/images/wechat_group.png)

>>>>>>> 4bcce44 (modify README.md and add README_zh.md #51)
Please submit any XenonDB bugs, issues, and feature requests to XenonDB GitHub Issue.
