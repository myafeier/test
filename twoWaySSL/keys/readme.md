# 双向校验证书签发帮助
> 参考站点：
>> https://smallstep.com/docs/getting-started/
>> https://smallstep.com/hello-mtls/doc/combined/go/php


## 使用流程
- Configuring your Certificate Authority
  ```
  step ca init
    ✔ What would you like to name your new PKI? (e.g. Smallstep): Example Inc.
    ✔ What DNS names or IP addresses would you like to add to your new CA? (e.g. ca.smallstep.com[,1.1.1.1,etc.]): localhost
    ✔ What address will your new CA listen at? (e.g. :443): 127.0.0.1:8443
    ✔ What would you like to name the first provisioner for your new CA? (e.g. you@smallstep.com): bob@example.com
    ✔ What do you want your password to be? [leave empty and we'll generate one]: abc123

    Generating root certificate...
    all done!
  ```
  
- Running your Certificate Authority
  `step-ca $(step path)/config/ca.json`
- Generate crt and key
  - generate server crt and key
  `step ca certificate localhost srv.crt srv.key`
  - copy ca.crt
  `step ca root ca.crt`
  - generate client crt and key
  `step ca certificate "client" client.crt client.key`
- 

